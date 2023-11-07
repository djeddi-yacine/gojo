package animeMovie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovieServer(ctx context.Context, req *ampb.CreateAnimeMovieServerRequest) (*ampb.CreateAnimeMovieServerResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie server")
	}

	if violations := validateCreateAnimeMovieServerRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError([]*errdetails.BadRequest_FieldViolation{violations})
	}

	srv, err := server.gojo.CreateAnimeMovieServer(ctx, req.AnimeID)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime movie server : %s", err)
	}

	res := &ampb.CreateAnimeMovieServerResponse{
		AnimeID:   srv.AnimeID,
		ServerID:  srv.ID,
		CreatedAt: timestamppb.New(srv.CreatedAt),
	}
	return res, nil
}

func validateCreateAnimeMovieServerRequest(req *ampb.CreateAnimeMovieServerRequest) *errdetails.BadRequest_FieldViolation {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		return shared.FieldViolation("animeID", err)
	}
	return nil
}
