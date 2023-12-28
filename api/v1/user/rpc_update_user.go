package usapiv1

import (
	"context"
	"time"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) UpdateUser(ctx context.Context, req *uspbv1.UpdateUserRequest) (*uspbv1.UpdateUserResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if violations := validateUpdateUserRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	if authPayload.Username != req.GetUsername() && authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update other users info")
	}

	arg := db.UpdateUserParams{
		Username: req.GetUsername(),
		FullName: pgtype.Text{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}

	if req.Password != nil {
		hashedPassword, err := utils.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password : %s", err)
		}
		arg.HashedPassword = pgtype.Text{
			String: hashedPassword,
			Valid:  true,
		}
		arg.PasswordChangedAt = pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		}
	}

	user, err := server.gojo.UpdateUser(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to update user", err)
	}

	res := &uspbv1.UpdateUserResponse{
		User: &uspbv1.User{
			Username:          user.Username,
			FullName:          user.FullName,
			Email:             user.Email,
			PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
			CreatedAt:         timestamppb.New(user.CreatedAt),
		},
	}
	return res, nil
}

func validateUpdateUserRequest(req *uspbv1.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, shv1.FieldViolation("username", err))
	}

	if req.Password != nil {
		if err := utils.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, shv1.FieldViolation("password", err))
		}
	}

	if req.FullName != nil {
		if err := utils.ValidateFullName(req.GetFullName()); err != nil {
			violations = append(violations, shv1.FieldViolation("fullName", err))
		}
	}

	if req.Email != nil {
		if err := utils.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, shv1.FieldViolation("email", err))
		}
	}
	return violations
}
