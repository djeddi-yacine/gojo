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

func (server *InfoServer) GetAllStudios(ctx context.Context, req *nfpbv1.GetAllStudiosRequest) (*nfpbv1.GetAllStudiosResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAllStudiosRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	res := &nfpbv1.GetAllStudiosResponse{}
	studios, err := server.gojo.GetAllStudiosTx(ctx, db.ListStudiosParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	})
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code == pgerrcode.CaseNotFound {
				return res, nil
			}
		}
		return nil, shv1.ApiError("failed to list all studios", err)
	}

	res.Studios = shv1.ConvertStudios(studios)

	return res, nil
}

func validateGetAllStudiosRequest(req *nfpbv1.GetAllStudiosRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
