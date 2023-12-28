package usapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *UserServer) VerifyEmail(ctx context.Context, req *uspbv1.VerifyEmailRequest) (*uspbv1.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	txResult, err := server.gojo.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailID:    req.GetEmailID(),
		SecretCode: req.GetSecretCode(),
	})
	if err != nil {
		return nil, shv1.ApiError("failed to verify email", err)
	}

	rsp := &uspbv1.VerifyEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
	}
	return rsp, nil
}

func validateVerifyEmailRequest(req *uspbv1.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetEmailID()); err != nil {
		violations = append(violations, shv1.FieldViolation("emailID", err))
	}

	if err := utils.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, shv1.FieldViolation("secretCode", err))
	}

	return violations
}
