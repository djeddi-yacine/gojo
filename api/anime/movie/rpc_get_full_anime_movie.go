package animeMovie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
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
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.NotFound, "there is no anime movie with this ID : %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get the anime movie : %s", err)
	}

	_, err = server.gojo.GetLanguage(ctx, req.GetLanguageID())
	if err != nil {
		if db.ErrorCode(err) == db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.NotFound, "there is no language with this ID : %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get the language : %s", err)
	}

	res := &ampb.GetFullAnimeMovieResponse{
		AnimeMovie: shared.ConvertAnimeMovie(animeMovie),
	}

	animeMeta, err := server.gojo.GetAnimeMovieMeta(ctx, db.GetAnimeMovieMetaParams{
		AnimeID:    req.GetAnimeID(),
		LanguageID: req.GetLanguageID(),
	})
	if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
		return nil, status.Errorf(codes.Internal, "no anime movie found with this languageID : %s", err)
	}

	if animeMeta != 0 {
		meta, err := server.gojo.GetMeta(ctx, animeMeta)
		if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.Internal, "error when return anime movie metadata : %s", err)
		}

		res.AnimeMeta = &nfpb.AnimeMetaResponse{
			LanguageID: req.GetLanguageID(),
			Meta:       shared.ConvertMeta(meta),
			CreatedAt:  timestamppb.New(meta.CreatedAt),
		}
	}

	animeMovieResources, err := server.gojo.GetAnimeMovieResource(ctx, req.GetAnimeID())
	if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
		return nil, status.Errorf(codes.Internal, "error when get anime movie movie resources : %s", err)
	}

	if animeMovieResources.ID != 0 {
		animeResources, err := server.gojo.GetAnimeResource(ctx, animeMovieResources.ResourceID)
		if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.Internal, "error when get anime movie resources : %s", err)
		}
		res.AnimeResources = shared.ConvertAnimeResource(animeResources)
	}

	animeMovieGenres, err := server.gojo.ListAnimeMovieGenres(ctx, req.GetAnimeID())
	if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
		return nil, status.Errorf(codes.Internal, "error when get anime movie genres : %s", err)
	}

	if animeMovieGenres != nil {
		genres := make([]db.Genre, len(animeMovieGenres))

		for i, amg := range animeMovieGenres {
			genres[i], err = server.gojo.GetGenre(ctx, amg)
			if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
				return nil, status.Errorf(codes.Internal, "error when list anime movie genres : %s", err)
			}
		}
		res.AnimeGenres = shared.ConvertGenres(genres)
	}

	animeMovieStudios, err := server.gojo.ListAnimeMovieStudios(ctx, req.GetAnimeID())
	if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
		return nil, status.Errorf(codes.Internal, "error when get anime movie studios : %s", err)
	}

	if animeMovieStudios != nil {
		studios := make([]db.Studio, len(animeMovieStudios))
		for i, ams := range animeMovieStudios {
			studios[i], err = server.gojo.GetStudio(ctx, ams)
			if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
				return nil, status.Errorf(codes.Internal, "error when list anime movie studios : %s", err)
			}
		}
		res.AnimeStudios = shared.ConvertStudios(studios)
	}

	var sv db.AnimeMovieServer
	sv, err = server.gojo.GetAnimeMovieServer(ctx, req.GetAnimeID())
	if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
		return nil, status.Errorf(codes.Internal, "error when get anime movie server ID : %s", err)
	}

	if sv.ID != 0 {
		res.ServerID = &sv.ID
		ss, err := server.gojo.ListAnimeMovieServerSubVideos(ctx, sv.ID)
		if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.Internal, "error when list anime movie server sub videos : %s", err)
		}

		subV := make([]db.AnimeMovieVideo, len(ss))
		for i, ks := range ss {
			subV[i], err = server.gojo.GetAnimeMovieVideo(ctx, ks.VideoID)
			if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
				return nil, status.Errorf(codes.Internal, "error when get anime movie server sub videos : %s", err)
			}
		}

		st, err := server.gojo.ListAnimeMovieServerSubTorrents(ctx, sv.ID)
		if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.Internal, "error when list anime movie server sub torrents : %s", err)
		}

		subT := make([]db.AnimeMovieTorrent, len(st))
		for i, kst := range st {
			subT[i], err = server.gojo.GetAnimeMovieTorrent(ctx, kst.ServerID)
			if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
				return nil, status.Errorf(codes.Internal, "error when get anime movie server sub torrents : %s", err)
			}
		}

		res.Sub = &ampb.AnimeMovieSubDataResponse{
			Videos:   shared.ConvertAnimeMovieVideos(subV),
			Torrents: shared.ConvertAnimeMovieTorrents(subT),
		}

		sd, err := server.gojo.ListAnimeMovieServerDubVideos(ctx, sv.ID)
		if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.Internal, "error when list anime movie server dub videos : %s", err)
		}

		dubV := make([]db.AnimeMovieVideo, len(sd))
		for i, kd := range sd {
			dubV[i], err = server.gojo.GetAnimeMovieVideo(ctx, kd.VideoID)
			if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
				return nil, status.Errorf(codes.Internal, "error when get anime movie server dub videos : %s", err)
			}
		}

		dt, err := server.gojo.ListAnimeMovieServerDubTorrents(ctx, sv.ID)
		if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
			return nil, status.Errorf(codes.Internal, "error when list anime movie server sub torrents : %s", err)
		}

		dubT := make([]db.AnimeMovieTorrent, len(dt))
		for i, kdt := range dt {
			subT[i], err = server.gojo.GetAnimeMovieTorrent(ctx, kdt.ServerID)
			if err != nil && db.ErrorCode(err) != db.ErrRecordNotFound.Error() {
				return nil, status.Errorf(codes.Internal, "error when get anime movie server dub torrents : %s", err)
			}
		}

		res.Dub = &ampb.AnimeMovieDubDataResponse{
			Videos:   shared.ConvertAnimeMovieVideos(dubV),
			Torrents: shared.ConvertAnimeMovieTorrents(dubT),
		}

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
