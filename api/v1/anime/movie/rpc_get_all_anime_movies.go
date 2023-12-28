package amapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) GetAllAnimeMovies(ctx context.Context, req *ampbv1.GetAllAnimeMoviesRequest) (*ampbv1.GetAllAnimeMoviesResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all anime movies")
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
	DBAnimeMovies, err := server.gojo.ListAnimeMovies(ctx, arg)
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, nil
		}
		return nil, shv1.ApiError("failed to list all anime movies", err)
	}

	var PBAnimeMovies []*ampbv1.AnimeMovieResponse
	for _, a := range DBAnimeMovies {
		PBAnimeMovies = append(PBAnimeMovies, convertAnimeMovie(a))
	}

	res := &ampbv1.GetAllAnimeMoviesResponse{
		AnimeMovies: PBAnimeMovies,
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
