package animeSerie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSeasonResource(ctx context.Context, req *aspb.CreateAnimeSeasonResourceRequest) (*aspb.CreateAnimeSeasonResourceResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie resource")
	}

	if violations := validateCreateAnimeSerieResourceRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
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
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie resource: %s", err)
	}

	res := &aspb.CreateAnimeSeasonResourceResponse{
		AnimeSeason:     shared.ConvertAnimeSerieSeason(data.AnimeSeason),
		SeasonResources: shared.ConvertAnimeResource(data.AnimeResource),
	}
	return res, nil
}

func validateCreateAnimeSerieResourceRequest(req *aspb.CreateAnimeSeasonResourceRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shared.FieldViolation("ID", err))
	}

	if req.SeasonResources != nil {
		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetTvdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tvdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetTmdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tmdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetLivechartID())); err != nil {
			violations = append(violations, shared.FieldViolation("livechartID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetAnidbID())); err != nil {
			violations = append(violations, shared.FieldViolation("anidbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetAnisearchID())); err != nil {
			violations = append(violations, shared.FieldViolation("anisearchID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetKitsuID())); err != nil {
			violations = append(violations, shared.FieldViolation("kitsuID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetMalID())); err != nil {
			violations = append(violations, shared.FieldViolation("malID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeasonResources().GetAnilistID())); err != nil {
			violations = append(violations, shared.FieldViolation("anilistID", err))
		}

	} else {
		violations = append(violations, shared.FieldViolation("SeasonResources", errors.New("you need to send the SeasonResources model")))
	}

	return violations
}
