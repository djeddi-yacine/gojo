package amapiv1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovieMetas(ctx context.Context, req *ampbv1.CreateAnimeMovieMetasRequest) (*ampbv1.CreateAnimeMovieMetasResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie metadata")
	}

	if violations := validateCreateAnimeMovieMetasRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieMetasTxParams{
		AnimeID: req.GetAnimeID(),
	}

	arg.AnimeMovieMetas = make([]db.AnimeMetaTxParam, len(req.AnimeMetas))
	for i, v := range req.AnimeMetas {
		arg.AnimeMovieMetas[i] = db.AnimeMetaTxParam{
			LanguageID: v.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    v.GetMeta().GetTitle(),
				Overview: v.GetMeta().GetOverview(),
			},
		}
	}

	metas, err := server.gojo.CreateAnimeMovieMetasTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime movie metadata", err)
	}

	res := &ampbv1.CreateAnimeMovieMetasResponse{
		AnimeID: req.GetAnimeID(),
	}

	res.AnimeMetas = make([]*nfpbv1.AnimeMetaResponse, len(metas.AnimeMovieMetas))
	for i, v := range metas.AnimeMovieMetas {
		res.AnimeMetas[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(v.Meta),
			LanguageID: v.LanguageID,
			CreatedAt:  timestamppb.New(v.Meta.CreatedAt),
		}
	}

	return res, nil
}

func validateCreateAnimeMovieMetasRequest(req *ampbv1.CreateAnimeMovieMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if req.AnimeMetas != nil {
		for _, v := range req.AnimeMetas {
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
		violations = append(violations, shv1.FieldViolation("animeMetas", errors.New("give at least one metadata")))
	}

	return violations
}
