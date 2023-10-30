package animeSerie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeSerieMetas(ctx context.Context, req *aspb.CreateAnimeSerieMetasRequest) (*aspb.CreateAnimeSerieMetasResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime Serie metadata")
	}

	if violations := validateCreateAnimeSerieMetasRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBAM = make([]db.CreateAnimeSerieMetaTxParam, len(req.AnimeMetas))
	for i, am := range req.AnimeMetas {
		DBAM[i] = db.CreateAnimeSerieMetaTxParam{
			LanguageID: am.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    am.GetMeta().GetTitle(),
				Overview: am.GetMeta().GetOverview(),
			},
		}
	}

	arg := db.CreateAnimeSerieMetasTxParams{
		AnimeID:                       req.GetAnimeID(),
		CreateAnimeSerieMetasTxParams: DBAM,
	}

	metas, err := server.gojo.CreateAnimeSerieMetasTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie metadata : %s", err)
	}

	var PBAM = make([]*nfpb.AnimeMetaResponse, len(metas.CreateAnimeSerieMetasTxResults))

	for i, am := range metas.CreateAnimeSerieMetasTxResults {
		PBAM[i] = &nfpb.AnimeMetaResponse{
			Meta:      shared.ConvertMeta(am.Meta),
			Language:  shared.ConvertLanguage(am.Language),
			CreatedAt: timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspb.CreateAnimeSerieMetasResponse{
		AnimeID:    req.GetAnimeID(),
		AnimeMetas: PBAM,
	}
	return res, nil
}

func validateCreateAnimeSerieMetasRequest(req *aspb.CreateAnimeSerieMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	if req.AnimeMetas != nil {
		for _, am := range req.AnimeMetas {
			if err := utils.ValidateInt(int64(am.GetLanguageID())); err != nil {
				violations = append(violations, shared.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shared.FieldViolation("title", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shared.FieldViolation("overview", err))
			}
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeMetas", errors.New("give at least one metadata")))
	}

	return violations
}
