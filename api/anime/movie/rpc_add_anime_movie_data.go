package animeMovie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) AddAnimeMovieData(ctx context.Context, req *ampb.AddAnimeMovieDataRequest) (*ampb.AddAnimeMovieDataResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime movie data")
	}

	if violations := validateAddAnimeMovieDataRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var DBSV []db.CreateAnimeMovieVideoParams
	if req.Sub != nil {
		DBSV = make([]db.CreateAnimeMovieVideoParams, len(req.GetSub().Videos))
		for i, s := range req.GetSub().Videos {
			DBSV[i].LanguageID = s.LanguageID
			DBSV[i].Autority = s.Autority
			DBSV[i].Referer = s.Referer
			DBSV[i].Link = s.Link
			DBSV[i].Quality = s.Quality
		}
	}

	var DBDV []db.CreateAnimeMovieVideoParams
	if req.Dub != nil {
		DBDV = make([]db.CreateAnimeMovieVideoParams, len(req.GetDub().Videos))
		for i, d := range req.GetDub().Videos {
			DBDV[i].LanguageID = d.LanguageID
			DBDV[i].Autority = d.Autority
			DBDV[i].Referer = d.Referer
			DBDV[i].Link = d.Link
			DBDV[i].Quality = d.Quality
		}
	}

	var DBST []db.CreateAnimeMovieTorrentParams
	if req.Sub != nil {
		DBST = make([]db.CreateAnimeMovieTorrentParams, len(req.GetSub().Torrents))
		for i, s := range req.GetSub().Torrents {
			DBST[i].LanguageID = s.LanguageID
			DBST[i].FileName = s.FileName
			DBST[i].TorrentHash = s.TorrentHash
			DBST[i].TorrentFile = s.TorrentFile
			DBST[i].Seeds = s.Seeds
			DBST[i].Peers = s.Peers
			DBST[i].Leechers = s.Leechers
			DBST[i].SizeBytes = s.SizeBytes
			DBST[i].Quality = s.Quality
		}
	}

	var DBDT []db.CreateAnimeMovieTorrentParams
	if req.Dub != nil {
		DBDT = make([]db.CreateAnimeMovieTorrentParams, len(req.GetDub().Torrents))
		for i, d := range req.GetDub().Torrents {
			DBDT[i].LanguageID = d.LanguageID
			DBDT[i].FileName = d.FileName
			DBDT[i].TorrentHash = d.TorrentHash
			DBDT[i].TorrentFile = d.TorrentFile
			DBDT[i].Seeds = d.Seeds
			DBDT[i].Peers = d.Peers
			DBDT[i].Leechers = d.Leechers
			DBDT[i].SizeBytes = d.SizeBytes
			DBDT[i].Quality = d.Quality
		}
	}

	arg := db.AddAnimeMovieDataTxParams{
		ServerID:    req.GetServerID(),
		SubVideos:   DBSV,
		DubVideos:   DBDV,
		SubTorrents: DBST,
		DubTorrents: DBDT,
	}

	data, err := server.gojo.AddAnimeMovieDataTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to add anime movie videos & torrents to the server", err)
	}

	res := &ampb.AddAnimeMovieDataResponse{
		AnimeMovie: shared.ConvertAnimeMovie(data.AnimeMovie),
		Sub: &ampb.AnimeMovieSubDataResponse{
			Videos:   shared.ConvertAnimeMovieVideos(data.AnimeMovieSubVideos),
			Torrents: shared.ConvertAnimeMovieTorrents(data.AnimeMovieSubTorrents),
		},
		Dub: &ampb.AnimeMovieDubDataResponse{
			Videos:   shared.ConvertAnimeMovieVideos(data.AnimeMovieDubVideos),
			Torrents: shared.ConvertAnimeMovieTorrents(data.AnimeMovieDubTorrents),
		},
	}

	return res, nil
}

