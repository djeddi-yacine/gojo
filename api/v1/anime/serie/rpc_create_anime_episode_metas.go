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

func (server *AnimeSerieServer) CreateAnimeEpisodeMetas(ctx context.Context, req *aspbv1.CreateAnimeEpisodeMetasRequest) (*aspbv1.CreateAnimeEpisodeMetasResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime serie episode metadata")
	}

	if violations := validateCreateAnimeEpisodeMetasRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	var DBEM = make([]db.AnimeMetaTxParam, len(req.EpisodeMetas))
	for i, am := range req.EpisodeMetas {
		DBEM[i] = db.AnimeMetaTxParam{
			LanguageID: am.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    am.GetMeta().GetTitle(),
				Overview: am.GetMeta().GetOverview(),
			},
		}
	}

	arg := db.CreateAnimeEpisodeMetasTxParams{
		EpisodeID:    req.GetEpisodeID(),
		EpisodeMetas: DBEM,
	}

	data, err := server.gojo.CreateAnimeEpisodeMetasTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to add anime serie episode metadata", err)
	}

	var PBSM = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeEpisodeMetas))

	for i, am := range data.AnimeEpisodeMetas {
		PBSM[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(am.Meta),
			LanguageID: am.LanguageID,
			CreatedAt:  timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspbv1.CreateAnimeEpisodeMetasResponse{
		EpisodeID:    req.GetEpisodeID(),
		EpisodeMetas: PBSM,
	}

	return res, nil
}

func validateCreateAnimeEpisodeMetasRequest(req *aspbv1.CreateAnimeEpisodeMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := utils.ValidateInt(int64(req.GetEpisodeID())); err != nil {
		violations = append(violations, shv1.FieldViolation("episodeID", err))
	}

	if req.EpisodeMetas != nil {
		for _, am := range req.EpisodeMetas {
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
		violations = append(violations, shv1.FieldViolation("episodeMetas", errors.New("episodeMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
