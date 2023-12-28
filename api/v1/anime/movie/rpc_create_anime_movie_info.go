package amapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieInfo(ctx context.Context, req *ampbv1.CreateAnimeMovieInfoRequest) (*ampbv1.CreateAnimeMovieInfoResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime movie info")
	}

	if violations := validateCreateAnimeMovieInfoRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieInfoTxParams{
		AnimeID:   req.GetAnimeID(),
		GenreIDs:  req.GetGenreIDs(),
		StudioIDs: req.GetStudioIDs(),
	}

	data, err := server.gojo.CreateAnimeMovieInfoTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to add anime movie info", err)
	}

	res := &ampbv1.CreateAnimeMovieInfoResponse{
		AnimeMovie: convertAnimeMovie(data.AnimeMovie),
		Genres:     shv1.ConvertGenres(data.Genres),
		Studios:    shv1.ConvertStudios(data.Studios),
	}

	return res, nil
}

func validateCreateAnimeMovieInfoRequest(req *ampbv1.CreateAnimeMovieInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.GenreIDs == nil && req.StudioIDs == nil {
		violations = append(violations, shv1.FieldViolation("studioIDs,genreIDs", errors.New("add at least one ID for studio or genre")))
	} else {
		if req.GenreIDs != nil {
			for i, g := range req.GetGenreIDs() {
				if err := utils.ValidateInt(int64(g)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("genreIDs at index [%d]", i), err))
				}
			}
		}

		if req.StudioIDs != nil {
			for i, s := range req.GetStudioIDs() {
				if err := utils.ValidateInt(int64(s)); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("studioIDs at index [%d]", i), err))
				}
			}

		}
	}

	return violations
}
