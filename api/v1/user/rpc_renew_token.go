package usapiv1

import (
	"context"
	"time"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) RenewTokens(ctx context.Context, req *uspbv1.RenewTokensRequest) (*uspbv1.RenewTokensResponse, error) {
	if violations := validateRenewTokensRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot use this refresh tokens : %s", err)
	}

	session, err := server.gojo.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		return nil, shv1.ApiError("failed to get session", err)
	}

	if session.IsBlocked {
		return nil, status.Errorf(codes.Unauthenticated, "blocked session")
	}

	if session.Username != refreshPayload.Username {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect session user")
	}

	if session.RefreshToken != req.RefreshToken {
		return nil, status.Errorf(codes.Unauthenticated, "mismatched session tokens")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "expired session")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		refreshPayload.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token : %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		refreshPayload.Role,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token : %s", err)
	}

	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     refreshPayload.Username,
		RefreshToken: refreshToken,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	s, err := server.gojo.CreateSession(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to renew session", err)
	}

	res := &uspbv1.RenewTokensResponse{
		SessionID:             s.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return res, nil
}

func validateRenewTokensRequest(req *uspbv1.RenewTokensRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateToken(req.GetRefreshToken()); err != nil {
		violations = append(violations, shv1.FieldViolation("refreshToken", err))
	}
	return violations
}
