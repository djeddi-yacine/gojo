package api

import (
	"context"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateLanguage(ctx context.Context, req *pb.CreateLanguageRequest) (*pb.CreateLanguageResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new language")
	}

	if violations := validateCreateLanguageRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	var DBlanguages []db.Language
	for _, l := range req.GetLanguages() {
		arg := db.CreateLanguageParams{
			LanguageName: l.LanguageName,
			LanguageCode: l.LanguageCode,
		}
		language, err := server.gojo.CreateLanguage(ctx, arg)
		if err != nil {
			if db.ErrorCode(err) == db.UniqueViolation {
				return nil, status.Errorf(codes.AlreadyExists, err.Error())
			}
			return nil, status.Errorf(codes.Internal, "failed to create language : %s", err)
		}
		DBlanguages = append(DBlanguages, language)
	}

	var PBlanguages []*pb.LanguageResponse
	for _, l := range DBlanguages {
		language := ConvertLanguage(l)
		PBlanguages = append(PBlanguages, language)
	}

	res := &pb.CreateLanguageResponse{
		Languages: PBlanguages,
	}

	return res, nil
}

func validateCreateLanguageRequest(req *pb.CreateLanguageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateLanguage(req.GetLanguages()); err != nil {
		violations = append(violations, fieldViolation("languages", err))
	}

	return violations
}
