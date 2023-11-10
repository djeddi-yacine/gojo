package animeSerie

import (
	"context"
	"math"

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

func (server *AnimeSerieServer) GetFullAnimeSerie(ctx context.Context, req *aspb.GetFullAnimeSerieRequest) (*aspb.GetFullAnimeSerieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get full anime serie")
	}

	violations := validateGetFullAnimeSerieRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	animeSerie, err := server.gojo.GetAnimeSerie(ctx, req.GetAnimeID())
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.NotFound, "there is no anime serie with this ID : %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get the anime serie : %s", err)
	}

	_, err = server.gojo.GetLanguage(ctx, req.GetLanguageID())
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.NotFound, "there is no language with this ID : %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get the language : %s", err)
	}

	res := &aspb.GetFullAnimeSerieResponse{
		AnimeSerie: shared.ConvertAnimeSerie(animeSerie),
	}

	animeMeta, err := server.gojo.GetAnimeSerieMeta(ctx, db.GetAnimeSerieMetaParams{
		AnimeID:    req.GetAnimeID(),
		LanguageID: req.GetLanguageID(),
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no anime serie found with this languageID : %s", err)
	}

	if animeMeta > 0 {
		meta, err := server.gojo.GetMeta(ctx, animeMeta)
		if err != nil && err != db.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "error when return anime serie metadata : %s", err)
		}

		res.AnimeMeta = &nfpb.AnimeMetaResponse{
			LanguageID: req.GetLanguageID(),
			Meta:       shared.ConvertMeta(meta),
			CreatedAt:  timestamppb.New(meta.CreatedAt),
		}
	}

	animeSerieResources, err := server.gojo.GetAnimeSerieResource(ctx, req.GetAnimeID())
	if err != nil && err != db.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "error when get anime serie serie resources : %s", err)
	}

	if animeSerieResources.AnimeID == req.AnimeID {
		animeResources, err := server.gojo.GetAnimeResource(ctx, animeSerieResources.ResourceID)
		if err != nil && err != db.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "error when get anime serie resources : %s", err)
		}
		res.AnimeResources = shared.ConvertAnimeResource(animeResources)
	}

	animeSerieGenres, err := server.gojo.ListAnimeSerieGenres(ctx, req.GetAnimeID())
	if err != nil && err != db.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "error when get anime serie genres : %s", err)
	}

	if len(animeSerieGenres) > 0 {
		genres := make([]db.Genre, len(animeSerieGenres))

		for i, amg := range animeSerieGenres {
			genres[i], err = server.gojo.GetGenre(ctx, amg)
			if err != nil && err != db.ErrRecordNotFound {
				return nil, status.Errorf(codes.Internal, "error when list anime serie genres : %s", err)
			}
		}
		res.AnimeGenres = shared.ConvertGenres(genres)
	}

	animeSerieStudios, err := server.gojo.ListAnimeSerieStudios(ctx, req.GetAnimeID())
	if err != nil && err != db.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "error when get anime serie studios : %s", err)
	}

	if len(animeSerieStudios) > 0 {
		studios := make([]db.Studio, len(animeSerieStudios))
		for i, ams := range animeSerieStudios {
			studios[i], err = server.gojo.GetStudio(ctx, ams)
			if err != nil && err != db.ErrRecordNotFound {
				return nil, status.Errorf(codes.Internal, "error when list anime serie studios : %s", err)
			}
		}
		res.AnimeStudios = shared.ConvertStudios(studios)
	}

	animeSerieSeasons, err := server.gojo.ListAnimeSerieSeasonsByAnimeID(ctx, db.ListAnimeSerieSeasonsByAnimeIDParams{
		AnimeID: req.GetAnimeID(),
		Limit:   math.MaxInt32,
		Offset:  0,
	})
	if err != nil && err != db.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "error when get anime serie studios : %s", err)
	}

	if len(animeSerieSeasons) > 0 {
		var meta db.AnimeSerieSeasonMeta
		var seasonMeta db.Meta
		var seasonEpisodes []db.AnimeSerieEpisode
		seasonDetails := make([]*aspb.AnimeSerieSeasonDetails, len(animeSerieSeasons))

		for i, assm := range animeSerieSeasons {
			meta, err = server.gojo.GetAnimeSerieSeasonMeta(ctx, db.GetAnimeSerieSeasonMetaParams{
				SeasonID:   assm.ID,
				LanguageID: req.GetLanguageID(),
			})
			if err != nil && err != db.ErrRecordNotFound {
				return nil, status.Errorf(codes.Internal, "error when list anime serie season metadata : %s", err)
			}

			seasonMeta, err = server.gojo.GetMeta(ctx, meta.MetaID)
			if err != nil && err != db.ErrRecordNotFound {
				return nil, status.Errorf(codes.Internal, "error when get anime serie season metadata : %s", err)
			}

			seasonEpisodes, err = server.gojo.ListAnimeSerieEpisodesBySeasonID(ctx, db.ListAnimeSerieEpisodesBySeasonIDParams{
				SeasonID: assm.ID,
				Limit:    math.MaxInt32,
				Offset:   0,
			})
			if err != nil && err != db.ErrRecordNotFound {
				return nil, status.Errorf(codes.Internal, "error when list anime serie season episodes : %s", err)
			}

			seasonDetails[i] = &aspb.AnimeSerieSeasonDetails{
				Season: shared.ConvertAnimeSerieSeason(assm),
				SeasonMetas: &nfpb.AnimeMetaResponse{
					LanguageID: req.GetLanguageID(),
					Meta:       shared.ConvertMeta(seasonMeta),
					CreatedAt:  timestamppb.New(seasonMeta.CreatedAt),
				},
				SeasonEpisodes: shared.ConvertAnimeMovieEpiosdes(seasonEpisodes),
			}
		}
		res.AnimeSeasons = &aspb.AnimeSerieSeasonDetailsRsponse{
			Seasons: seasonDetails,
		}
	}

	return res, nil
}

func validateGetFullAnimeSerieRequest(req *aspb.GetFullAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shared.FieldViolation("languageID", err))
	}

	return violations
}
