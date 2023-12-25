package animeMovie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/pb/shpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) GetAnimeMovieImages(ctx context.Context, req *ampb.GetAnimeMovieImagesRequest) (*ampb.GetAnimeMovieImagesResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get full anime movie")
	}

	violations := validateGetAnimeMovieImagesRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.AnimeID,
		Target: ping.ANIME_MOVIE,
	}

	res := &ampb.GetAnimeMovieImagesResponse{}

	if err = server.ping.Handle(ctx, cache.Images(), &res.AnimeImages, func() error {
		animePosterIDs, err := server.gojo.ListAnimeMoviePosterImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime movie posters images IDs", err)
		}

		var animePosters []db.AnimeImage
		if len(animePosterIDs) > 0 {
			animePosters = make([]db.AnimeImage, len(animePosterIDs))

			for i, p := range animePosterIDs {
				poster, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime movie poster image", err)
				}
				animePosters[i] = poster
			}
		}

		animeBackdropIDs, err := server.gojo.ListAnimeMovieBackdropImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime movie backdrops images IDs", err)
		}

		var animeBackdrops []db.AnimeImage
		if len(animeBackdropIDs) > 0 {
			animeBackdrops = make([]db.AnimeImage, len(animeBackdropIDs))

			for i, b := range animeBackdropIDs {
				backdrop, err := server.gojo.GetAnimeImage(ctx, b)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime movie backdrop image", err)
				}
				animeBackdrops[i] = backdrop
			}
		}

		animeLogoIDs, err := server.gojo.ListAnimeMovieLogoImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shared.DatabaseError("cannot get anime movie logos images IDs", err)
		}

		var animeLogos []db.AnimeImage
		if len(animeLogoIDs) > 0 {
			animeLogos = make([]db.AnimeImage, len(animeLogoIDs))

			for i, l := range animeLogoIDs {
				logo, err := server.gojo.GetAnimeImage(ctx, l)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shared.DatabaseError("cannot get anime movie logo image", err)
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
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func validateGetAnimeMovieImagesRequest(req *ampb.GetAnimeMovieImagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	return violations
}