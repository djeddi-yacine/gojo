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

func (server *AnimeSerieServer) GetFullAnimeSeason(ctx context.Context, req *aspbv1.GetFullAnimeSeasonRequest) (*aspbv1.GetFullAnimeSeasonResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetFullAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.SeasonID,
		Target: ping.AnimeSeason,
	}

	res := &aspbv1.GetFullAnimeSeasonResponse{}

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
		meta, err = server.gojo.GetAnimeSeasonMetaWithLanguageDirectly(ctx, db.GetAnimeSeasonMetaWithLanguageDirectlyParams{
			SeasonID:   req.GetSeasonID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime season found with this language ID", err)
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

	var resources db.AnimeResource
	if err = server.ping.Handle(ctx, cache.Resources(), &resources, func() error {
		resources, err = server.gojo.GetAnimeSeasonResourceDirectly(ctx, req.GetSeasonID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get resources data", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.SeasonResources = av1.ConvertAnimeResource(resources)

	var gIDs []int32
	if err = server.ping.Handle(ctx, cache.Genre(), &gIDs, func() error {
		gIDs, err = server.gojo.ListAnimeSeasonGenres(ctx, req.GetSeasonID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season genres", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	genres, err := server.gojo.ListGenresTx(ctx, gIDs)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code != pgerrcode.CaseNotFound {
				return nil, shv1.ApiError("cannot list anime season genres", err)
			}
		}
	}

	res.SeasonGenres = shv1.ConvertGenres(genres)

	var sIDs []int32
	if err = server.ping.Handle(ctx, cache.Studio(), &sIDs, func() error {
		sIDs, err = server.gojo.ListAnimeSeasonStudios(ctx, req.GetSeasonID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season studios", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	studios, err := server.gojo.ListStudiosTx(ctx, sIDs)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code != pgerrcode.CaseNotFound {
				return nil, shv1.ApiError("cannot list anime season studios", err)
			}
		}
	}

	res.SeasonStudios = shv1.ConvertStudios(studios)

	var tIDs []int64
	if err = server.ping.Handle(ctx, cache.Tags(), &tIDs, func() error {
		tIDs, err = server.gojo.ListAnimeSeasonTags(ctx, req.GetSeasonID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season tags IDs", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	tags, err := server.gojo.ListAnimeTagsTx(ctx, tIDs)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code != pgerrcode.CaseNotFound {
				return nil, shv1.ApiError("cannot get anime season tag", err)
			}
		}
	}

	res.SeasonTags = av1.ConvertAnimeTags(tags)

	var pIDs []int64
	if err = server.ping.Handle(ctx, cache.Posters(), &pIDs, func() error {
		pIDs, err = server.gojo.ListAnimeSeasonPosterImages(ctx, req.GetSeasonID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season posters images IDs", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	posters, err := server.gojo.ListAnimeImagesTx(ctx, pIDs)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code != pgerrcode.CaseNotFound {
				return nil, shv1.ApiError("cannot get anime season posters images", err)
			}
		}
	}

	res.SeasonPosters = av1.ConvertAnimeImages(posters)

	var rIDs []int64
	if err = server.ping.Handle(ctx, cache.Trailers(), &rIDs, func() error {
		rIDs, err = server.gojo.ListAnimeSeasonTrailers(ctx, req.GetSeasonID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season trailers IDs", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	trailers, err := server.gojo.ListAnimeTrailersTx(ctx, rIDs)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code != pgerrcode.CaseNotFound {
				return nil, shv1.ApiError("cannot get anime season trailers", err)
			}
		}
	}

	res.SeasonTrailers = av1.ConvertAnimeTrailers(trailers)

	return res, nil
}

func validateGetFullAnimeSeasonRequest(req *aspbv1.GetFullAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
