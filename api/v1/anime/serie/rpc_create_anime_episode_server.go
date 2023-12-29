package asapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeEpisodeServer(ctx context.Context, req *aspbv1.CreateAnimeEpisodeServerRequest) (*aspbv1.CreateAnimeEpisodeServerResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create episode server")
	}

	if violations := validateCreateAnimeEpisodeServerRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError([]*errdetails.BadRequest_FieldViolation{violations})
	}

	srv, err := server.gojo.CreateAnimeEpisodeServerTx(ctx, req.EpisodeID)
	if err != nil {
		return nil, shv1.ApiError("failed to create episode server", err)
	}

	res := &aspbv1.CreateAnimeEpisodeServerResponse{
		EpisodeID: srv.EpisodeID,
		ServerID:  srv.ID,
		CreatedAt: timestamppb.New(srv.CreatedAt),
	}
	return res, nil
}

func validateCreateAnimeEpisodeServerRequest(req *aspbv1.CreateAnimeEpisodeServerRequest) *errdetails.BadRequest_FieldViolation {
	if err := utils.ValidateInt(req.GetEpisodeID()); err != nil {
		return shv1.FieldViolation("episodeID", err)
	}
	return nil
}
