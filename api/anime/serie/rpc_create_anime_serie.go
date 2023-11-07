package animeSerie

import (
	"context"
	"errors"
	"time"

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
			OriginalTitle:     req.AnimeSerie.GetOriginalTitle(),
			Aired:             req.AnimeSerie.GetAired().AsTime(),
			ReleaseYear:       req.AnimeSerie.GetReleaseYear(),
			Rating:            req.AnimeSerie.GetRating(),
			PortriatPoster:    req.AnimeSerie.GetPortriatPoster(),
			PortriatBlurHash:  req.AnimeSerie.GetPortriatBlurHash(),
			LandscapePoster:   req.AnimeSerie.GetLandscapePoster(),
			LandscapeBlurHash: req.AnimeSerie.GetLandscapeBlurHash(),
		},
		CreateAnimeResourceParams: db.CreateAnimeResourceParams{
			TmdbID:          req.Resources.GetTMDbID(),
			ImdbID:          req.Resources.GetIMDbID(),
			WikipediaUrl:    req.Resources.GetWikipediaUrl(),
			OfficialWebsite: req.Resources.GetOfficialWebsite(),
			CrunchyrollUrl:  req.Resources.GetCrunchyrollUrl(),
			SocialMedia:     req.Resources.GetSocialMedia(),
		},
	}

	data, err := server.gojo.CreateAnimeSerieTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie : %s", err)
	}

	res := &aspb.CreateAnimeSerieResponse{
		AnimeSerie: shared.ConvertAnimeSerie(data.AnimeSerie),
		Resources:  shared.ConvertAnimeResource(data.Resource),
	}
	return res, nil
}

func validateCreateAnimeSerieRequest(req *aspb.CreateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.GetAnimeSerie() != nil {
		if err := utils.ValidateString(req.GetAnimeSerie().GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shared.FieldViolation("originalTitle", err))
		}

		if err := utils.ValidateDate(req.GetAnimeSerie().GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, shared.FieldViolation("aired", err))
		}

		if err := utils.ValidateYear(req.GetAnimeSerie().GetReleaseYear()); err != nil {
			violations = append(violations, shared.FieldViolation("releaseYear", err))
		}

		if err := utils.ValidateString(req.GetAnimeSerie().GetRating(), 2, 30); err != nil {
			violations = append(violations, shared.FieldViolation("rating", err))
		}

		if err := utils.ValidateImage(req.GetAnimeSerie().GetPortriatPoster()); err != nil {
			violations = append(violations, shared.FieldViolation("portriatPoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeSerie().GetPortriatBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("portriatBlurHash", err))
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

	if req.GetResources() != nil {
		if err := utils.ValidateInt(int64(req.GetResources().GetTMDbID())); err != nil {
			violations = append(violations, shared.FieldViolation("TMDbID", err))
		}

		if err := utils.ValidateURL(req.GetResources().GetWikipediaUrl(), "wikipedia"); err != nil {
			violations = append(violations, shared.FieldViolation("wikipediaUrl", err))
		}

		if req.GetResources().GetCrunchyrollUrl() != "" {
			if err := utils.ValidateURL(req.GetResources().GetCrunchyrollUrl(), "crunchyroll"); err != nil {
				violations = append(violations, shared.FieldViolation("crunchyrollUrl", err))
			}
		}

		if req.GetResources().GetOfficialWebsite() != "" {
			if err := utils.ValidateURL(req.GetResources().GetOfficialWebsite(), ""); err != nil {
				violations = append(violations, shared.FieldViolation("officialWebsite", err))
			}
		}

		if req.GetResources().GetIMDbID() != "" {
			if err := utils.ValidateIMDbID(req.GetResources().GetIMDbID()); err != nil {
				violations = append(violations, shared.FieldViolation("IMDbID", err))
			}
		}

		if req.GetResources().GetSocialMedia() != nil {
			for _, s := range req.GetResources().SocialMedia {
				if err := utils.ValidateURL(s, ""); err != nil {
					violations = append(violations, shared.FieldViolation("socialMedia", err))
				}
			}
		}
	} else {
		violations = append(violations, shared.FieldViolation("resources", errors.New("you need to send the resources model")))
	}

	return violations
}
