package amapiv1

import (
	"context"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) GetFullAnimeMovie(ctx context.Context, req *ampbv1.GetFullAnimeMovieRequest) (*ampbv1.GetFullAnimeMovieResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetFullAnimeMovieRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.GetAnimeID(),
		Target: ping.AnimeMovie,
	}

	res := &ampbv1.GetFullAnimeMovieResponse{}

	var movie db.AnimeMovie
	if err = server.ping.Handle(ctx, cache.Main(), &movie, func() error {
		movie, err = server.gojo.GetAnimeMovie(ctx, req.GetAnimeID())
		if err != nil {
			return shv1.ApiError("cannot get anime movie", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeMovie = server.convertAnimeMovie(movie)

	var meta db.Meta
	if err = server.ping.Handle(ctx, cache.Meta(), &meta, func() error {
		meta, err = server.gojo.GetAnimeMovieMetaWithLanguageDirectly(ctx, db.GetAnimeMovieMetaWithLanguageDirectlyParams{
			AnimeID:    req.GetAnimeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime movie found with this language ID", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeMeta = &nfpbv1.AnimeMetaResponse{
		LanguageID: req.GetLanguageID(),
		Meta:       shv1.ConvertMeta(meta),
		CreatedAt:  timestamppb.New(meta.CreatedAt),
	}

	var resources db.AnimeResource
	if err = server.ping.Handle(ctx, cache.Resources(), &resources, func() error {
		resources, err = server.gojo.GetAnimeMovieResourceDirectly(ctx, req.GetAnimeID())
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

	res.AnimeResources = av1.ConvertAnimeResource(resources)

	var gIDs []int32
	if err = server.ping.Handle(ctx, cache.Genre(), &gIDs, func() error {
		gIDs, err = server.gojo.ListAnimeMovieGenres(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie genres", err)
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
				return nil, shv1.ApiError("cannot list anime movie genres", err)
			}
		}
	}

	res.AnimeGenres = shv1.ConvertGenres(genres)

	var sIDs []int32
	if err = server.ping.Handle(ctx, cache.Studio(), &sIDs, func() error {
		sIDs, err = server.gojo.ListAnimeMovieStudios(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie studios", err)
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
				return nil, shv1.ApiError("cannot list anime movie studios", err)
			}
		}
	}

	res.AnimeStudios = shv1.ConvertStudios(studios)

	var serverID int64
	if err = server.ping.Handle(ctx, cache.Server(), &serverID, func() error {
		serverID, err = server.gojo.GetAnimeMovieServerByAnimeID(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie server ID", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.ServerID = serverID

	if serverID != 0 {
		var subVideos []db.AnimeMovieVideo
		if err = server.ping.Handle(ctx, cache.SubV(), &subVideos, func() error {
			v, err := server.gojo.ListAnimeMovieServerSubVideos(ctx, res.ServerID)
			if err != nil {
				if dberr := db.ErrorDB(err); dberr != nil {
					if dberr.Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot list anime movie server sub videos", err)
					}
				}
			}

			if len(v) > 0 {
				subVideos = make([]db.AnimeMovieVideo, len(v))
				for i, x := range v {
					subVideos[i], err = server.gojo.GetAnimeMovieVideo(ctx, x.VideoID)
					if err != nil {
						if dberr := db.ErrorDB(err); dberr != nil {
							if dberr.Code != pgerrcode.CaseNotFound {
								return shv1.ApiError("cannot get anime movie server sub videos", err)
							}
						}
					}
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		var subTorrents []db.AnimeMovieTorrent
		if err = server.ping.Handle(ctx, cache.SubT(), &subTorrents, func() error {
			v, err := server.gojo.ListAnimeMovieServerSubTorrents(ctx, res.ServerID)
			if err != nil {
				if dberr := db.ErrorDB(err); dberr != nil {
					if dberr.Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot list anime movie server sub torrents", err)
					}
				}
			}

			if len(v) > 0 {
				subTorrents = make([]db.AnimeMovieTorrent, len(v))
				for i, x := range v {
					subTorrents[i], err = server.gojo.GetAnimeMovieTorrent(ctx, x.TorrentID)
					if err != nil {
						if dberr := db.ErrorDB(err); dberr != nil {
							if dberr.Code != pgerrcode.CaseNotFound {
								return shv1.ApiError("cannot get anime movie server sub torrents", err)
							}
						}
					}
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.Sub = &apbv1.AnimeSubDataResponse{
			Videos:   server.convertAnimeMovieVideos(subVideos),
			Torrents: server.convertAnimeMovieTorrents(subTorrents),
		}

		var dubVideos []db.AnimeMovieVideo
		if err = server.ping.Handle(ctx, cache.DubV(), &dubVideos, func() error {
			v, err := server.gojo.ListAnimeMovieServerDubVideos(ctx, res.ServerID)
			if err != nil {
				if dberr := db.ErrorDB(err); dberr != nil {
					if dberr.Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot list anime movie server dub videos", err)
					}
				}
			}

			if len(v) > 0 {
				dubVideos = make([]db.AnimeMovieVideo, len(v))
				for i, x := range v {
					dubVideos[i], err = server.gojo.GetAnimeMovieVideo(ctx, x.VideoID)
					if err != nil {
						if dberr := db.ErrorDB(err); dberr != nil {
							if dberr.Code != pgerrcode.CaseNotFound {
								return shv1.ApiError("cannot get anime movie server dub videos", err)
							}
						}
					}
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		var dubTorrents []db.AnimeMovieTorrent
		if err = server.ping.Handle(ctx, cache.DubT(), &dubTorrents, func() error {
			v, err := server.gojo.ListAnimeMovieServerDubTorrents(ctx, res.ServerID)
			if err != nil {
				if dberr := db.ErrorDB(err); dberr != nil {
					if dberr.Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot list anime movie server dub torrents", err)
					}
				}
			}

			if len(v) > 0 {
				dubTorrents = make([]db.AnimeMovieTorrent, len(v))
				for i, x := range v {
					dubTorrents[i], err = server.gojo.GetAnimeMovieTorrent(ctx, x.TorrentID)
					if err != nil {
						if dberr := db.ErrorDB(err); dberr != nil {
							if dberr.Code != pgerrcode.CaseNotFound {
								return shv1.ApiError("cannot get anime movie server dub torrents", err)
							}
						}
					}
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.Dub = &apbv1.AnimeDubDataResponse{
			Videos:   server.convertAnimeMovieVideos(dubVideos),
			Torrents: server.convertAnimeMovieTorrents(dubTorrents),
		}
	}

	var link db.AnimeLink
	if err = server.ping.Handle(ctx, cache.Links(), &link, func() error {
		link, err = server.gojo.GetAnimeMovieLinksDirectly(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie links", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeLinks = av1.ConvertAnimeLink(link)

	var tIDs []int64
	if err = server.ping.Handle(ctx, cache.Tags(), &tIDs, func() error {
		tIDs, err = server.gojo.ListAnimeMovieTags(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie tags IDs", err)
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
				return nil, shv1.ApiError("cannot get anime movie tag", err)
			}
		}
	}

	res.AnimeTags = av1.ConvertAnimeTags(tags)

	var pIDs []int64
	if err = server.ping.Handle(ctx, cache.Posters(), &pIDs, func() error {
		pIDs, err = server.gojo.ListAnimeMoviePosterImages(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie posters images IDs", err)
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
				return nil, shv1.ApiError("cannot get anime movie posters images", err)
			}
		}
	}

	var bIDs []int64
	if err = server.ping.Handle(ctx, cache.Backdrops(), &bIDs, func() error {
		bIDs, err = server.gojo.ListAnimeMovieBackdropImages(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie backdrops images IDs", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	backdrops, err := server.gojo.ListAnimeImagesTx(ctx, bIDs)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code != pgerrcode.CaseNotFound {
				return nil, shv1.ApiError("cannot get anime movie backdrops images", err)
			}
		}
	}

	var lIDs []int64
	if err = server.ping.Handle(ctx, cache.Logos(), &lIDs, func() error {
		lIDs, err = server.gojo.ListAnimeMovieLogoImages(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie logos images IDs", err)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	logos, err := server.gojo.ListAnimeImagesTx(ctx, lIDs)
	if err != nil {
		if dberr := db.ErrorDB(err); dberr != nil {
			if dberr.Code != pgerrcode.CaseNotFound {
				return nil, shv1.ApiError("cannot get anime movie logos images", err)
			}
		}
	}

	res.AnimeImages = &apbv1.AnimeImageResponse{
		Posters:   av1.ConvertAnimeImages(posters),
		Backdrops: av1.ConvertAnimeImages(backdrops),
		Logos:     av1.ConvertAnimeImages(logos),
	}

	var rIDs []int64
	if err = server.ping.Handle(ctx, cache.Trailers(), &rIDs, func() error {
		rIDs, err = server.gojo.ListAnimeMovieTrailers(ctx, req.GetAnimeID())
		if err != nil {
			if dberr := db.ErrorDB(err); dberr != nil {
				if dberr.Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie trailers IDs", err)
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
				return nil, shv1.ApiError("cannot get anime movie trailers", err)
			}
		}
	}

	res.AnimeTrailers = av1.ConvertAnimeTrailers(trailers)

	return res, nil
}

func validateGetFullAnimeMovieRequest(req *ampbv1.GetFullAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
