package amapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) QueryAnimeMovie(ctx context.Context, req *ampbv1.QueryAnimeMovieRequest) (*ampbv1.QueryAnimeMovieResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all anime movies")
	}

	violations := validateQueryAnimeMovieRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.QueryAnimeMovieTxParams{
		Query:  req.GetQuery(),
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageNumber() - 1) * req.GetPageSize(),
	}

	data, err := server.gojo.QueryAnimeMovieTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to query anime movies", err)
	}

	var animeMovies []*ampbv1.AnimeMovieResponse
	for _, a := range data.AnimeMovies {
		animeMovies = append(animeMovies, convertAnimeMovie(a))
	}

	res := &ampbv1.QueryAnimeMovieResponse{
		AnimeMovies: animeMovies,
	}
	return res, nil
}

func validateQueryAnimeMovieRequest(req *ampbv1.QueryAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetQuery(), 1, 150); err != nil {
		violations = append(violations, shv1.FieldViolation("query", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
