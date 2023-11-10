package animeMovie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovieMetas(ctx context.Context, req *ampb.CreateAnimeMovieMetasRequest) (*ampb.CreateAnimeMovieMetasResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie metadata")
	}

	if violations := validateCreateAnimeMovieMetasRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
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

	arg := db.CreateAnimeMovieMetasTxParams{
		AnimeID:                       req.GetAnimeID(),
		CreateAnimeMovieMetasTxParams: DBAM,
	}

	metas, err := server.gojo.CreateAnimeMovieMetasTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime movie metadata : %s", err)
	}

	var anime_metas = make([]*nfpb.AnimeMetaResponse, len(metas.AnimeMovieMetasTxResults))

	for i, am := range metas.AnimeMovieMetasTxResults {
		anime_metas[i] = &nfpb.AnimeMetaResponse{
			Meta:       shared.ConvertMeta(am.Meta),
			LanguageID: am.LanguageID,
			CreatedAt:  timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &ampb.CreateAnimeMovieMetasResponse{
		AnimeID:    req.GetAnimeID(),
		AnimeMetas: anime_metas,
	}
	return res, nil
}

func validateCreateAnimeMovieMetasRequest(req *ampb.CreateAnimeMovieMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
