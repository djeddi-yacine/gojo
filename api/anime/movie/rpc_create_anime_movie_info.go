package animeMovie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieInfo(ctx context.Context, req *ampb.CreateAnimeMovieInfoRequest) (*ampb.CreateAnimeMovieInfoResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime movie info")
	}

	if violations := validateCreateAnimeMovieInfoRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieInfoTxParams{
		AnimeID:   req.GetAnimeID(),
		GenreIDs:  req.GetGenres().GenreID,
		StudioIDs: req.GetStudios().StudioID,
	}

	data, err := server.gojo.CreateAnimeMovieInfoTx(ctx, arg)
	if err != nil {
		return nil, shared.ApiError("failed to add anime movie info", err)
	}

	res := &ampb.CreateAnimeMovieInfoResponse{
		AnimeMovie: shared.ConvertAnimeMovie(data.AnimeMovie),
		Genres:     shared.ConvertGenres(data.Genres),
		Studios:    shared.ConvertStudios(data.Studios),
	}

	return res, nil
}

func validateCreateAnimeMovieInfoRequest(req *ampb.CreateAnimeMovieInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	if req.Genres == nil && req.Studios == nil {
		violations = append(violations, shared.FieldViolation("studios,genres", errors.New("add at least one studio or genre")))
	} else {
		if req.Genres != nil {
			for _, g := range req.GetGenres().GenreID {
				if err := utils.ValidateInt(int64(g)); err != nil {
					violations = append(violations, shared.FieldViolation("genreID", err))
				}
			}

		}

		if req.Studios != nil {
			for _, s := range req.GetStudios().StudioID {
				if err := utils.ValidateInt(int64(s)); err != nil {
					violations = append(violations, shared.FieldViolation("studioID", err))
				}
			}

		}
	}

	return violations
}
