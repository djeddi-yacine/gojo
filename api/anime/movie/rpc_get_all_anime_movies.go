package animeMovie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) GetAllAnimeMovies(ctx context.Context, req *ampb.GetAllAnimeMoviesRequest) (*ampb.GetAllAnimeMoviesResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all anime movies")
	}

	violations := validateGetAllAnimeMoviesRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.ListAnimeMoviesParams{
		ReleaseYear: req.GetYear(),
		Limit:       req.GetPageSize(),
		Offset:      (req.GetPageNumber() - 1) * req.GetPageSize(),
	}
	DBAnimeMovies, err := server.gojo.ListAnimeMovies(ctx, arg)
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, nil
		}
		return nil, shared.ApiError("failed to list all anime movies", err)
	}

	var PBAnimeMovies []*ampb.AnimeMovieResponse
	for _, a := range DBAnimeMovies {
		PBAnimeMovies = append(PBAnimeMovies, shared.ConvertAnimeMovie(a))
	}

	res := &ampb.GetAllAnimeMoviesResponse{
		AnimeMovies: PBAnimeMovies,
	}
	return res, nil
}

func validateGetAllAnimeMoviesRequest(req *ampb.GetAllAnimeMoviesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shared.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shared.FieldViolation("pageSize", err))
	}

	if req.GetYear() != 0 {
		if err := utils.ValidateYear(req.GetYear()); err != nil {
			violations = append(violations, shared.FieldViolation("year", err))
		}
	}

	return violations
}
