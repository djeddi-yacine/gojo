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
			OriginalTitle: req.AnimeMovie.GetOriginalTitle(),
			Aired:         req.AnimeMovie.GetAired().AsTime(),
			ReleaseYear:   req.AnimeMovie.GetReleaseYear(),
			Rating:        req.AnimeMovie.GetRating(),
			Duration:      req.AnimeMovie.GetDuration().AsDuration(),
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

	data, err := server.gojo.CreateAnimeMovieTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime movie : %s", err)
	}

	res := &ampb.CreateAnimeMovieResponse{
		AnimeMovie: shared.ConvertAnimeMovie(data.AnimeMovie),
		Resources:  shared.ConvertAnimeResource(data.Resource),
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
	} else {
		violations = append(violations, shared.FieldViolation("animeMovie", errors.New("you need to send the animeMovie model")))
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
