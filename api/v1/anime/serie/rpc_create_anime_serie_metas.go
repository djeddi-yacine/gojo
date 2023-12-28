package asapiv1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSerieMetas(ctx context.Context, req *aspbv1.CreateAnimeSerieMetasRequest) (*aspbv1.CreateAnimeSerieMetasResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime Serie metadata")
	}

	if violations := validateCreateAnimeSerieMetasRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	var DBAM = make([]db.AnimeMetaTxParam, len(req.AnimeMetas))
	for i, am := range req.AnimeMetas {
		DBAM[i] = db.AnimeMetaTxParam{
			LanguageID: am.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    am.GetMeta().GetTitle(),
				Overview: am.GetMeta().GetOverview(),
			},
		}
	}

	arg := db.CreateAnimeSerieMetasTxParams{
		AnimeID:         req.GetAnimeID(),
		AnimeSerieMetas: DBAM,
	}

	metas, err := server.gojo.CreateAnimeSerieMetasTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime serie metadata", err)
	}

	var PBAM = make([]*nfpbv1.AnimeMetaResponse, len(metas.AnimeSerieMetas))

	for i, am := range metas.AnimeSerieMetas {
		PBAM[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(am.Meta),
			LanguageID: am.LanguageID,
			CreatedAt:  timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspbv1.CreateAnimeSerieMetasResponse{
		AnimeID:    req.GetAnimeID(),
		AnimeMetas: PBAM,
	}

	return res, nil
}

func validateCreateAnimeSerieMetasRequest(req *aspbv1.CreateAnimeSerieMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.AnimeMetas != nil {
		for _, am := range req.AnimeMetas {
			if err := utils.ValidateInt(int64(am.GetLanguageID())); err != nil {
				violations = append(violations, shv1.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shv1.FieldViolation("title", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shv1.FieldViolation("overview", err))
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeMetas", errors.New("give at least one metadata")))
	}

	return violations
}
