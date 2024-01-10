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

func (server *InfoServer) GetAllLanguages(ctx context.Context, req *nfpbv1.GetAllLanguagesRequest) (*nfpbv1.GetAllLanguagesResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAllLanguagesRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.ListLanguagesParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}
	dbLanguages, err := server.gojo.ListLanguages(ctx, arg)
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, nil
		}
		return nil, shv1.ApiError("failed to list all languages", err)
	}

	pbLanguages := make([]*nfpbv1.LanguageResponse, len(dbLanguages))
	for i, x := range dbLanguages {
		pbLanguages[i] = shv1.ConvertLanguage(x)
	}

	res := &nfpbv1.GetAllLanguagesResponse{
		Languages: pbLanguages,
	}
	return res, nil
}

func validateGetAllLanguagesRequest(req *nfpbv1.GetAllLanguagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
