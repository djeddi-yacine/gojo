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

func (server *Server) CreateStudios(ctx context.Context, req *pb.CreateStudiosRequest) (*pb.CreateStudiosResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new studio")
	}

	if violations := validateCreateStudioRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	result, err := server.gojo.CreateStudiosTx(ctx, db.CreateStudiosTxParams{
		Names: req.GetNames(),
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create studio : %s", err)
	}

	var PBStudios []*pb.Studio
	for _, s := range result.Studios {
		studio := ConvertStudio(s)
		PBStudios = append(PBStudios, studio)
	}

	res := &pb.CreateStudiosResponse{
		Studios: PBStudios,
	}

	return res, nil
}

func validateCreateStudioRequest(req *pb.CreateStudiosRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateGenreAndStudio(req.GetNames()); err != nil {
		violations = append(violations, fieldViolation("names", err))
	}

	return violations
}
