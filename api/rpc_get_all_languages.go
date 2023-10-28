package api

import (
	"context"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAllLanguages(ctx context.Context, req *nfpb.GetAllLanguagesRequest) (*nfpb.GetAllLanguagesResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get all Languages")
	}

	violations := validateGetAllLanguagesRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListLanguagesParams{
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	}
	DBLanguages, err := server.gojo.ListLanguages(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to list the Languages : %s", err)
	}

	var PBLanguages []*nfpb.LanguageResponse
	for _, g := range DBLanguages {
		PBLanguages = append(PBLanguages, ConvertLanguage(g))
	}

	res := &nfpb.GetAllLanguagesResponse{
		Languages: PBLanguages,
	}
	return res, nil
}

func validateGetAllLanguagesRequest(req *nfpb.GetAllLanguagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, fieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, fieldViolation("pageSize", err))
	}

	return violations
}
