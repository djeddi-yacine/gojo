package asapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeSerieServer) GetAllAnimeSeries(ctx context.Context, req *aspbv1.GetAllAnimeSeriesRequest) (*aspbv1.GetAllAnimeSeriesResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAllAnimeSeriesRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.ListAnimeSeriesParams{
		FirstYear: req.GetYear(),
		Limit:     req.GetPageSize(),
		Offset:    (req.GetPageNumber() - 1) * req.GetPageSize(),
	}
	DBAnimeSeries, err := server.gojo.ListAnimeSeries(ctx, arg)
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, nil
		}
		return nil, shv1.ApiError("failed to list anime serie seasons", err)
	}

	var PBAnimeSeries []*aspbv1.AnimeSerieResponse
	for _, a := range DBAnimeSeries {
		PBAnimeSeries = append(PBAnimeSeries, convertAnimeSerie(a))
	}

	res := &aspbv1.GetAllAnimeSeriesResponse{
		AnimeSeries: PBAnimeSeries,
	}

	return res, nil
}

func validateGetAllAnimeSeriesRequest(req *aspbv1.GetAllAnimeSeriesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	if req.GetYear() != 0 {
		if err := utils.ValidateYear(req.GetYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("year", err))
		}
	}

	return violations
}
