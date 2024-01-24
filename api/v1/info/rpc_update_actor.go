package nfapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) UpdateActor(ctx context.Context, req *nfpbv1.UpdateActorRequest) (*nfpbv1.UpdateActorResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update actor")
	}

	if violations := validateUpdateActorRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	data, err := server.gojo.UpdateActor(ctx, db.UpdateActorParams{
		ID: req.GetActorID(),
		FullName: pgtype.Text{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Biography: pgtype.Text{
			String: req.GetBiography(),
			Valid:  req.Biography != nil,
		},
		Gender: pgtype.Text{
			String: req.GetGender(),
			Valid:  req.Gender != nil,
		},
		ImageUrl: pgtype.Text{
			String: req.GetImage(),
			Valid:  req.Image != nil,
		},
		ImageBlurHash: pgtype.Text{
			String: req.GetImageBlurHash(),
			Valid:  req.ImageBlurHash != nil,
		},
	})
	if err != nil {
		return nil, shv1.ApiError("failed to update actor", err)
	}

	res := &nfpbv1.UpdateActorResponse{
		Actor: shv1.ConvertActor(data),
	}

	return res, nil
}

func validateUpdateActorRequest(req *nfpbv1.UpdateActorRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetActorID())); err != nil {
		violations = append(violations, shv1.FieldViolation("actorID", err))
	}

	if req.FullName != nil {
		if err := utils.ValidateString(req.GetFullName(), 2, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("fullName", err))
		}
	}
	if req.Biography != nil {
		if err := utils.ValidateString(req.GetBiography(), 2, 5000); err != nil {
			violations = append(violations, shv1.FieldViolation("biography", err))
		}
	}

	if req.Gender != nil {
		if err := utils.ValidateString(req.GetGender(), 2, 20); err != nil {
			violations = append(violations, shv1.FieldViolation("gender", err))
		}
	}

	if req.Image != nil {
		if err := utils.ValidateURL(req.GetImage(), ""); err != nil {
			violations = append(violations, shv1.FieldViolation("image", err))
		}
	}

	if req.ImageBlurHash != nil {
		if err := utils.ValidateString(req.GetImageBlurHash(), 2, 32); err != nil {
			violations = append(violations, shv1.FieldViolation("imageBlurhash", err))
		}
	}

	return violations
}
