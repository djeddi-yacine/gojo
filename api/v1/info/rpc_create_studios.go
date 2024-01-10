package nfapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) CreateStudios(ctx context.Context, req *nfpbv1.CreateStudiosRequest) (*nfpbv1.CreateStudiosResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new studio")
	}

	if violations := validateCreateStudioRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	result, err := server.gojo.CreateStudiosTx(ctx, req.GetNames())
	if err != nil {
		return nil, shv1.ApiError("failed to create studio", err)
	}

	res := &nfpbv1.CreateStudiosResponse{
		Studios: shv1.ConvertStudios(result),
	}

	return res, nil
}

func validateCreateStudioRequest(req *nfpbv1.CreateStudiosRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateGenreAndStudio(req.GetNames()); err != nil {
		violations = append(violations, shv1.FieldViolation("names", err))
	}

	return violations
}
