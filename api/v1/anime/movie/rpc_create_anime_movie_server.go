package amapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovieServer(ctx context.Context, req *ampbv1.CreateAnimeMovieServerRequest) (*ampbv1.CreateAnimeMovieServerResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie server")
	}

	if violations := validateCreateAnimeMovieServerRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError([]*errdetails.BadRequest_FieldViolation{violations})
	}

	srv, err := server.gojo.CreateAnimeMovieServer(ctx, req.AnimeID)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie server", err)
	}

	res := &ampbv1.CreateAnimeMovieServerResponse{
		AnimeID:   srv.AnimeID,
		ServerID:  srv.ID,
		CreatedAt: timestamppb.New(srv.CreatedAt),
	}
	return res, nil
}

func validateCreateAnimeMovieServerRequest(req *ampbv1.CreateAnimeMovieServerRequest) *errdetails.BadRequest_FieldViolation {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		return shv1.FieldViolation("animeID", err)
	}
	return nil
}