func validateAddAnimeMovieDataRequest(req *ampb.AddAnimeMovieDataRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetServerID()); err != nil {
		violations = append(violations, shared.FieldViolation("serverID", err))
	}

	if req.Sub == nil && req.Dub == nil {
		violations = append(violations, shared.FieldViolation("sub,dub", errors.New("add one video or torrent at least")))
	} else {
		if req.Sub != nil {
			if req.GetSub().Videos != nil {
				for _, s := range req.GetSub().GetVideos() {
					if err := utils.ValidateInt(int64(s.GetLanguageID())); err != nil {
						violations = append(violations, shared.FieldViolation("languageID", err))
					}

					if err := utils.ValidateURL(s.GetAutority(), ""); err != nil {
						violations = append(violations, shared.FieldViolation("autority", err))
					}

					if err := utils.ValidateURL(s.GetReferer(), ""); err != nil {
						violations = append(violations, shared.FieldViolation("link", err))
					}

					if err := utils.ValidateURL(s.GetLink(), ""); err != nil {
						violations = append(violations, shared.FieldViolation("link", err))
					}

					if err := utils.ValidateQuality(s.GetQuality()); err != nil {
						violations = append(violations, shared.FieldViolation("quality", err))
					}

				}
			}
			if req.GetSub().Torrents != nil {
				for _, s := range req.GetSub().GetTorrents() {
					if err := utils.ValidateInt(int64(s.GetLanguageID())); err != nil {
						violations = append(violations, shared.FieldViolation("languageID", err))
					}

					if err := utils.ValidateString(s.GetFileName(), 1, 100); err != nil {
						violations = append(violations, shared.FieldViolation("fileName", err))
					}

					if err := utils.ValidateString(s.GetTorrentHash(), 32, 64); err != nil {
						violations = append(violations, shared.FieldViolation("torrentHash", err))
					}

					if err := utils.ValidateString(s.GetTorrentFile(), 0, 200); err != nil {
						violations = append(violations, shared.FieldViolation("torrentFile", err))
					}

					if err := utils.ValidateInt(int64(s.GetSeeds() + 1)); err != nil {
						violations = append(violations, shared.FieldViolation("seeds", err))
					}

					if err := utils.ValidateInt(int64(s.GetPeers() + 1)); err != nil {
						violations = append(violations, shared.FieldViolation("peers", err))
					}

					if err := utils.ValidateInt(int64(s.GetLeechers() + 1)); err != nil {
						violations = append(violations, shared.FieldViolation("leechers", err))
					}

					if err := utils.ValidateInt(s.GetSizeBytes()); err != nil {
						violations = append(violations, shared.FieldViolation("sizeBytes", err))
					}

					if err := utils.ValidateQuality(s.GetQuality()); err != nil {
						violations = append(violations, shared.FieldViolation("quality", err))
					}

				}
			}
		}

		if req.Dub != nil {
			if req.GetDub().Videos != nil {
				for _, d := range req.GetDub().GetVideos() {
					if err := utils.ValidateInt(int64(d.GetLanguageID())); err != nil {
						violations = append(violations, shared.FieldViolation("languageID", err))
					}

					if err := utils.ValidateURL(d.GetAutority(), ""); err != nil {
						violations = append(violations, shared.FieldViolation("autority", err))
					}

					if err := utils.ValidateURL(d.GetReferer(), ""); err != nil {
						violations = append(violations, shared.FieldViolation("link", err))
					}

					if err := utils.ValidateURL(d.GetLink(), ""); err != nil {
						violations = append(violations, shared.FieldViolation("link", err))
					}

					if err := utils.ValidateQuality(d.GetQuality()); err != nil {
						violations = append(violations, shared.FieldViolation("quality", err))
					}

				}
			}
			if req.GetDub().Torrents != nil {
				for _, d := range req.GetDub().GetTorrents() {
					if err := utils.ValidateInt(int64(d.GetLanguageID())); err != nil {
						violations = append(violations, shared.FieldViolation("languageID", err))
					}

					if err := utils.ValidateString(d.GetFileName(), 1, 100); err != nil {
						violations = append(violations, shared.FieldViolation("fileName", err))
					}

					if err := utils.ValidateString(d.GetTorrentHash(), 32, 64); err != nil {
						violations = append(violations, shared.FieldViolation("torrentHash", err))
					}

					if err := utils.ValidateString(d.GetTorrentFile(), 0, 200); err != nil {
						violations = append(violations, shared.FieldViolation("torrentFile", err))
					}

					if err := utils.ValidateInt(int64(d.GetSeeds() + 1)); err != nil {
						violations = append(violations, shared.FieldViolation("seeds", err))
					}

					if err := utils.ValidateInt(int64(d.GetPeers() + 1)); err != nil {
						violations = append(violations, shared.FieldViolation("peers", err))
					}

					if err := utils.ValidateInt(int64(d.GetLeechers() + 1)); err != nil {
						violations = append(violations, shared.FieldViolation("leechers", err))
					}

					if err := utils.ValidateInt(d.GetSizeBytes()); err != nil {
						violations = append(violations, shared.FieldViolation("sizeBytes", err))
					}

					if err := utils.ValidateQuality(d.GetQuality()); err != nil {
						violations = append(violations, shared.FieldViolation("quality", err))
					}

				}
			}
		}
	}

	return violations
}
