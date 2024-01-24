package nfapiv1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) UpdateStudio(ctx context.Context, req *nfpbv1.UpdateStudioRequest) (*nfpbv1.UpdateStudioResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update studio")
	}

	if violations := validateUpdateStudioRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	data, err := server.gojo.UpdateStudio(ctx, db.UpdateStudioParams{
		ID: req.GetStudioID(),
		StudioName: pgtype.Text{
			String: req.GetStudioName(),
			Valid:  req.StudioName != nil,
		},
	})
	if err != nil {
		return nil, shv1.ApiError("failed to update studio", err)
	}

	res := &nfpbv1.UpdateStudioResponse{
		Studio: shv1.ConvertStudio(data),
	}

	return res, nil
}

func validateUpdateStudioRequest(req *nfpbv1.UpdateStudioRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetStudioID())); err != nil {
		violations = append(violations, shv1.FieldViolation("studioID", err))
	}

	if req.StudioName != nil {
		if err := utils.ValidateString(req.GetStudioName(), 2, 15); err != nil {
			violations = append(violations, shv1.FieldViolation("studioName", err))
		}
	} else {
		violations = append(violations, shv1.FieldViolation("studioName", errors.New("put the studio name")))
	}

	return violations
}
