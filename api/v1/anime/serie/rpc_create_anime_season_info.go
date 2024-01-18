package asapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSeasonInfo(ctx context.Context, req *aspbv1.CreateAnimeSeasonInfoRequest) (*aspbv1.CreateAnimeSeasonInfoResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime movie info")
	}

	if violations := validateCreateAnimeSeasonInfoRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	data, err := server.gojo.CreateAnimeSeasonInfoTx(ctx, db.CreateAnimeSeasonInfoTxParams{
		SeasonID:  req.GetSeasonID(),
		GenreIDs:  req.GetGenreIDs(),
		StudioIDs: req.GetStudioIDs(),
	})
	if err != nil {
		return nil, shv1.ApiError("failed to add anime serie info", err)
	}

	res := &aspbv1.CreateAnimeSeasonInfoResponse{
		AnimeSeason: convertAnimeSeason(data.AnimeSeason),
		Genres:      shv1.ConvertGenres(data.Genres),
		Studios:     shv1.ConvertStudios(data.Studios),
	}

	return res, nil
}

func validateCreateAnimeSeasonInfoRequest(req *aspbv1.CreateAnimeSeasonInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.GenreIDs == nil && req.StudioIDs == nil {
		violations = append(violations, shv1.FieldViolation("studioIDs,genreIDs", errors.New("add at least one ID for studio or genre")))
	} else {
		if req.GenreIDs != nil {
			for i, v := range req.GetGenreIDs() {
				if err := utils.ValidateInt(int64(v)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("genreIDs at index [%d]", i), err))
				}
			}
		}

		if req.StudioIDs != nil {
			for i, v := range req.GetStudioIDs() {
				if err := utils.ValidateInt(int64(v)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("studioIDs at index [%d]", i), err))
				}
			}

		}
	}

	return violations
}
