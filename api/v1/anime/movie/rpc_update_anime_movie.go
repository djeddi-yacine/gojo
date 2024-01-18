package amapiv1

import (
	"context"
	"fmt"
	"time"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) UpdateAnimeMovie(ctx context.Context, req *ampbv1.UpdateAnimeMovieRequest) (*ampbv1.UpdateAnimeMovieResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update anime movie")
	}

	if violations := validateUpdateAnimeMovieRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.UpdateAnimeMovieParams{
		ID: req.GetAnimeID(),
		OriginalTitle: pgtype.Text{
			String: req.GetOriginalTitle(),
			Valid:  req.OriginalTitle != nil,
		},
		Aired: pgtype.Timestamptz{
			Time:  req.GetAired().AsTime(),
			Valid: req.Aired != nil,
		},
		ReleaseYear: pgtype.Int4{
			Int32: req.GetReleaseYear(),
			Valid: req.ReleaseYear != nil,
		},
		Rating: pgtype.Text{
			String: req.GetRating(),
			Valid:  req.Rating != nil,
		},
		Duration: pgtype.Interval{
			Microseconds: req.GetDuration().AsDuration().Microseconds(),
			Valid:        req.Duration != nil,
		},
		PortraitPoster: pgtype.Text{
			String: req.GetPortraitPoster(),
			Valid:  req.PortraitPoster != nil,
		},
		PortraitBlurHash: pgtype.Text{
			String: req.GetPortraitBlurHash(),
			Valid:  req.PortraitBlurHash != nil,
		},
		LandscapePoster: pgtype.Text{
			String: req.GetLandscapePoster(),
			Valid:  req.LandscapePoster != nil,
		},
		LandscapeBlurHash: pgtype.Text{
			String: req.GetLandscapeBlurHash(),
			Valid:  req.LandscapeBlurHash != nil,
		},
	}

	anime, err := server.gojo.UpdateAnimeMovie(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to update anime movie", err)
	}

	res := &ampbv1.UpdateAnimeMovieResponse{
		AnimeMovie: convertAnimeMovie(anime),
	}

	return res, nil
}

func validateUpdateAnimeMovieRequest(req *ampbv1.UpdateAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.OriginalTitle != nil {
		if err := utils.ValidateString(req.GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shv1.FieldViolation("originalTitle", err))
		}
	}

	if req.Aired != nil {
		if err := utils.ValidateDate(req.GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, shv1.FieldViolation("aired", err))
		}
	}

	if req.ReleaseYear != nil {
		if err := utils.ValidateYear(req.GetReleaseYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("releaseYear", err))
		}
	}

	if req.Rating != nil {
		if err := utils.ValidateString(req.GetRating(), 2, 30); err != nil {
			violations = append(violations, shv1.FieldViolation("rating", err))
		}
	}

	if req.Duration != nil {
		fmt.Println(req.GetDuration().String())
		if err := utils.ValidateDuration(req.GetDuration().AsDuration().String()); err != nil {
			violations = append(violations, shv1.FieldViolation("duration", err))
		}
	}

	if req.PortraitPoster != nil {
		if err := utils.ValidateImage(req.GetPortraitPoster()); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitPoster", err))
		}
	}

	if req.PortraitBlurHash != nil {
		if err := utils.ValidateString(req.GetPortraitBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitBlurHash", err))
		}
	}

	if req.LandscapePoster != nil {
		if err := utils.ValidateImage(req.GetLandscapePoster()); err != nil {
			violations = append(violations, shv1.FieldViolation("landscapePoster", err))
		}
	}

	if req.LandscapeBlurHash != nil {
		if err := utils.ValidateString(req.GetLandscapeBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("landscapeBlurHash", err))
		}
	}

	return violations
}
