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

func (server *Server) CreateGenres(ctx context.Context, req *pb.CreateGenresRequest) (*pb.CreateGenresResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new genre")
	}

	if violations := validateCreateGenreRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	result, err := server.gojo.CreateGenresTx(ctx, db.CreateGenresTxParams{
		Names: req.GetNames(),
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create genre : %s", err)
	}

	var PBgenres []*pb.Genre
	for _, g := range result.Genres {
		genre := ConvertGenre(g)
		PBgenres = append(PBgenres, genre)
	}

	res := &pb.CreateGenresResponse{
		Genres: PBgenres,
	}

	return res, nil
}

func validateCreateGenreRequest(req *pb.CreateGenresRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateGenreAndStudio(req.GetNames()); err != nil {
		violations = append(violations, fieldViolation("names", err))
	}

	return violations
}
