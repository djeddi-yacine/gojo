package animeMovie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/pb/shpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) GetFullAnimeMovie(ctx context.Context, req *ampb.GetFullAnimeMovieRequest) (*ampb.GetFullAnimeMovieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get full anime movie")
	}

	violations := validateGetFullAnimeMovieRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	animeMovie, err := server.gojo.GetAnimeMovie(ctx, req.GetAnimeID())
	if err != nil {
		return nil, shared.DatabaseError("cannot get anime movie", err)
	}

	_, err = server.gojo.GetLanguage(ctx, req.GetLanguageID())
	if err != nil {
		return nil, shared.DatabaseError("no language found with this language ID", err)
	}

	res := &ampb.GetFullAnimeMovieResponse{
		AnimeMovie: shared.ConvertAnimeMovie(animeMovie),
	}

	animeMeta, err := server.gojo.GetAnimeMovieMeta(ctx, db.GetAnimeMovieMetaParams{
		AnimeID:    req.GetAnimeID(),
		LanguageID: req.GetLanguageID(),
	})
	if err != nil {
		return nil, shared.DatabaseError("no anime movie found with this language ID", err)
	}

	if animeMeta > 0 {
		meta, err := server.gojo.GetMeta(ctx, animeMeta)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("cannot get anime movie metadata", err)
		}

		res.AnimeMeta = &nfpb.AnimeMetaResponse{
			LanguageID: req.GetLanguageID(),
			Meta:       shared.ConvertMeta(meta),
			CreatedAt:  timestamppb.New(meta.CreatedAt),
		}
	}

	animeMovieResources, err := server.gojo.GetAnimeMovieResource(ctx, req.GetAnimeID())
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie resources", err)
	}

	animeResources, err := server.gojo.GetAnimeResource(ctx, animeMovieResources.ResourceID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get resources data", err)
	}
	res.AnimeResources = shared.ConvertAnimeResource(animeResources)

	animeMovieGenres, err := server.gojo.ListAnimeMovieGenres(ctx, req.GetAnimeID())
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie genres", err)
	}

	if len(animeMovieGenres) > 0 {
		genres := make([]db.Genre, len(animeMovieGenres))

		for i, amg := range animeMovieGenres {
			genres[i], err = server.gojo.GetGenre(ctx, amg)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot list anime movie genres", err)
			}
		}
		res.AnimeGenres = shared.ConvertGenres(genres)
	}

	animeMovieStudios, err := server.gojo.ListAnimeMovieStudios(ctx, req.GetAnimeID())
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie studios", err)
	}

	if len(animeMovieStudios) > 0 {
		studios := make([]db.Studio, len(animeMovieStudios))
		for i, ams := range animeMovieStudios {
			studios[i], err = server.gojo.GetStudio(ctx, ams)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot list anime movie studios", err)
			}
		}
		res.AnimeStudios = shared.ConvertStudios(studios)
	}

	var sv db.AnimeMovieServer
	sv, err = server.gojo.GetAnimeMovieServer(ctx, req.GetAnimeID())
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie server ID", err)
	}

	if sv.AnimeID == req.GetAnimeID() {
		res.ServerID = &sv.ID
		ss, err := server.gojo.ListAnimeMovieServerSubVideos(ctx, sv.ID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("cannot list anime movie server sub videos", err)
		}

		subV := make([]db.AnimeMovieVideo, len(ss))
		for i, ks := range ss {
			subV[i], err = server.gojo.GetAnimeMovieVideo(ctx, ks.VideoID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime movie server sub videos", err)
			}
		}

		st, err := server.gojo.ListAnimeMovieServerSubTorrents(ctx, sv.ID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("cannot list anime movie server sub torrents", err)
		}

		subT := make([]db.AnimeMovieTorrent, len(st))
		for i, kst := range st {
			subT[i], err = server.gojo.GetAnimeMovieTorrent(ctx, kst.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime movie server sub torrents", err)
			}
		}

		res.Sub = &ampb.AnimeMovieSubDataResponse{
			Videos:   shared.ConvertAnimeMovieVideos(subV),
			Torrents: shared.ConvertAnimeMovieTorrents(subT),
		}

		sd, err := server.gojo.ListAnimeMovieServerDubVideos(ctx, sv.ID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("cannot list anime movie server dub videos", err)
		}

		dubV := make([]db.AnimeMovieVideo, len(sd))
		for i, kd := range sd {
			dubV[i], err = server.gojo.GetAnimeMovieVideo(ctx, kd.VideoID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime movie server dub videos", err)
			}
		}

		dt, err := server.gojo.ListAnimeMovieServerDubTorrents(ctx, sv.ID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("cannot list anime movie server dub torrents", err)
		}

		dubT := make([]db.AnimeMovieTorrent, len(dt))
		for i, kdt := range dt {
			subT[i], err = server.gojo.GetAnimeMovieTorrent(ctx, kdt.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime movie server dub torrents", err)
			}
		}

		res.Dub = &ampb.AnimeMovieDubDataResponse{
			Videos:   shared.ConvertAnimeMovieVideos(dubV),
			Torrents: shared.ConvertAnimeMovieTorrents(dubT),
		}
	}

	animePosterIDs, err := server.gojo.ListAnimeMoviePosterImages(ctx, req.AnimeID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie posters images IDs", err)
	}

	var animePosters []db.AnimeImage
	if len(animePosterIDs) > 0 {
		animePosters = make([]db.AnimeImage, len(animePosterIDs))

		for i, p := range animePosterIDs {
			poster, err := server.gojo.GetAnimeImage(ctx, p)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime movie poster image", err)
			}
			animePosters[i] = poster
		}
	}

	animeBackdropIDs, err := server.gojo.ListAnimeMovieBackdropImages(ctx, req.AnimeID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie backdrops images IDs", err)
	}

	var animeBackdrops []db.AnimeImage
	if len(animeBackdropIDs) > 0 {
		animeBackdrops = make([]db.AnimeImage, len(animeBackdropIDs))

		for i, p := range animeBackdropIDs {
			backdrop, err := server.gojo.GetAnimeImage(ctx, p)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime movie backdrop image", err)
			}
			animeBackdrops[i] = backdrop
		}
	}

	animeLogoIDs, err := server.gojo.ListAnimeMovieLogoImages(ctx, req.AnimeID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie logos images IDs", err)
	}

	var animeLogos []db.AnimeImage
	if len(animeLogoIDs) > 0 {
		animeLogos = make([]db.AnimeImage, len(animeLogoIDs))

		for i, p := range animeLogoIDs {
			logo, err := server.gojo.GetAnimeImage(ctx, p)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime movie logo image", err)
			}
			animeLogos[i] = logo
		}
	}

	res.AnimeImages = &shpb.AnimeImageResponse{
		Posters:   shared.ConvertAnimeImages(animePosters),
		Backdrops: shared.ConvertAnimeImages(animeBackdrops),
		Logos:     shared.ConvertAnimeImages(animeLogos),
	}

	animeLinkID, err := server.gojo.GetAnimeMovieLink(ctx, req.AnimeID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime movie logos images IDs", err)
	}

	if animeLinkID.AnimeID == req.AnimeID {
		animeLink, err := server.gojo.GetAnimeLink(ctx, animeLinkID.ID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("cannot get anime movie links", err)
		}

		res.AnimeLinks = shared.ConvertAnimeLink(animeLink)
	}

	return res, nil
}

func validateGetFullAnimeMovieRequest(req *ampb.GetFullAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("animeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shared.FieldViolation("languageID", err))
	}

	return violations
}
