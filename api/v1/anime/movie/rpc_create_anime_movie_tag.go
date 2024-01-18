package amapiv1

import (
	"context"
	"errors"
	"fmt"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieTag(ctx context.Context, req *ampbv1.CreateAnimeMovieTagRequest) (*ampbv1.CreateAnimeMovieTagResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie tags")
	}

	if violations := validateCreateAnimeMovieTagRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieTagTxParams{
		AnimeID: req.GetAnimeID(),
	}

	if req.AnimeTags != nil {
		arg.AnimeTags = make([]string, len(req.GetAnimeTags()))
		for i, v := range req.GetAnimeTags() {
			arg.AnimeTags[i] = v
		}
	}

	data, err := server.gojo.CreateAnimeMovieTagTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie tags", err)
	}

	res := &ampbv1.CreateAnimeMovieTagResponse{
		AnimeMovie: convertAnimeMovie(data.AnimeMovie),
		AnimeTags:  aapiv1.ConvertAnimeTags(data.AnimeTags),
	}

	return res, nil
}

func validateCreateAnimeMovieTagRequest(req *ampbv1.CreateAnimeMovieTagRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.AnimeTags != nil {
		if len(req.GetAnimeTags()) > 0 {
			for i, v := range req.GetAnimeTags() {
				if err := utils.ValidateString(v, 1, 300); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTags >  tag at index [%d]", i), err))
				}
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("animeTags", errors.New("you need to send the other tags in AnimeTags model")))
	}

	return violations
}
