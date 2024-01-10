package nfapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *InfoServer) GetAllGenres(ctx context.Context, req *nfpbv1.GetAllGenresRequest) (*nfpbv1.GetAllGenresResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAllGenresRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.ListGenresParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}
	dbGenres, err := server.gojo.ListGenres(ctx, arg)
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, nil
		}
		return nil, shv1.ApiError("failed to list all genres", err)
	}

	res := &nfpbv1.GetAllGenresResponse{
		Genres: shv1.ConvertGenres(dbGenres),
	}
	return res, nil
}

func validateGetAllGenresRequest(req *nfpbv1.GetAllGenresRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
