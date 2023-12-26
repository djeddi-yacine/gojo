package user

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *UserServer) VerifyEmail(ctx context.Context, req *uspb.VerifyEmailRequest) (*uspb.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	txResult, err := server.gojo.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailID:    req.GetEmailID(),
		SecretCode: req.GetSecretCode(),
	})
	if err != nil {
		return nil, shared.ApiError("failed to verify email", err)
	}

	rsp := &uspb.VerifyEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
	}
	return rsp, nil
}

func validateVerifyEmailRequest(req *uspb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetEmailID()); err != nil {
		violations = append(violations, shared.FieldViolation("emailID", err))
	}

	if err := utils.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, shared.FieldViolation("secretCode", err))
	}

	return violations
}
