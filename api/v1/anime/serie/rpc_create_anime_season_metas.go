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

func (server *AnimeSerieServer) CreateAnimeSeasonMetas(ctx context.Context, req *aspbv1.CreateAnimeSeasonMetasRequest) (*aspbv1.CreateAnimeSeasonMetasResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime serie season metadata")
	}

	if violations := validateCreateAnimeSeasonMetasRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSeasonMetasTxParams{
		SeasonID: req.GetSeasonID(),
	}

	arg.SeasonMetas = make([]db.AnimeMetaTxParam, len(req.SeasonMetas))
	for i, v := range req.SeasonMetas {
		arg.SeasonMetas[i] = db.AnimeMetaTxParam{
			LanguageID: v.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    v.GetMeta().GetTitle(),
				Overview: v.GetMeta().GetOverview(),
			},
		}
	}

	data, err := server.gojo.CreateAnimeSeasonMetasTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to add anime serie season metadata", err)
	}

	res := &aspbv1.CreateAnimeSeasonMetasResponse{
		SeasonID: req.GetSeasonID(),
	}

	titles := make([]string, len(data.AnimeSeasonMetas))
	res.SeasonMetas = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeSeasonMetas))
	for i, v := range data.AnimeSeasonMetas {
		res.SeasonMetas[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(v.Meta),
			LanguageID: v.LanguageID,
			CreatedAt:  timestamppb.New(v.Meta.CreatedAt),
		}

		titles[i] = v.Meta.Title
	}

	server.meilisearch.AddDocuments(&utils.Document{
		ID:     req.GetSeasonID(),
		Titles: utils.RemoveDuplicatesTitles(titles),
	})

	return res, nil
}

func validateCreateAnimeSeasonMetasRequest(req *aspbv1.CreateAnimeSeasonMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := utils.ValidateInt(int64(req.GetSeasonID())); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.SeasonMetas != nil {
		for _, v := range req.SeasonMetas {
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
		violations = append(violations, shv1.FieldViolation("seasonMetas", errors.New("seasonMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
