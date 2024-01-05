package asapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeSerieServer) QueryAnimeSeason(ctx context.Context, req *aspbv1.QueryAnimeSeasonRequest) (*aspbv1.QueryAnimeSeasonResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateQueryAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.QueryAnimeSeasonTxParams{
		Query:  req.GetQuery(),
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageNumber() - 1) * req.GetPageSize(),
	}

	data, err := server.gojo.QueryAnimeSeasonTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to query anime seasons", err)
	}

	var animeSeasons []*aspbv1.AnimeSeasonResponse
	for _, a := range data.AnimeSeasons {
		animeSeasons = append(animeSeasons, convertAnimeSeason(a))
	}

	res := &aspbv1.QueryAnimeSeasonResponse{
		AnimeSeasons: animeSeasons,
	}
	return res, nil
}

func validateQueryAnimeSeasonRequest(req *aspbv1.QueryAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetQuery(), 1, 150); err != nil {
		violations = append(violations, shv1.FieldViolation("query", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
