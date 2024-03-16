package asapiv1

import (
	"context"
	"errors"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSeasonResource(ctx context.Context, req *aspbv1.CreateAnimeSeasonResourceRequest) (*aspbv1.CreateAnimeSeasonResourceResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie resource")
	}

	if violations := validateCreateAnimeSerieResourceRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSeasonResourceTxParams{
		SeasonID: req.GetSeasonID(),
		CreateAnimeResourceParams: db.CreateAnimeResourceParams{
			TvdbID:        req.GetSeasonResources().GetTvdbID(),
			TmdbID:        req.GetSeasonResources().GetTmdbID(),
			ImdbID:        req.GetSeasonResources().GetImdbID(),
			LivechartID:   req.GetSeasonResources().GetLivechartID(),
			AnimePlanetID: req.GetSeasonResources().GetAnimePlanetID(),
			AnisearchID:   req.GetSeasonResources().GetAnisearchID(),
			AnidbID:       req.GetSeasonResources().GetAnidbID(),
			KitsuID:       req.GetSeasonResources().GetKitsuID(),
			MalID:         req.GetSeasonResources().GetMalID(),
			NotifyMoeID:   req.GetSeasonResources().GetNotifyMoeID(),
			AnilistID:     req.GetSeasonResources().GetAnilistID(),
		},
	}

	data, err := server.gojo.CreateAnimeSeasonResourceTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime serie resources", err)
	}

	res := &aspbv1.CreateAnimeSeasonResourceResponse{
		AnimeSeason:     server.convertAnimeSeason(data.AnimeSeason),
		SeasonResources: av1.ConvertAnimeResource(data.AnimeResource),
	}
	return res, nil
}

func validateCreateAnimeSerieResourceRequest(req *aspbv1.CreateAnimeSeasonResourceRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.SeasonResources != nil {
		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetTvdbID() + 1)); err != nil {
			violations = append(violations, shv1.FieldViolation("tvdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetTmdbID() + 1)); err != nil {
			violations = append(violations, shv1.FieldViolation("tmdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetLivechartID())); err != nil {
			violations = append(violations, shv1.FieldViolation("livechartID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetAnidbID())); err != nil {
			violations = append(violations, shv1.FieldViolation("anidbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetAnisearchID())); err != nil {
			violations = append(violations, shv1.FieldViolation("anisearchID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetKitsuID())); err != nil {
			violations = append(violations, shv1.FieldViolation("kitsuID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetMalID())); err != nil {
			violations = append(violations, shv1.FieldViolation("malID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetAnilistID())); err != nil {
			violations = append(violations, shv1.FieldViolation("anilistID", err))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("SeasonResources", errors.New("you need to send the SeasonResources model")))
	}

	return violations
}
