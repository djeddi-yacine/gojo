package animeSerie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) GetAllAnimeSeries(ctx context.Context, req *aspb.GetAllAnimeSeriesRequest) (*aspb.GetAllAnimeSeriesResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all anime Series")
	}

	violations := validateGetAllAnimeSeriesRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
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
		return nil, shared.DatabaseError("Failed to list all anime series", err)
	}

	var PBAnimeSeries []*aspb.AnimeSerieResponse
	for _, a := range DBAnimeSeries {
		PBAnimeSeries = append(PBAnimeSeries, shared.ConvertAnimeSerie(a))
	}

	res := &aspb.GetAllAnimeSeriesResponse{
		AnimeSeries: PBAnimeSeries,
	}
	return res, nil
}

func validateGetAllAnimeSeriesRequest(req *aspb.GetAllAnimeSeriesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shared.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shared.FieldViolation("pageSize", err))
	}

	if req.GetYear() != 0 {
		if err := utils.ValidateYear(req.GetYear()); err != nil {
			violations = append(violations, shared.FieldViolation("year", err))
		}
	}

	return violations
}
