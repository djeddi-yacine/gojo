package animeSerie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/pb/shpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
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

	cache := &ping.CacheKey{
		ID:     req.AnimeID,
		Target: ping.ANIME_SERIE,
	}

	res := &aspb.GetFullAnimeSerieResponse{}

	server.ping.Handle(ctx, cache.Main(), &res.AnimeSerie, func() error {
		animeSerie, err := server.gojo.GetAnimeSerie(ctx, req.GetAnimeID())
		if err != nil {
			return shared.DatabaseError("failed to get the anime serie", err)
		}

		_, err = server.gojo.GetLanguage(ctx, req.GetLanguageID())
		if err != nil {
			return shared.DatabaseError("failed to get the language", err)
		}

		res.AnimeSerie = shared.ConvertAnimeSerie(animeSerie)
		return nil
	})

	server.ping.Handle(ctx, cache.Meta(uint32(req.LanguageID)), &res.AnimeMeta, func() error {
		_, err = server.gojo.GetLanguage(ctx, req.GetLanguageID())
		if err != nil {
			return shared.DatabaseError("failed to get the language", err)
		}

		animeMeta, err := server.gojo.GetAnimeSerieMeta(ctx, db.GetAnimeSerieMetaParams{
			AnimeID:    req.GetAnimeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shared.DatabaseError("no anime serie found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err := server.gojo.GetMeta(ctx, animeMeta)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shared.DatabaseError("failed to get anime serie metadata", err)
			}

			res.AnimeMeta = &nfpb.AnimeMetaResponse{
				LanguageID: req.GetLanguageID(),
				Meta:       shared.ConvertMeta(meta),
				CreatedAt:  timestamppb.New(meta.CreatedAt),
			}
		}

		return nil
	})

	server.ping.Handle(ctx, cache.Links(), &res.AnimeLinks, func() error {
		animeLinkID, err := server.gojo.GetAnimeSerieLink(ctx, req.GetAnimeID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("failed to get anime serie links ID", err)
		}

		if animeLinkID.AnimeID == req.AnimeID {
			animeLinks, err := server.gojo.GetAnimeLink(ctx, animeLinkID.LinkID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shared.DatabaseError("failed to get anime serie links", err)
			}
			res.AnimeLinks = shared.ConvertAnimeLink(animeLinks)
		}

		return nil
	})

	server.ping.Handle(ctx, cache.Images(), &res.AnimeImages, func() error {
		animePosterIDs, err := server.gojo.ListAnimeSeriePosterImages(ctx, req.AnimeID)
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

		animeBackdropIDs, err := server.gojo.ListAnimeSerieBackdropImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime serie backdrops images IDs", err)
		}

		var animeBackdrops []db.AnimeImage
		if len(animeBackdropIDs) > 0 {
			animeBackdrops = make([]db.AnimeImage, len(animeBackdropIDs))

			for i, p := range animeBackdropIDs {
				backdrop, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime serie backdrop image", err)
				}
				animeBackdrops[i] = backdrop
			}
		}

		animeLogoIDs, err := server.gojo.ListAnimeSerieLogoImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime serie logos images IDs", err)
		}

		var animeLogos []db.AnimeImage
		if len(animeLogoIDs) > 0 {
			animeLogos = make([]db.AnimeImage, len(animeLogoIDs))

			for i, p := range animeLogoIDs {
				logo, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime serie logo image", err)
				}
				animeLogos[i] = logo
			}
		}

		res.AnimeImages = &shpb.AnimeImageResponse{
			Posters:   shared.ConvertAnimeImages(animePosters),
			Backdrops: shared.ConvertAnimeImages(animeBackdrops),
			Logos:     shared.ConvertAnimeImages(animeLogos),
		}

		return nil
	})

	server.ping.Handle(ctx, cache.Trailers(), &res.AnimeTrailers, func() error {
		animeTrailerIDs, err := server.gojo.ListAnimeSerieTrailers(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime serie trailers IDs", err)
		}

		var animeTrailers []db.AnimeTrailer
		if len(animeTrailerIDs) > 0 {
			animeTrailers = make([]db.AnimeTrailer, len(animeTrailerIDs))

			for i, t := range animeTrailerIDs {
				trailer, err := server.gojo.GetAnimeTrailer(ctx, t.TrailerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime serie trailer", err)
				}
				animeTrailers[i] = trailer
			}
		}

		res.AnimeTrailers = shared.ConvertAnimeTrailers(animeTrailers)
		return nil
	})

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
