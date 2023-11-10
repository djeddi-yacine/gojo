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

func (server *AnimeSerieServer) AddAnimeSerieSeasonMetas(ctx context.Context, req *aspb.AddAnimeSerieSeasonMetasRequest) (*aspb.AddAnimeSerieSeasonMetasResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime serie season metas")
	}

	if violations := validateAddAnimeSerieSeasonMetasRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBSM = make([]db.AnimeMetaTxParam, len(req.SeasonMetas))
	for i, am := range req.SeasonMetas {
		DBSM[i] = db.AnimeMetaTxParam{
			LanguageID: am.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    am.GetMeta().GetTitle(),
				Overview: am.GetMeta().GetOverview(),
			},
		}
	}

	arg := db.AddAnimeSerieSeasonMetasTxParams{
		SeasonID:    int64(req.GetSeasonID()),
		SeasonMetas: DBSM,
	}

	data, err := server.gojo.AddAnimeSerieSeasonMetasTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to add anime serie season metas : %s", err)
	}

	var PBSM = make([]*nfpb.AnimeMetaResponse, len(data.AnimeSerieSeasonMetas))

	for i, am := range data.AnimeSerieSeasonMetas {
		PBSM[i] = &nfpb.AnimeMetaResponse{
			Meta:       shared.ConvertMeta(am.Meta),
			LanguageID: am.LanguageID,
			CreatedAt:  timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspb.AddAnimeSerieSeasonMetasResponse{
		Season:      shared.ConvertAnimeSerieSeason(data.AnimeSerieSeason),
		SeasonMetas: PBSM,
	}
	return res, nil
}

func validateAddAnimeSerieSeasonMetasRequest(req *aspb.AddAnimeSerieSeasonMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := utils.ValidateInt(int64(req.GetSeasonID())); err != nil {
		violations = append(violations, shared.FieldViolation("seasonID", err))
	}

	if req.SeasonMetas != nil {
		for _, am := range req.SeasonMetas {
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
		violations = append(violations, shared.FieldViolation("seasonMetas", errors.New("seasonMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
