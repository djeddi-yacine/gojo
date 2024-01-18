package asapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) UpdateAnimeSerie(ctx context.Context, req *aspbv1.UpdateAnimeSerieRequest) (*aspbv1.UpdateAnimeSerieResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update anime serie")
	}

	if violations := validateUpdateAnimeSerieRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.UpdateAnimeSerieParams{
		ID: req.GetAnimeID(),
		OriginalTitle: pgtype.Text{
			String: req.GetOriginalTitle(),
			Valid:  req.OriginalTitle != nil,
		},
		FirstYear: pgtype.Int4{
			Int32: req.GetFirstYear(),
			Valid: req.FirstYear != nil,
		},
		LastYear: pgtype.Int4{
			Int32: req.GetLastYear(),
			Valid: req.LastYear != nil,
		},
		MalID: pgtype.Int4{
			Int32: req.GetMalID(),
			Valid: req.MalID != nil,
		},
		TvdbID: pgtype.Int4{
			Int32: req.GetTvdbID(),
			Valid: req.TvdbID != nil,
		},
		TmdbID: pgtype.Int4{
			Int32: req.GetTmdbID(),
			Valid: req.LastYear != nil,
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

	anime, err := server.gojo.UpdateAnimeSerie(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to update anime serie", err)
	}

	res := &aspbv1.UpdateAnimeSerieResponse{
		AnimeSerie: convertAnimeSerie(anime),
	}

	return res, nil
}

func validateUpdateAnimeSerieRequest(req *aspbv1.UpdateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.OriginalTitle != nil {
		if err := utils.ValidateString(req.GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shv1.FieldViolation("originalTitle", err))
		}
	}

	if req.FirstYear != nil {
		if err := utils.ValidateYear(req.GetFirstYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("firstYear", err))
		}
	}

	if req.LastYear != nil {
		if err := utils.ValidateYear(req.GetLastYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("lastYear", err))
		}
	}

	if req.MalID != nil {
		if err := utils.ValidateYear(req.GetMalID()); err != nil {
			violations = append(violations, shv1.FieldViolation("malID", err))
		}
	}

	if req.TvdbID != nil {
		if err := utils.ValidateYear(req.GetTvdbID()); err != nil {
			violations = append(violations, shv1.FieldViolation("tvdbID", err))
		}
	}

	if req.TmdbID != nil {
		if err := utils.ValidateYear(req.GetTmdbID()); err != nil {
			violations = append(violations, shv1.FieldViolation("tmdbID", err))
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
