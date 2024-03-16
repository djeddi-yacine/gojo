package asapiv1

import (
	"context"
	"errors"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
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

func (server *AnimeSerieServer) CreateAnimeSerie(ctx context.Context, req *aspbv1.CreateAnimeSerieRequest) (*aspbv1.CreateAnimeSerieResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie")
	}

	if violations := validateCreateAnimeSerieRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSerieTxParams{
		CreateAnimeSerieParams: db.CreateAnimeSerieParams{
			OriginalTitle:     req.GetAnimeSerie().GetOriginalTitle(),
			FirstYear:         req.GetAnimeSerie().GetFirstYear(),
			LastYear:          req.GetAnimeSerie().GetLastYear(),
			MalID:             req.GetAnimeSerie().GetMalID(),
			TvdbID:            req.GetAnimeSerie().GetTvdbID(),
			TmdbID:            req.GetAnimeSerie().GetTmdbID(),
			PortraitPoster:    req.GetAnimeSerie().GetPortraitPoster(),
			PortraitBlurHash:  req.GetAnimeSerie().GetPortraitBlurHash(),
			LandscapePoster:   req.GetAnimeSerie().GetLandscapePoster(),
			LandscapeBlurHash: req.GetAnimeSerie().GetLandscapeBlurHash(),
		},
		CreateAnimeLinkParams: db.CreateAnimeLinkParams{
			OfficialWebsite: req.GetAnimeLinks().GetOfficialWebsite(),
			WikipediaUrl:    req.GetAnimeLinks().GetWikipediaUrl(),
			CrunchyrollUrl:  req.GetAnimeLinks().GetCrunchyrollUrl(),
			SocialMedia:     req.GetAnimeLinks().GetSocialMedia(),
		},
	}

	arg.CreateAnimeMetasParams = make([]db.AnimeMetaTxParam, len(req.AnimeMetas))
	for i, v := range req.AnimeMetas {
		arg.CreateAnimeMetasParams[i] = db.AnimeMetaTxParam{
			LanguageID: v.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    v.GetMeta().GetTitle(),
				Overview: v.GetMeta().GetOverview(),
			},
		}
	}

	data, err := server.gojo.CreateAnimeSerieTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime serie", err)
	}

	res := &aspbv1.CreateAnimeSerieResponse{
		AnimeSerie: server.convertAnimeSerie(data.AnimeSerie),
		AnimeLinks: av1.ConvertAnimeLink(data.AnimeLink),
	}

	titles := make([]string, len(data.AnimeMetas))
	res.AnimeMetas = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeMetas))
	for i, v := range data.AnimeMetas {
		res.AnimeMetas[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(v.Meta),
			LanguageID: v.LanguageID,
			CreatedAt:  timestamppb.New(v.Meta.CreatedAt),
		}
		titles[i] = v.Meta.Title
	}

	return res, nil
}

func validateCreateAnimeSerieRequest(req *aspbv1.CreateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.AnimeSerie != nil {
		if err := utils.ValidateString(req.GetAnimeSerie().GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shv1.FieldViolation("originalTitle", err))
		}

		if err := utils.ValidateYear(req.GetAnimeSerie().GetFirstYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("firstYear", err))
		}

		if err := utils.ValidateYear(req.GetAnimeSerie().GetLastYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("lastYear", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeSerie().GetMalID())); err != nil {
			violations = append(violations, shv1.FieldViolation("malID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeSerie().GetTmdbID())); err != nil {
			violations = append(violations, shv1.FieldViolation("tmdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeSerie().GetTvdbID())); err != nil {
			violations = append(violations, shv1.FieldViolation("tvdbID", err))
		}

		if err := utils.ValidateImage(req.GetAnimeSerie().GetPortraitPoster()); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitPoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeSerie().GetPortraitBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitBlurHash", err))
		}

		if err := utils.ValidateImage(req.GetAnimeSerie().GetLandscapePoster()); err != nil {
			violations = append(violations, shv1.FieldViolation("landscapePoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeSerie().GetLandscapeBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("landscapeBlurHash", err))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeSerie", errors.New("you need to send the animeSerie model")))
	}

	if req.AnimeLinks != nil {
		if err := utils.ValidateURL(req.GetAnimeLinks().GetOfficialWebsite(), ""); err != nil {
			violations = append(violations, shv1.FieldViolation("officialWebsite", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetCrunchyrollUrl(), "crunchyroll"); err != nil {
			violations = append(violations, shv1.FieldViolation("crunchyrollUrl", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetWikipediaUrl(), "wikipedia"); err != nil {
			violations = append(violations, shv1.FieldViolation("wikipediaUrl", err))
		}

		if len(req.GetAnimeLinks().GetSocialMedia()) > 0 {
			for _, v := range req.GetAnimeLinks().GetSocialMedia() {
				if err := utils.ValidateURL(v, ""); err != nil {
					violations = append(violations, shv1.FieldViolation("socialMedia", err))
				}
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeLinks", errors.New("you need to send the AnimeLinks model")))
	}

	if req.AnimeMetas != nil {
		for _, v := range req.AnimeMetas {
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
		violations = append(violations, shv1.FieldViolation("animeMetas", errors.New("give at least one metadata")))
	}

	return violations
}
