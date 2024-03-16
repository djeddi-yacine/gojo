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

	arg := db.CreateAnimeSeasonTxParams{
		Season: db.CreateAnimeSeasonParams{
			AnimeID:             req.GetSeason().GetAnimeID(),
			SeasonOriginalTitle: req.GetSeason().GetSeasonOriginalTitle(),
			ReleaseYear:         req.GetSeason().GetReleaseYear(),
			Aired:               req.GetSeason().GetAired().AsTime(),
			Rating:              req.GetSeason().GetRating(),
			PortraitPoster:      req.GetSeason().GetPortraitPoster(),
			PortraitBlurHash:    req.GetSeason().GetPortraitBlurHash(),
			ShowType:            req.GetSeason().GetShowType(),
		},
	}

	arg.SeasonMetas = make([]db.AnimeMetaTxParam, len(req.SeasonMetas))
	for i, v := range req.SeasonMetas {
		arg.SeasonMetas[i] = db.AnimeMetaTxParam{
			LanguageID: v.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    v.GetMeta().GetTitle(),
				Overview: v.GetMeta().GetOverview(),
			},
		}
	}

	data, err := server.gojo.CreateAnimeSeasonTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime serie season", err)
	}

	res := &aspbv1.CreateAnimeSeasonResponse{
		Season: server.convertAnimeSeason(data.AnimeSeason),
	}

	titles := make([]string, len(data.AnimeSeasonMetas))
	res.SeasonMetas = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeSeasonMetas))
	for i, v := range data.AnimeSeasonMetas {
		res.SeasonMetas[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(v.Meta),
			LanguageID: v.LanguageID,
			CreatedAt:  timestamppb.New(v.Meta.CreatedAt),
		}

		titles[i] = v.Meta.Title
	}

	server.meilisearch.AddDocuments(&utils.Document{
		ID:     data.AnimeSeason.ID,
		Titles: utils.RemoveDuplicatesTitles(titles),
	})

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

		if err := utils.ValidateShow(req.GetSeason().GetShowType()); err != nil {
			violations = append(violations, shv1.FieldViolation("showType", err))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("season", errors.New("season :you need to send the season model")))
	}

	if req.SeasonMetas != nil {
		for _, v := range req.SeasonMetas {
			if err := utils.ValidateInt(int64(v.GetLanguageID())); err != nil {
				violations = append(violations, shv1.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(v.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shv1.FieldViolation("title", err))
			}

			if err := utils.ValidateString(v.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shv1.FieldViolation("overview", err))
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("seasonMetas", errors.New("seasonMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
