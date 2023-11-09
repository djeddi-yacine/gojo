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

func (server *AnimeSerieServer) CreateAnimeSerieEpisode(ctx context.Context, req *aspb.CreateAnimeSerieEpisodeRequest) (*aspb.CreateAnimeSerieEpisodeResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie episode")
	}

	if violations := validateCreateAnimeSerieEpisodeRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
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

	arg := db.CreateAnimeSerieEpisodeTxParams{
		Episode: db.CreateAnimeSerieEpisodeParams{
			SeasonID:           req.GetEpisode().GetSeasonID(),
			EpisodeNumber:      req.GetEpisode().GetEpisodeNumber(),
			Thumbnails:         req.GetEpisode().GetThumbnails(),
			ThumbnailsBlurHash: req.GetEpisode().GetThumbnailsBlurHash(),
		},
		EpisodeMetas: DBEM,
	}

	data, err := server.gojo.CreateAnimeSerieEpisodeTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie episode : %s", err)
	}

	var PBSM = make([]*nfpb.AnimeMetaResponse, len(data.AnimeSerieEpisodeMetas))

	for i, am := range data.AnimeSerieEpisodeMetas {
		PBSM[i] = &nfpb.AnimeMetaResponse{
			Meta:      shared.ConvertMeta(am.Meta),
			Language:  shared.ConvertLanguage(am.Language),
			CreatedAt: timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspb.CreateAnimeSerieEpisodeResponse{
		Season:       shared.ConvertAnimeSerieSeason(data.AnimeSerieSeason),
		Episode:      shared.ConvertAnimeSerieEpisode(data.AnimeSerieEpisode),
		EpisodeMetas: PBSM,
	}
	return res, nil
}

func validateCreateAnimeSerieEpisodeRequest(req *aspb.CreateAnimeSerieEpisodeRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.GetEpisode() != nil {
		if err := utils.ValidateInt(req.GetEpisode().GetSeasonID()); err != nil {
			violations = append(violations, shared.FieldViolation("seasonID", err))
		}

		if err := utils.ValidateInt(int64(req.GetEpisode().GetEpisodeNumber())); err != nil {
			violations = append(violations, shared.FieldViolation("episodeNumber", err))
		}

		if err := utils.ValidateImage(req.GetEpisode().GetThumbnails()); err != nil {
			violations = append(violations, shared.FieldViolation("thumbnails", err))
		}

		if err := utils.ValidateString(req.GetEpisode().GetThumbnailsBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("thumbnailsBlurHash", err))
		}

	} else {
		violations = append(violations, shared.FieldViolation("episode", errors.New("episode :you need to send the episode model")))
	}

	if req.EpisodeMetas != nil {
		for _, am := range req.EpisodeMetas {
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
		violations = append(violations, shared.FieldViolation("episodeMetas", errors.New("episodeMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
