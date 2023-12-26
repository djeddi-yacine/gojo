package info

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) CreateStudios(ctx context.Context, req *nfpb.CreateStudiosRequest) (*nfpb.CreateStudiosResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new studio")
	}

	if violations := validateCreateStudioRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	result, err := server.gojo.CreateStudiosTx(ctx, db.CreateStudiosTxParams{
		Names: req.GetNames(),
	})
	if err != nil {
		return nil, shared.ApiError("failed to create studio", err)
	}

	var PBStudios []*nfpb.Studio
	for _, s := range result.Studios {
		studio := shared.ConvertStudio(s)
		PBStudios = append(PBStudios, studio)
	}

	res := &nfpb.CreateStudiosResponse{
		Studios: PBStudios,
	}

	return res, nil
}

func validateCreateStudioRequest(req *nfpb.CreateStudiosRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateGenreAndStudio(req.GetNames()); err != nil {
		violations = append(violations, shared.FieldViolation("names", err))
	}

	return violations
}
