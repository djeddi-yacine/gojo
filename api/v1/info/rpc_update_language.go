package nfapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) UpdateLanguage(ctx context.Context, req *nfpbv1.UpdateLanguageRequest) (*nfpbv1.UpdateLanguageResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new language")
	}

	if violations := validateUpdateLanguageRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	data, err := server.gojo.UpdateLanguage(ctx, db.UpdateLanguageParams{
		ID: req.GetLanguageID(),
		LanguageCode: pgtype.Text{
			String: req.GetLanguageCode(),
			Valid:  req.LanguageCode != nil,
		},
		LanguageName: pgtype.Text{
			String: req.GetLanguageName(),
			Valid:  req.LanguageName != nil,
		},
	})
	if err != nil {
		return nil, shv1.ApiError("failed to update language", err)
	}

	res := &nfpbv1.UpdateLanguageResponse{
		Languages: shv1.ConvertLanguage(data),
	}

	return res, nil
}

func validateUpdateLanguageRequest(req *nfpbv1.UpdateLanguageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	if req.LanguageCode != nil {
		if err := utils.ValidateString(req.GetLanguageCode(), 2, 3); err != nil {
			violations = append(violations, shv1.FieldViolation("languageCode", err))
		}
	}

	if req.LanguageName != nil {
		if err := utils.ValidateString(req.GetLanguageName(), 2, 15); err != nil {
			violations = append(violations, shv1.FieldViolation("languageName", err))
		}
	}

	return violations
}
