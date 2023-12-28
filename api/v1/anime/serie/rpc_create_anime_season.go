package asapiv1

import (
	"context"
	"errors"
	"time"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSeason(ctx context.Context, req *aspbv1.CreateAnimeSeasonRequest) (*aspbv1.CreateAnimeSeasonResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie season")
	}

	if violations := validateCreateAnimeSeasonRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
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

	arg := db.CreateAnimeSeasonTxParams{
		Season: db.CreateAnimeSeasonParams{
			AnimeID:             req.GetSeason().GetAnimeID(),
			SeasonOriginalTitle: req.GetSeason().GetSeasonOriginalTitle(),
			ReleaseYear:         req.GetSeason().GetReleaseYear(),
			Aired:               req.GetSeason().GetAired().AsTime(),
			Rating:              req.GetSeason().GetRating(),
			PortraitPoster:      req.GetSeason().GetPortraitPoster(),
			PortraitBlurHash:    req.GetSeason().GetPortraitBlurHash(),
		},
		SeasonMetas: DBSM,
	}

	data, err := server.gojo.CreateAnimeSeasonTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, shv1.ApiError("failed to create anime serie season", err)
	}

	var PBSM = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeSeasonMetas))

	for i, am := range data.AnimeSeasonMetas {
		PBSM[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(am.Meta),
			LanguageID: am.LanguageID,
			CreatedAt:  timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspbv1.CreateAnimeSeasonResponse{
		Season:      convertAnimeSeason(data.AnimeSeason),
		SeasonMetas: PBSM,
	}
	return res, nil
}

func validateCreateAnimeSeasonRequest(req *aspbv1.CreateAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Season != nil {
		if err := utils.ValidateInt(req.GetSeason().GetAnimeID()); err != nil {
			violations = append(violations, shv1.FieldViolation("animeID", err))
		}

		if err := utils.ValidateString(req.GetSeason().GetSeasonOriginalTitle(), 1, 300); err != nil {
			violations = append(violations, shv1.FieldViolation("seasonOriginalTitle", err))
		}

		if err := utils.ValidateInt(int64(req.GetSeason().GetReleaseYear())); err != nil {
			violations = append(violations, shv1.FieldViolation("releaseYear", err))
		}

		if err := utils.ValidateImage(req.GetSeason().GetPortraitPoster()); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitPoster", err))
		}

		if err := utils.ValidateString(req.GetSeason().GetPortraitBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitBlurHash", err))
		}

		if err := utils.ValidateDate(req.GetSeason().GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, shv1.FieldViolation("aired", err))
		}

		if err := utils.ValidateString(req.GetSeason().GetRating(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("rating", err))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("season", errors.New("season :you need to send the season model")))
	}

	if req.SeasonMetas != nil {
		for _, am := range req.SeasonMetas {
			if err := utils.ValidateInt(int64(am.GetLanguageID())); err != nil {
				violations = append(violations, shv1.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shv1.FieldViolation("title", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shv1.FieldViolation("overview", err))
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("seasonMetas", errors.New("seasonMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
