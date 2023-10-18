package api

import (
	"context"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAllAnimeMovies(ctx context.Context, req *pb.GetAllAnimeMoviesRequest) (*pb.GetAllAnimeMoviesResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all anime movies")
	}

	violations := validateGetAllAnimeMoviesRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListAnimeMoviesParams{
		ReleaseYear: req.GetYear(),
		Limit:       req.GetPageSize(),
		Offset:      (req.GetPageNumber() - 1) * req.GetPageSize(),
	}
	DBAnimeMovies, err := server.gojo.ListAnimeMovies(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to list the anime movies : %s", err)
	}

	var PBAnimeMovies []*pb.AnimeMovieResponse
	for _, a := range DBAnimeMovies {
		PBAnimeMovies = append(PBAnimeMovies, ConvertAnimeMovie(a))
	}

	res := &pb.GetAllAnimeMoviesResponse{
		AnimeMovies: PBAnimeMovies,
	}
	return res, nil
}

func validateGetAllAnimeMoviesRequest(req *pb.GetAllAnimeMoviesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, fieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, fieldViolation("pageSize", err))
	}

	if req.GetYear() != 0 {
		if err := utils.ValidateYear(req.GetYear()); err != nil {
			violations = append(violations, fieldViolation("year", err))
		}
	}

	return violations
}
