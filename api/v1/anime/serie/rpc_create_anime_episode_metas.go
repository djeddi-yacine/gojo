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

	arg := db.CreateAnimeEpisodeMetasTxParams{
		EpisodeID: req.GetEpisodeID(),
	}

	arg.EpisodeMetas = make([]db.AnimeMetaTxParam, len(req.EpisodeMetas))
	for i, v := range req.EpisodeMetas {
		arg.EpisodeMetas[i] = db.AnimeMetaTxParam{
			LanguageID: v.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    v.GetMeta().GetTitle(),
				Overview: v.GetMeta().GetOverview(),
			},
		}
	}

	data, err := server.gojo.CreateAnimeEpisodeMetasTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to add anime serie episode metadata", err)
	}

	res := &aspbv1.CreateAnimeEpisodeMetasResponse{
		EpisodeID: req.GetEpisodeID(),
	}

	res.EpisodeMetas = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeEpisodeMetas))
	for i, v := range data.AnimeEpisodeMetas {
		res.EpisodeMetas[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(v.Meta),
			LanguageID: v.LanguageID,
			CreatedAt:  timestamppb.New(v.Meta.CreatedAt),
		}
	}

	return res, nil
}

func validateCreateAnimeEpisodeMetasRequest(req *aspbv1.CreateAnimeEpisodeMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := utils.ValidateInt(int64(req.GetEpisodeID())); err != nil {
		violations = append(violations, shv1.FieldViolation("episodeID", err))
	}

	if req.EpisodeMetas != nil {
		for _, v := range req.EpisodeMetas {
			if err := utils.ValidateInt(int64(v.GetLanguageID())); err != nil {
				violations = append(violations, shv1.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(v.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shv1.FieldViolation("title", err))
			}

			if err := utils.ValidateString(v.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shv1.FieldViolation("overview", err))
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("episodeMetas > meta", errors.New("you need to send at least one of meta model")))
	}

	return violations
}
