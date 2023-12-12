package info

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) GetAllLanguages(ctx context.Context, req *nfpb.GetAllLanguagesRequest) (*nfpb.GetAllLanguagesResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all Languages")
	}

	violations := validateGetAllLanguagesRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.ListLanguagesParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}
	DBLanguages, err := server.gojo.ListLanguages(ctx, arg)
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, nil
		}
		return nil, shared.DatabaseError("failed to list all languages", err)
	}

	var PBLanguages []*nfpb.LanguageResponse
	for _, g := range DBLanguages {
		PBLanguages = append(PBLanguages, shared.ConvertLanguage(g))
	}

	res := &nfpb.GetAllLanguagesResponse{
		Languages: PBLanguages,
	}
	return res, nil
}

func validateGetAllLanguagesRequest(req *nfpb.GetAllLanguagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shared.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shared.FieldViolation("pageSize", err))
	}

	return violations
}
