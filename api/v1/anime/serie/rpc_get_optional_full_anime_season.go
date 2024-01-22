package asapiv1

import (
	"context"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) GetOptionalFullAnimeSeason(ctx context.Context, req *aspbv1.GetOptionalFullAnimeSeasonRequest) (*aspbv1.GetOptionalFullAnimeSeasonResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetOptionalFullAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.SeasonID,
		Target: ping.AnimeSeason,
	}

	res := &aspbv1.GetOptionalFullAnimeSeasonResponse{}

	var season db.AnimeSeason
	if err = server.ping.Handle(ctx, cache.Main(), &season, func() error {
		season, err = server.gojo.GetAnimeSeason(ctx, req.GetSeasonID())
		if err != nil {
			return shv1.ApiError("cannot get anime season", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeSeason = convertAnimeSeason(season)

	var meta db.Meta
	if err = server.ping.Handle(ctx, cache.Meta(), &meta, func() error {
		animeMeta, err := server.gojo.GetAnimeSeasonMeta(ctx, db.GetAnimeSeasonMetaParams{
			SeasonID:   req.GetSeasonID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime season found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err = server.gojo.GetMeta(ctx, animeMeta)
			if err != nil {
				return shv1.ApiError("cannot get anime season metadata", err)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	res.SeasonMeta = &nfpbv1.AnimeMetaResponse{
		LanguageID: req.GetLanguageID(),
		Meta:       shv1.ConvertMeta(meta),
		CreatedAt:  timestamppb.New(meta.CreatedAt),
	}

	if req.GetWithResources() {
		var resources db.AnimeResource
		if err = server.ping.Handle(ctx, cache.Resources(), &resources, func() error {
			ID, err := server.gojo.GetAnimeSeasonResource(ctx, req.GetSeasonID())
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season resources", err)
				} else {
					return nil
				}
			}

			resources, err = server.gojo.GetAnimeResource(ctx, ID.ResourceID)
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get resources data", err)
				} else {
					return nil
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.SeasonResources = av1.ConvertAnimeResource(resources)
	}

	if req.GetWithGenres() {
		var gIDs []int32
		if err = server.ping.Handle(ctx, cache.Genre(), &gIDs, func() error {
			gIDs, err = server.gojo.ListAnimeSeasonGenres(ctx, req.GetSeasonID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime season genres", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		genres, err := server.gojo.ListGenresTx(ctx, gIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot list anime season genres", err)
		}

		res.SeasonGenres = shv1.ConvertGenres(genres)
	}

	if req.GetWithStudios() {
		var sIDs []int32
		if err = server.ping.Handle(ctx, cache.Studio(), &sIDs, func() error {
			sIDs, err = server.gojo.ListAnimeSeasonStudios(ctx, req.GetSeasonID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime season studios", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		studios, err := server.gojo.ListStudiosTx(ctx, sIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot list anime season studios", err)
		}

		res.SeasonStudios = shv1.ConvertStudios(studios)
	}

	if req.GetWithTags() {
		var tIDs []int64
		if err = server.ping.Handle(ctx, cache.Tags(), &tIDs, func() error {
			tIDs, err = server.gojo.ListAnimeSeasonTags(ctx, req.GetSeasonID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime season tags IDs", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		tags, err := server.gojo.ListAnimeTagsTx(ctx, tIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot get anime season tag", err)
		}

		res.SeasonTags = av1.ConvertAnimeTags(tags)
	}

	if req.GetWithPosters() {
		var pIDs []int64
		if err = server.ping.Handle(ctx, cache.Posters(), &pIDs, func() error {
			pIDs, err = server.gojo.ListAnimeSeasonPosterImages(ctx, req.GetSeasonID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime season posters images IDs", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		posters, err := server.gojo.ListAnimeImagesTx(ctx, pIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot get anime season posters images", err)
		}

		res.SeasonPosters = av1.ConvertAnimeImages(posters)
	}

	if req.GetWithTrailer() {
		var rIDs []int64
		if err = server.ping.Handle(ctx, cache.Trailers(), &rIDs, func() error {
			rIDs, err = server.gojo.ListAnimeSeasonTrailers(ctx, req.GetSeasonID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime season trailers IDs", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		trailers, err := server.gojo.ListAnimeTrailersTx(ctx, rIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot get anime season trailers", err)
		}

		res.SeasonTrailers = av1.ConvertAnimeTrailers(trailers)
	}

	return res, nil
}

func validateGetOptionalFullAnimeSeasonRequest(req *aspbv1.GetOptionalFullAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
