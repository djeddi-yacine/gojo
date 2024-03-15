package amapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeMovieServer) GetAllAnimeMovies(ctx context.Context, req *ampbv1.GetAllAnimeMoviesRequest) (*ampbv1.GetAllAnimeMoviesResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAllAnimeMoviesRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.ListAnimeMoviesParams{
		ReleaseYear: req.GetYear(),
		Limit:       req.GetPageSize(),
		Offset:      (req.GetPageNumber() - 1) * req.GetPageSize(),
	}

	data, err := server.gojo.ListAnimeMovies(ctx, arg)
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, nil
		}
		return nil, shv1.ApiError("failed to list all anime movies", err)
	}

	res := &ampbv1.GetAllAnimeMoviesResponse{}

	res.AnimeMovies = make([]*ampbv1.AnimeMovieResponse, len(data))
	for i, v := range data {
		res.AnimeMovies[i] = server.convertAnimeMovie(v)
	}

	return res, nil
}

func validateGetAllAnimeMoviesRequest(req *ampbv1.GetAllAnimeMoviesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	if req.GetYear() != 0 {
		if err := utils.ValidateYear(req.GetYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("year", err))
		}
	}

	return violations
}
