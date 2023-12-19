package animeMovie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) QueryAnimeMovie(ctx context.Context, req *ampb.QueryAnimeMovieRequest) (*ampb.QueryAnimeMovieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all anime movies")
	}

	violations := validateQueryAnimeMovieRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.QueryAnimeMovieTxParams{
		Query:  req.GetQuery(),
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageNumber() - 1) * req.GetPageSize(),
	}

	data, err := server.gojo.QueryAnimeMovieTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to query anime movies", err)
	}

	var animeMovies []*ampb.AnimeMovieResponse
	for _, a := range data.AnimeMovies {
		animeMovies = append(animeMovies, shared.ConvertAnimeMovie(a))
	}

	res := &ampb.QueryAnimeMovieResponse{
		AnimeMovies: animeMovies,
	}
	return res, nil
}

func validateQueryAnimeMovieRequest(req *ampb.QueryAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetQuery(), 1, 150); err != nil {
		violations = append(violations, shared.FieldViolation("query", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shared.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shared.FieldViolation("pageSize", err))
	}

	return violations
}
