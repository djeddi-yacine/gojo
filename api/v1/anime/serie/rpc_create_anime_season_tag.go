package asapiv1

import (
	"context"
	"errors"
	"fmt"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSeasonTag(ctx context.Context, req *aspbv1.CreateAnimeSeasonTagRequest) (*aspbv1.CreateAnimeSeasonTagResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season tags")
	}

	if violations := validateCreateAnimeSeasonTagRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
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
		return nil, shv1.ApiError("failed to create anime season tags", err)
	}

	res := &aspbv1.CreateAnimeSeasonTagResponse{
		AnimeSeason: convertAnimeSeason(data.AnimeSeason),
		SeasonTags:  aapiv1.ConvertAnimeTags(data.SeasonTags),
	}

	return res, nil
}

func validateCreateAnimeSeasonTagRequest(req *aspbv1.CreateAnimeSeasonTagRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.SeasonTags != nil {
		if len(req.GetSeasonTags()) > 0 {
			for i, t := range req.GetSeasonTags() {
				if err := utils.ValidateString(t, 1, 300); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("seasonTags >  tag at index [%d]", i), err))
				}
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("seasonTags", errors.New("you need to send the other tags in SeasonTags model")))
	}

	return violations
}
