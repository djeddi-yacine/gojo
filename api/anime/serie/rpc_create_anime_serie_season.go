package animeSerie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSerieSeason(ctx context.Context, req *aspb.CreateAnimeSerieSeasonRequest) (*aspb.CreateAnimeSerieSeasonResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie season")
	}

	if violations := validateCreateAnimeSerieSeasonRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBSM = make([]db.AnimeMetaTxParam, len(req.SeasonMetas))
	for i, am := range req.SeasonMetas {
		DBSM[i] = db.AnimeMetaTxParam{
			LanguageID: am.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    am.GetMeta().GetTitle(),
				Overview: am.GetMeta().GetOverview(),
			},
		}
	}

	arg := db.CreateAnimeSerieSeasonTxParams{
		Season: db.CreateAnimeSerieSeasonParams{
			AnimeID:          req.GetSeason().GetAnimeID(),
			ReleaseYear:      req.GetSeason().GetReleaseYear(),
			Aired:            req.GetSeason().GetAired().AsTime(),
			Rating:           req.GetSeason().GetRating(),
			PortriatPoster:   req.GetSeason().GetPortriatPoster(),
			PortriatBlurHash: req.GetSeason().GetPortriatBlurHash(),
		},
		SeasonMetas: DBSM,
	}

	data, err := server.gojo.CreateAnimeSerieSeasonTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie season : %s", err)
	}

	var PBSM = make([]*nfpb.AnimeMetaResponse, len(data.AnimeSerieSeasonMetas))

	for i, am := range data.AnimeSerieSeasonMetas {
		PBSM[i] = &nfpb.AnimeMetaResponse{
			Meta:       shared.ConvertMeta(am.Meta),
			LanguageID: am.LanguageID,
			CreatedAt:  timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspb.CreateAnimeSerieSeasonResponse{
		Season:      shared.ConvertAnimeSerieSeason(data.AnimeSerieSeason),
		SeasonMetas: PBSM,
	}
	return res, nil
}

func validateCreateAnimeSerieSeasonRequest(req *aspb.CreateAnimeSerieSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Season != nil {
		if err := utils.ValidateInt(req.GetSeason().GetAnimeID()); err != nil {
			violations = append(violations, shared.FieldViolation("animeID", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeason().GetReleaseYear())); err != nil {
			violations = append(violations, shared.FieldViolation("releaseYear", err))
		}

		if err := utils.ValidateImage(req.GetSeason().GetPortriatPoster()); err != nil {
			violations = append(violations, shared.FieldViolation("portriatPoster", err))
		}

		if err := utils.ValidateString(req.GetSeason().GetPortriatBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("portriatBlurHash", err))
		}

		if err := utils.ValidateDate(req.GetSeason().GetAired().String()); err != nil {
			violations = append(violations, shared.FieldViolation("aired", err))
		}

		if err := utils.ValidateString(req.GetSeason().GetRating(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("rating", err))
		}

	} else {
		violations = append(violations, shared.FieldViolation("season", errors.New("season :you need to send the season model")))
	}

	if req.SeasonMetas != nil {
		for _, am := range req.SeasonMetas {
			if err := utils.ValidateInt(int64(am.GetLanguageID())); err != nil {
				violations = append(violations, shared.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shared.FieldViolation("title", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shared.FieldViolation("overview", err))
			}
		}
	} else {
		violations = append(violations, shared.FieldViolation("seasonMetas", errors.New("seasonMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
