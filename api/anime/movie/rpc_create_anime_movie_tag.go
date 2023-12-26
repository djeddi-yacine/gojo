package animeMovie

import (
	"context"
	"errors"
	"fmt"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovieTag(ctx context.Context, req *ampb.CreateAnimeMovieTagRequest) (*ampb.CreateAnimeMovieTagResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie tags")
	}

	if violations := validateCreateAnimeMovieTagRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBT []string
	if req.AnimeTags != nil {
		DBT = make([]string, len(req.GetAnimeTags()))
		for i, t := range req.GetAnimeTags() {
			DBT[i] = t
		}
	}

	arg := db.CreateAnimeMovieTagTxParams{
		AnimeID:   req.GetAnimeID(),
		AnimeTags: DBT,
	}

	data, err := server.gojo.CreateAnimeMovieTagTx(ctx, arg)
	if err != nil {
		return nil, shared.ApiError("failed to create anime movie tags", err)
	}

	var animeTags []*ampb.AnimeMovieTag
	if len(data.AnimeTags) > 0 {
		animeTags = make([]*ampb.AnimeMovieTag, len(data.AnimeTags))
		for i, t := range data.AnimeTags {
			animeTags[i] = &ampb.AnimeMovieTag{
				ID:        t.ID,
				Tag:       t.Tag,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	res := &ampb.CreateAnimeMovieTagResponse{
		AnimeMovie: shared.ConvertAnimeMovie(data.AnimeMovie),
		AnimeTags:  animeTags,
	}

	return res, nil
}

func validateCreateAnimeMovieTagRequest(req *ampb.CreateAnimeMovieTagRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	if req.AnimeTags != nil {
		if len(req.GetAnimeTags()) > 0 {
			for i, t := range req.GetAnimeTags() {
				if err := utils.ValidateString(t, 1, 300); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("animeTags >  tag at index [%d]", i), err))
				}
			}
		}
	} else {
		violations = append(violations, shared.FieldViolation("animeTags", errors.New("you need to send the other tags in AnimeTags model")))
	}

	return violations
}
