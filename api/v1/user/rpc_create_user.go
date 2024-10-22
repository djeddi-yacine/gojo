package usapiv1

import (
	"context"
	"time"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) CreateUser(ctx context.Context, req *uspbv1.CreateUserRequest) (*uspbv1.CreateUserResponse, error) {
	if violations := validateCreateUserRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password : %s", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.GetUsername(),
			Email:          req.GetEmail(),
			HashedPassword: hashedPassword,
			FullName:       req.GetFullName(),
		},
		AfterCreate: func(user db.User) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				Username: user.Username,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue("critical"),
			}

			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	txResult, err := server.gojo.CreateUserTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create the user", err)
	}

	res := &uspbv1.CreateUserResponse{
		User: &uspbv1.User{
			ID:                txResult.User.ID,
			Username:          txResult.User.Username,
			FullName:          txResult.User.FullName,
			Email:             txResult.User.Email,
			PasswordChangedAt: timestamppb.New(txResult.User.PasswordChangedAt),
			CreatedAt:         timestamppb.New(txResult.User.CreatedAt),
		},
	}
	return res, nil
}

func validateCreateUserRequest(req *uspbv1.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, shv1.FieldViolation("username", err))
	}

	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, shv1.FieldViolation("password", err))
	}

	if err := utils.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, shv1.FieldViolation("fullName", err))
	}

	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, shv1.FieldViolation("email", err))
	}

	return violations
}
