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

	res := &aspbv1.GetAllAnimeSeriesResponse{}

	data, err := server.gojo.ListAnimeSeries(ctx, db.ListAnimeSeriesParams{
		FirstYear: req.GetYear(),
		Limit:     req.GetPageSize(),
		Offset:    (req.GetPageNumber() - 1) * req.GetPageSize(),
	})
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code == pgerrcode.CaseNotFound {
				return res, nil
			}
		}

		return nil, shv1.ApiError("failed to list anime serie seasons", err)
	}

	res.AnimeSeries = make([]*aspbv1.AnimeSerieResponse, len(data))
	for i, v := range data {
		res.AnimeSeries[i] = convertAnimeSerie(v)
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
