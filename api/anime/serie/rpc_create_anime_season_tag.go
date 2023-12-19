package animeSerie

import (
	"context"
	"errors"
	"fmt"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSeasonTag(ctx context.Context, req *aspb.CreateAnimeSeasonTagRequest) (*aspb.CreateAnimeSeasonTagResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season tags")
	}

	if violations := validateCreateAnimeSeasonTagRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBT []string
	if req.SeasonTags != nil {
		DBT = make([]string, len(req.GetSeasonTags()))
		for i, t := range req.GetSeasonTags() {
			DBT[i] = t
		}
	}

	arg := db.CreateAnimeSeasonTagTxParams{
		SeasonID:   req.GetSeasonID(),
		SeasonTags: DBT,
	}

	data, err := server.gojo.CreateAnimeSeasonTagTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to create anime season tags", err)
	}

	var seasonTags []*aspb.AnimeSeasonTag
	if len(data.SeasonTags) > 0 {
		seasonTags = make([]*aspb.AnimeSeasonTag, len(data.SeasonTags))
		for i, t := range data.SeasonTags {
			seasonTags[i] = &aspb.AnimeSeasonTag{
				ID:        t.ID,
				Tag:       t.Tag,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	res := &aspb.CreateAnimeSeasonTagResponse{
		AnimeSeason: shared.ConvertAnimeSeason(data.AnimeSeason),
		SeasonTags:  seasonTags,
	}

	return res, nil
}

func validateCreateAnimeSeasonTagRequest(req *aspb.CreateAnimeSeasonTagRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shared.FieldViolation("seasonID", err))
	}

	if req.SeasonTags != nil {
		if len(req.GetSeasonTags()) > 0 {
			for i, t := range req.GetSeasonTags() {
				if err := utils.ValidateString(t, 1, 300); err != nil {
					violations = append(violations, shared.FieldViolation(fmt.Sprintf("seasonTags >  tag at index [%d]", i), err))
				}
			}
		}
	} else {
		violations = append(violations, shared.FieldViolation("seasonTags", errors.New("you need to send the other tags in SeasonTags model")))
	}

	return violations
}
