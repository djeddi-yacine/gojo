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

func (server *AnimeSerieServer) QueryAnimeSeason(ctx context.Context, req *aspb.QueryAnimeSeasonRequest) (*aspb.QueryAnimeSeasonResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all anime seasons")
	}

	violations := validateQueryAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.QueryAnimeSeasonTxParams{
		Query: req.GetQuery(),
	}

	data, err := server.gojo.QueryAnimeSeasonTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to query anime seasons", err)
	}

	var animeSeasons []*aspb.AnimeSeasonResponse
	for _, a := range data.AnimeSeasons {
		animeSeasons = append(animeSeasons, shared.ConvertAnimeSeason(a))
	}

	res := &aspb.QueryAnimeSeasonResponse{
		AnimeSeasons: animeSeasons,
	}
	return res, nil
}

func validateQueryAnimeSeasonRequest(req *aspb.QueryAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetQuery(), 1, 150); err != nil {
		violations = append(violations, shared.FieldViolation("query", err))
	}

	return violations
}
