package animeMovie

import (
	"context"
	"errors"
	"time"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovie(ctx context.Context, req *ampb.CreateAnimeMovieRequest) (*ampb.CreateAnimeMovieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie")
	}

	if violations := validateCreateAnimeMovieRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieTxParams{
		CreateAnimeMovieParams: db.CreateAnimeMovieParams{
			OriginalTitle:     req.GetAnimeMovie().GetOriginalTitle(),
			Aired:             req.GetAnimeMovie().GetAired().AsTime(),
			ReleaseYear:       req.GetAnimeMovie().GetReleaseYear(),
			Rating:            req.GetAnimeMovie().GetRating(),
			Duration:          req.GetAnimeMovie().GetDuration().AsDuration(),
			PortraitPoster:    req.GetAnimeMovie().GetPortraitPoster(),
			PortraitBlurHash:  req.GetAnimeMovie().GetPortraitBlurHash(),
			LandscapePoster:   req.GetAnimeMovie().GetLandscapePoster(),
			LandscapeBlurHash: req.GetAnimeMovie().GetLandscapeBlurHash(),
		},
		CreateAnimeResourceParams: db.CreateAnimeResourceParams{
			TvdbID:        req.GetAnimeResources().GetTvdbID(),
			TmdbID:        req.GetAnimeResources().GetTmdbID(),
			ImdbID:        req.GetAnimeResources().GetImdbID(),
			LivechartID:   req.GetAnimeResources().GetLivechartID(),
			AnimePlanetID: req.GetAnimeResources().GetAnimePlanetID(),
			AnisearchID:   req.GetAnimeResources().GetAnisearchID(),
			AnidbID:       req.GetAnimeResources().GetAnidbID(),
			KitsuID:       req.GetAnimeResources().GetKitsuID(),
			MalID:         req.GetAnimeResources().GetMalID(),
			NotifyMoeID:   req.GetAnimeResources().GetNotifyMoeID(),
			AnilistID:     req.GetAnimeResources().GetAnilistID(),
		},
		CreateAnimeLinkParams: db.CreateAnimeLinkParams{
			OfficialWebsite: req.GetAnimeLinks().GetOfficialWebsite(),
			WikipediaUrl:    req.GetAnimeLinks().GetWikipediaUrl(),
			CrunchyrollUrl:  req.GetAnimeLinks().GetCrunchyrollUrl(),
			SocialMedia:     req.GetAnimeLinks().GetSocialMedia(),
		},
	}

	data, err := server.gojo.CreateAnimeMovieTx(ctx, arg)
	if err != nil {
		return nil, shared.ApiError("cannot create anime movie", err)
	}

	res := &ampb.CreateAnimeMovieResponse{
		AnimeMovie:     shared.ConvertAnimeMovie(data.AnimeMovie),
		AnimeResources: shared.ConvertAnimeResource(data.AnimeResource),
		AnimeLinks:     shared.ConvertAnimeLink(data.AnimeLink),
	}
	return res, nil
}

func validateCreateAnimeMovieRequest(req *ampb.CreateAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.AnimeMovie != nil {
		if err := utils.ValidateString(req.GetAnimeMovie().GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shared.FieldViolation("originalTitle", err))
		}

		if err := utils.ValidateDate(req.GetAnimeMovie().GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, shared.FieldViolation("aired", err))
		}

		if err := utils.ValidateYear(req.GetAnimeMovie().GetReleaseYear()); err != nil {
			violations = append(violations, shared.FieldViolation("releaseYear", err))
		}

		if err := utils.ValidateString(req.GetAnimeMovie().GetRating(), 2, 30); err != nil {
			violations = append(violations, shared.FieldViolation("rating", err))
		}

		if err := utils.ValidateDuration(req.GetAnimeMovie().GetDuration().AsDuration().String()); err != nil {
			violations = append(violations, shared.FieldViolation("duration", err))
		}

		if err := utils.ValidateImage(req.GetAnimeMovie().GetPortraitPoster()); err != nil {
			violations = append(violations, shared.FieldViolation("portraitPoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeMovie().GetPortraitBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("portraitBlurHash", err))
		}

		if err := utils.ValidateImage(req.GetAnimeMovie().GetLandscapePoster()); err != nil {
			violations = append(violations, shared.FieldViolation("landscapePoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeMovie().GetLandscapeBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("landscapeBlurHash", err))
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeMovie", errors.New("you need to send the animeMovie model")))
	}

	if req.AnimeResources != nil {
		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetTvdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tvdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetTmdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tmdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetLivechartID())); err != nil {
			violations = append(violations, shared.FieldViolation("livechartID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnidbID())); err != nil {
			violations = append(violations, shared.FieldViolation("anidbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnisearchID())); err != nil {
			violations = append(violations, shared.FieldViolation("anisearchID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetKitsuID())); err != nil {
			violations = append(violations, shared.FieldViolation("kitsuID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetMalID())); err != nil {
			violations = append(violations, shared.FieldViolation("malID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnilistID())); err != nil {
			violations = append(violations, shared.FieldViolation("anilistID", err))
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeResources", errors.New("you need to send the AnimeResources model")))
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
