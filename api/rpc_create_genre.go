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

func (server *Server) CreateGenre(ctx context.Context, req *pb.CreateGenreRequest) (*pb.CreateGenreResponse, error) {
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

	var DBgenres []db.Genre
	for _, name := range req.GetNames() {
		genre, err := server.gojo.CreateGenre(ctx, name)
		if err != nil {
			if db.ErrorCode(err) == db.UniqueViolation {
				return nil, status.Errorf(codes.AlreadyExists, err.Error())
			}
			return nil, status.Errorf(codes.Internal, "failed to create genre : %s", err)
		}
		DBgenres = append(DBgenres, genre)
	}

	var PBgenres []*pb.Genre
	for _, g := range DBgenres {
		genre := ConvertGenre(g)
		PBgenres = append(PBgenres, genre)
	}

	res := &pb.CreateGenreResponse{
		Genres: PBgenres,
	}

	return res, nil
}

func validateCreateGenreRequest(req *pb.CreateGenreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateGenreAndStudio(req.GetNames()); err != nil {
		violations = append(violations, fieldViolation("names", err))
	}

	return violations
}
