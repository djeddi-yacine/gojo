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

func (server *AnimeSerieServer) CreateAnimeSerie(ctx context.Context, req *aspb.CreateAnimeSerieRequest) (*aspb.CreateAnimeSerieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie")
	}

	if violations := validateCreateAnimeSerieRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
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

	data, err := server.gojo.CreateAnimeSerieTx(ctx, arg)
	if err != nil {
		return nil, shared.ApiError("failed to create anime serie", err)
	}

	res := &aspb.CreateAnimeSerieResponse{
		AnimeSerie: shared.ConvertAnimeSerie(data.AnimeSerie),
		AnimeLinks: shared.ConvertAnimeLink(data.AnimeLink),
	}
	return res, nil
}

func validateCreateAnimeSerieRequest(req *aspb.CreateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.AnimeSerie != nil {
		if err := utils.ValidateString(req.GetAnimeSerie().GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shared.FieldViolation("originalTitle", err))
		}

		if err := utils.ValidateYear(req.GetAnimeSerie().GetFirstYear()); err != nil {
			violations = append(violations, shared.FieldViolation("firstYear", err))
		}

		if err := utils.ValidateYear(req.GetAnimeSerie().GetLastYear()); err != nil {
			violations = append(violations, shared.FieldViolation("lastYear", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeSerie().GetMalID())); err != nil {
			violations = append(violations, shared.FieldViolation("malID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeSerie().GetTmdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tmdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeSerie().GetTvdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tvdbID", err))
		}

		if err := utils.ValidateImage(req.GetAnimeSerie().GetPortraitPoster()); err != nil {
			violations = append(violations, shared.FieldViolation("portraitPoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeSerie().GetPortraitBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("portraitBlurHash", err))
		}

		if err := utils.ValidateImage(req.GetAnimeSerie().GetLandscapePoster()); err != nil {
			violations = append(violations, shared.FieldViolation("landscapePoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeSerie().GetLandscapeBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("landscapeBlurHash", err))
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeSerie", errors.New("you need to send the animeSerie model")))
	}

	if req.AnimeLinks != nil {
		if err := utils.ValidateURL(req.GetAnimeLinks().GetOfficialWebsite(), ""); err != nil {
			violations = append(violations, shared.FieldViolation("officialWebsite", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetCrunchyrollUrl(), "crunchyroll"); err != nil {
			violations = append(violations, shared.FieldViolation("crunchyrollUrl", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetWikipediaUrl(), "wikipedia"); err != nil {
			violations = append(violations, shared.FieldViolation("wikipediaUrl", err))
		}

		if len(req.GetAnimeLinks().GetSocialMedia()) > 0 {
			for _, l := range req.GetAnimeLinks().GetSocialMedia() {
				if err := utils.ValidateURL(l, ""); err != nil {
					violations = append(violations, shared.FieldViolation("socialMedia", err))
				}
			}
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeLinks", errors.New("you need to send the AnimeLinks model")))
	}

	return violations
}
