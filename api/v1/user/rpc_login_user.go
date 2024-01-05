package usapiv1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) LoginUser(ctx context.Context, req *uspbv1.LoginUserRequest) (*uspbv1.LoginUserResponse, error) {
	if violations := validateLoginUserRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	md := shv1.ExtractMetadata(ctx)
	user, err := server.gojo.LoginUserTx(ctx, db.LoginUserTxParams{
		Username:        req.Username,
		Password:        req.Password,
		DeviceName:      req.DeviceName,
		DeviceHash:      req.DeviceHash,
		OperatingSystem: req.OperatingSystem,
		MacAddress:      req.MacAddress,
		UserAgent:       md.UserAgent,
		ClientIp:        md.ClientIP,
	})
	if err != nil {
		if errors.Is(err, db.ErrInccorectPassword) {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
		if errors.Is(err, db.ErrFailedPrecondition) {
			return nil, status.Errorf(codes.FailedPrecondition, err.Error())
		}
		return nil, shv1.ApiError("failed to login user", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		user.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token : %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		user.Role,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token : %s", err)
	}

	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := server.gojo.CreateSession(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to renew session", err)
	}

	res := &uspbv1.LoginUserResponse{
		User: &uspbv1.User{
			ID:                user.ID,
			Username:          user.Username,
			FullName:          user.FullName,
			Email:             user.Email,
			PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
			CreatedAt:         timestamppb.New(user.CreatedAt),
		},
		SessionID:             session.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return res, nil
}

func validateLoginUserRequest(req *uspbv1.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, shv1.FieldViolation("username", err))
	}

	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, shv1.FieldViolation("password", err))
	}

	if err := utils.ValidateString(req.GetDeviceName(), 2, 100); err != nil {
		violations = append(violations, shv1.FieldViolation("deviceName", err))
	}

	if err := utils.ValidateString(req.GetDeviceHash(), 20, 35); err != nil {
		violations = append(violations, shv1.FieldViolation("deviceHash", err))
	}

	if err := utils.ValidateString(req.GetOperatingSystem(), 3, 100); err != nil {
		violations = append(violations, shv1.FieldViolation("operatingSystem", err))
	}

	if err := utils.ValidateMAC(req.GetMacAddress()); err != nil {
		violations = append(violations, shv1.FieldViolation("macAddress", err))
	}

	return violations
}
