package animeSerie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) GetFullAnimeSeason(ctx context.Context, req *aspb.GetFullAnimeSeasonRequest) (*aspb.GetFullAnimeSeasonResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get full anime serie")
	}

	violations := validateGetFullAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	cache := &CacheKey{
		id:     req.SeasonID,
		target: SEASON_KEY,
	}

	res := &aspb.GetFullAnimeSeasonResponse{}

	server.do(ctx, cache.Anime(), &res.AnimeSeason, func() error {
		animeSeason, err := server.gojo.GetAnimeSeason(ctx, req.GetSeasonID())
		if err != nil {
			return shared.DatabaseError("failed to get the anime season", err)
		}

		res.AnimeSeason = shared.ConvertAnimeSeason(animeSeason)
		return nil
	})

	server.do(ctx, cache.Meta(uint32(req.LanguageID)), &res.SeasonMeta, func() error {
		_, err = server.gojo.GetLanguage(ctx, req.GetLanguageID())
		if err != nil {
			return shared.DatabaseError("failed to get the language", err)
		}

		animeMeta, err := server.gojo.GetAnimeSeasonMeta(ctx, db.GetAnimeSeasonMetaParams{
			SeasonID:   req.GetSeasonID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shared.DatabaseError("no anime season found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err := server.gojo.GetMeta(ctx, animeMeta)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shared.DatabaseError("failed to get anime season metadata", err)
			}

			res.SeasonMeta = &nfpb.AnimeMetaResponse{
				LanguageID: req.GetLanguageID(),
				Meta:       shared.ConvertMeta(meta),
				CreatedAt:  timestamppb.New(meta.CreatedAt),
			}
		}

		return nil
	})

	server.do(ctx, cache.Genre(), &res.SeasonGenres, func() error {
		seasonGenres, err := server.gojo.ListAnimeSeasonGenres(ctx, req.GetSeasonID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("failed to get anime serie genres", err)
		}

		if len(seasonGenres) > 0 {
			genres := make([]db.Genre, len(seasonGenres))

			for i, amg := range seasonGenres {
				genres[i], err = server.gojo.GetGenre(ctx, amg)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("failed when list anime serie genres", err)
				}
			}
			res.SeasonGenres = shared.ConvertGenres(genres)
		}
		return nil
	})

	server.do(ctx, cache.Studio(), &res.SeasonStudios, func() error {
		seasonStudios, err := server.gojo.ListAnimeSeasonStudios(ctx, req.GetSeasonID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("failed to get anime serie studios", err)
		}

		if len(seasonStudios) > 0 {
			studios := make([]db.Studio, len(seasonStudios))
			for i, ams := range seasonStudios {
				studios[i], err = server.gojo.GetStudio(ctx, ams)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("failed when list anime serie studios", err)
				}
			}
			res.SeasonStudios = shared.ConvertStudios(studios)
		}
		return nil
	})

	server.do(ctx, cache.Resources(), &res.SeasonResoures, func() error {
		seasonResourceID, err := server.gojo.GetAnimeSeasonResource(ctx, req.GetSeasonID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime season resources", err)
		}

		seasonResources, err := server.gojo.GetAnimeResource(ctx, seasonResourceID.ResourceID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get resources data", err)
		}
		res.SeasonResoures = shared.ConvertAnimeResource(seasonResources)
		return nil
	})

	server.do(ctx, cache.Tags(), &res.SeasonTags, func() error {
		animeTagIDs, err := server.gojo.ListAnimeSeasonTags(ctx, req.SeasonID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime season tags IDs", err)
		}

		var animeTags []db.AnimeTag
		if len(animeTagIDs) > 0 {
			animeTags = make([]db.AnimeTag, len(animeTagIDs))

			for i, t := range animeTagIDs {
				tag, err := server.gojo.GetAnimeTag(ctx, t.TagID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime season tag", err)
				}
				animeTags[i] = tag
			}
		}

		if len(animeTags) > 0 {
			res.SeasonTags = make([]*aspb.AnimeSeasonTag, len(animeTags))

			for i, t := range animeTags {
				res.SeasonTags[i] = &aspb.AnimeSeasonTag{
					ID:        t.ID,
					Tag:       t.Tag,
					CreatedAt: timestamppb.New(t.CreatedAt),
				}
			}
		}

		return nil
	})

	server.do(ctx, cache.Images(), &res.SeasonPosters, func() error {
		animePosterIDs, err := server.gojo.ListAnimeSeriePosterImages(ctx, req.SeasonID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime serie posters images IDs", err)
		}

		var animePosters []db.AnimeImage
		if len(animePosterIDs) > 0 {
			animePosters = make([]db.AnimeImage, len(animePosterIDs))

			for i, p := range animePosterIDs {
				poster, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime serie poster image", err)
				}
				animePosters[i] = poster
			}
		}

		res.SeasonPosters = shared.ConvertAnimeImages(animePosters)

		return nil
	})

	server.do(ctx, cache.Trailers(), &res.SeasonTrailers, func() error {
		animeTrailerIDs, err := server.gojo.ListAnimeSerieTrailers(ctx, req.SeasonID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime season trailers IDs", err)
		}

		var animeTrailers []db.AnimeTrailer
		if len(animeTrailerIDs) > 0 {
			animeTrailers = make([]db.AnimeTrailer, len(animeTrailerIDs))

			for i, t := range animeTrailerIDs {
				trailer, err := server.gojo.GetAnimeTrailer(ctx, t.TrailerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime season trailer", err)
				}
				animeTrailers[i] = trailer
			}
		}

		res.SeasonTrailers = shared.ConvertAnimeTrailers(animeTrailers)
		return nil
	})

	return res, nil
}

func validateGetFullAnimeSeasonRequest(req *aspb.GetFullAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shared.FieldViolation("seasonID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shared.FieldViolation("languageID", err))
	}

	return violations
}
