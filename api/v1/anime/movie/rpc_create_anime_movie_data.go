package amapiv1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieData(ctx context.Context, req *ampbv1.CreateAnimeMovieDataRequest) (*ampbv1.CreateAnimeMovieDataResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime movie data")
	}

	if violations := validateCreateAnimeMovieDataRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	var DBSV []db.CreateAnimeMovieVideoParams
	if req.Sub != nil {
		DBSV = make([]db.CreateAnimeMovieVideoParams, len(req.GetSub().Videos))
		for i, s := range req.GetSub().Videos {
			DBSV[i].LanguageID = s.LanguageID
			DBSV[i].Authority = s.Authority
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
			DBDV[i].Authority = d.Authority
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

	arg := db.CreateAnimeMovieDataTxParams{
		ServerID:    req.GetServerID(),
		SubVideos:   DBSV,
		DubVideos:   DBDV,
		SubTorrents: DBST,
		DubTorrents: DBDT,
	}

	data, err := server.gojo.CreateAnimeMovieDataTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to add anime movie videos & torrents to the server", err)
	}

	res := &ampbv1.CreateAnimeMovieDataResponse{
		AnimeMovie: convertAnimeMovie(data.AnimeMovie),
		Sub: &ampbv1.AnimeMovieSubDataResponse{
			Videos:   convertAnimeMovieVideos(data.AnimeMovieSubVideos),
			Torrents: convertAnimeMovieTorrents(data.AnimeMovieSubTorrents),
		},
		Dub: &ampbv1.AnimeMovieDubDataResponse{
			Videos:   convertAnimeMovieVideos(data.AnimeMovieDubVideos),
			Torrents: convertAnimeMovieTorrents(data.AnimeMovieDubTorrents),
		},
	}

	return res, nil
}

func validateCreateAnimeMovieDataRequest(req *ampbv1.CreateAnimeMovieDataRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetServerID()); err != nil {
		violations = append(violations, shv1.FieldViolation("serverID", err))
	}

	if req.Sub == nil && req.Dub == nil {
		violations = append(violations, shv1.FieldViolation("sub,dub", errors.New("add one video or torrent at least")))
	} else {
		if req.Sub != nil {
			if req.GetSub().Videos != nil {
				for _, s := range req.GetSub().GetVideos() {
					if err := utils.ValidateInt(int64(s.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateURL(s.GetAuthority(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("authority", err))
					}

					if err := utils.ValidateURL(s.GetReferer(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateURL(s.GetLink(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateQuality(s.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
			if req.GetSub().Torrents != nil {
				for _, s := range req.GetSub().GetTorrents() {
					if err := utils.ValidateInt(int64(s.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateString(s.GetFileName(), 1, 100); err != nil {
						violations = append(violations, shv1.FieldViolation("fileName", err))
					}

					if err := utils.ValidateString(s.GetTorrentHash(), 32, 64); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentHash", err))
					}

					if err := utils.ValidateString(s.GetTorrentFile(), 0, 200); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentFile", err))
					}

					if err := utils.ValidateInt(int64(s.GetSeeds() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("seeds", err))
					}

					if err := utils.ValidateInt(int64(s.GetPeers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("peers", err))
					}

					if err := utils.ValidateInt(int64(s.GetLeechers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("leechers", err))
					}

					if err := utils.ValidateInt(s.GetSizeBytes()); err != nil {
						violations = append(violations, shv1.FieldViolation("sizeBytes", err))
					}

					if err := utils.ValidateQuality(s.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
		}

		if req.Dub != nil {
			if req.GetDub().Videos != nil {
				for _, d := range req.GetDub().GetVideos() {
					if err := utils.ValidateInt(int64(d.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateURL(d.GetAuthority(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("authority", err))
					}

					if err := utils.ValidateURL(d.GetReferer(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateURL(d.GetLink(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateQuality(d.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
			if req.GetDub().Torrents != nil {
				for _, d := range req.GetDub().GetTorrents() {
					if err := utils.ValidateInt(int64(d.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateString(d.GetFileName(), 1, 100); err != nil {
						violations = append(violations, shv1.FieldViolation("fileName", err))
					}

					if err := utils.ValidateString(d.GetTorrentHash(), 32, 64); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentHash", err))
					}

					if err := utils.ValidateString(d.GetTorrentFile(), 0, 200); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentFile", err))
					}

					if err := utils.ValidateInt(int64(d.GetSeeds() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("seeds", err))
					}

					if err := utils.ValidateInt(int64(d.GetPeers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("peers", err))
					}

					if err := utils.ValidateInt(int64(d.GetLeechers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("leechers", err))
					}

					if err := utils.ValidateInt(d.GetSizeBytes()); err != nil {
						violations = append(violations, shv1.FieldViolation("sizeBytes", err))
					}

					if err := utils.ValidateQuality(d.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
		}
	}

	return violations
}
