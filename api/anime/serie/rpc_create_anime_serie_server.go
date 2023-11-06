package animeSerie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSerieServer(ctx context.Context, req *aspb.CreateAnimeSerieServerRequest) (*aspb.CreateAnimeSerieServerResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create episode server")
	}

	if violations := validateCreateAnimeSerieServerRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError([]*errdetails.BadRequest_FieldViolation{violations})
	}

	srv, err := server.gojo.CreateAnimeSerieServer(ctx, req.EpisodeID)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create episode server : %s", err)
	}

	res := &aspb.CreateAnimeSerieServerResponse{
		EpisodeID: srv.EpisodeID,
		ServerID:  srv.ID,
	}
	return res, nil
}

func validateCreateAnimeSerieServerRequest(req *aspb.CreateAnimeSerieServerRequest) *errdetails.BadRequest_FieldViolation {
	if err := utils.ValidateInt(req.GetEpisodeID()); err != nil {
		return shared.FieldViolation("episodeID", err)
	}
	return nil
}
