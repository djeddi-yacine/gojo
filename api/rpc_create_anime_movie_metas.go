package api

import (
	"context"
	"errors"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAnimeMovieMetas(ctx context.Context, req *pb.CreateAnimeMovieMetasRequest) (*pb.CreateAnimeMovieMetasResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie metadata")
	}

	if violations := validateCreateAnimeMovieMetasRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	var DBAM = make([]db.CreateAnimeMovieMetaTxParam, len(req.AnimeMetas))
	for i, am := range req.AnimeMetas {
		DBAM[i] = db.CreateAnimeMovieMetaTxParam{
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

	var PBAM = make([]*pb.AnimeMetaResponse, len(metas.CreateAnimeMovieMetasTxResults))

	for i, am := range metas.CreateAnimeMovieMetasTxResults {
		PBAM[i] = &pb.AnimeMetaResponse{
			Meta:      ConvertMeta(am.Meta),
			Language:  ConvertLanguage(am.Language),
			CreatedAt: timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &pb.CreateAnimeMovieMetasResponse{
		AnimeID:    req.GetAnimeID(),
		AnimeMetas: PBAM,
	}
	return res, nil
}

func validateCreateAnimeMovieMetasRequest(req *pb.CreateAnimeMovieMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, fieldViolation("animeID", err))
	}

	if req.AnimeMetas != nil {
		for _, am := range req.AnimeMetas {
			if err := utils.ValidateInt(int64(am.GetLanguageID())); err != nil {
				violations = append(violations, fieldViolation("languageID", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, fieldViolation("title", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, fieldViolation("overview", err))
			}
		}

	} else {
		violations = append(violations, fieldViolation("animeMetas", errors.New("give at least one metadata")))
	}

	return violations
}
