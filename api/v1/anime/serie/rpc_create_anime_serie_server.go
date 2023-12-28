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

func (server *AnimeSerieServer) CreateAnimeSerieServer(ctx context.Context, req *aspbv1.CreateAnimeSerieServerRequest) (*aspbv1.CreateAnimeSerieServerResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create episode server")
	}

	if violations := validateCreateAnimeSerieServerRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError([]*errdetails.BadRequest_FieldViolation{violations})
	}

	srv, err := server.gojo.CreateAnimeEpisodeServer(ctx, req.EpisodeID)
	if err != nil {
		return nil, shv1.ApiError("failed to create episode server", err)
	}

	res := &aspbv1.CreateAnimeSerieServerResponse{
		EpisodeID: srv.EpisodeID,
		ServerID:  srv.ID,
		CreatedAt: timestamppb.New(srv.CreatedAt),
	}
	return res, nil
}

func validateCreateAnimeSerieServerRequest(req *aspbv1.CreateAnimeSerieServerRequest) *errdetails.BadRequest_FieldViolation {
	if err := utils.ValidateInt(req.GetEpisodeID()); err != nil {
		return shv1.FieldViolation("episodeID", err)
	}
	return nil
}
