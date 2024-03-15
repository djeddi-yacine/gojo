package nfapiv1

import (
	"context"
	"errors"
	"fmt"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) CreateLanguages(ctx context.Context, req *nfpbv1.CreateLanguagesRequest) (*nfpbv1.CreateLanguagesResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new language")
	}

	if violations := validateCreateLanguagesRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	languages := make([]db.CreateLanguageParams, len(req.GetLanguages()))
	for i, x := range req.GetLanguages() {
		languages[i] = db.CreateLanguageParams{
			LanguageName: x.LanguageName,
			LanguageCode: x.LanguageCode,
		}
	}

	result, err := server.gojo.CreateLanguagesTx(ctx, languages)
	if err != nil {
		return nil, shv1.ApiError("failed to create new language", err)
	}

	res := &nfpbv1.CreateLanguagesResponse{}

	res.Languages = make([]*nfpbv1.LanguageResponse, len(result))
	for i, x := range result {
		res.Languages[i] = shv1.ConvertLanguage(x)
	}

	return res, nil
}

func validateCreateLanguagesRequest(req *nfpbv1.CreateLanguagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetLanguages() != nil {
		for i, v := range req.GetLanguages() {
			if err := utils.ValidateString(v.LanguageCode, 2, 3); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("languages > code at [%d]", i), err))
			}
			if err := utils.ValidateString(v.LanguageName, 2, 15); err != nil {
				violations = append(violations, shv1.FieldViolation(fmt.Sprintf("languages > name at [%d]", i), err))
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("languages", errors.New("you need to send the Language model")))
	}

	return violations
}
