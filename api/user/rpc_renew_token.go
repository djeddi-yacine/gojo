package user

import (
	"context"
	"time"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) RenewTokens(ctx context.Context, req *uspb.RenewTokensRequest) (*uspb.RenewTokensResponse, error) {
	if violations := validateRenewTokensRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot use this refresh tokens : %s", err)
	}

	session, err := server.gojo.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		return nil, shared.DatabaseError("failed to get session", err)
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

	arg := db.RenewSessionTxParams{
		CreateSessionParams: db.CreateSessionParams{
			ID:           refreshPayload.ID,
			Username:     refreshPayload.Username,
			RefreshToken: refreshToken,
			UserAgent:    session.UserAgent,
			ClientIp:     session.ClientIp,
			IsBlocked:    false,
			ExpiresAt:    refreshPayload.ExpiredAt,
		},
		AfterRenew: func(username string) error {
			taskPayload := &worker.PayloadDeleteSession{
				Username: username,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}

			return server.taskDistributor.DistributeTaskDeleteSession(ctx, taskPayload, opts...)
		},
	}

	s, err := server.gojo.RenewSessionTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to renew session", err)
	}

	res := &uspb.RenewTokensResponse{
		SessionID:             s.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return res, nil
}

func validateRenewTokensRequest(req *uspb.RenewTokensRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateToken(req.GetRefreshToken()); err != nil {
		violations = append(violations, shared.FieldViolation("refreshToken", err))
	}
	return violations
}
