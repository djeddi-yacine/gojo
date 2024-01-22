package asapiv1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeEpisodeData(ctx context.Context, req *aspbv1.CreateAnimeEpisodeDataRequest) (*aspbv1.CreateAnimeEpisodeDataResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add anime serie data")
	}

	if violations := validateCreateAnimeEpisodeDataRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeEpisodeDataTxParams{
		ServerID: req.GetServerID(),
	}

	if req.Sub != nil {
		arg.SubVideos = make([]db.CreateAnimeEpisodeVideoParams, len(req.GetSub().GetVideos()))
		for i, v := range req.GetSub().GetVideos() {
			arg.SubVideos[i].LanguageID = v.LanguageID
			arg.SubVideos[i].Authority = v.Authority
			arg.SubVideos[i].Referer = v.Referer
			arg.SubVideos[i].Link = v.Link
			arg.SubVideos[i].Quality = v.Quality
		}

		arg.SubTorrents = make([]db.CreateAnimeEpisodeTorrentParams, len(req.GetSub().GetTorrents()))
		for i, v := range req.GetSub().GetTorrents() {
			arg.SubTorrents[i].LanguageID = v.LanguageID
			arg.SubTorrents[i].FileName = v.FileName
			arg.SubTorrents[i].TorrentHash = v.TorrentHash
			arg.SubTorrents[i].TorrentFile = v.TorrentFile
			arg.SubTorrents[i].Seeds = v.Seeds
			arg.SubTorrents[i].Peers = v.Peers
			arg.SubTorrents[i].Leechers = v.Leechers
			arg.SubTorrents[i].SizeBytes = v.SizeBytes
			arg.SubTorrents[i].Quality = v.Quality
		}
	}

	if req.Dub != nil {
		arg.DubVideos = make([]db.CreateAnimeEpisodeVideoParams, len(req.GetDub().GetVideos()))
		for i, v := range req.GetDub().GetVideos() {
			arg.DubVideos[i].LanguageID = v.LanguageID
			arg.DubVideos[i].Authority = v.Authority
			arg.DubVideos[i].Referer = v.Referer
			arg.DubVideos[i].Link = v.Link
			arg.DubVideos[i].Quality = v.Quality
		}

		arg.DubTorrents = make([]db.CreateAnimeEpisodeTorrentParams, len(req.GetDub().GetTorrents()))
		for i, v := range req.GetDub().GetTorrents() {
			arg.DubTorrents[i].LanguageID = v.LanguageID
			arg.DubTorrents[i].FileName = v.FileName
			arg.DubTorrents[i].TorrentHash = v.TorrentHash
			arg.DubTorrents[i].TorrentFile = v.TorrentFile
			arg.DubTorrents[i].Seeds = v.Seeds
			arg.DubTorrents[i].Peers = v.Peers
			arg.DubTorrents[i].Leechers = v.Leechers
			arg.DubTorrents[i].SizeBytes = v.SizeBytes
			arg.DubTorrents[i].Quality = v.Quality
		}
	}

	data, err := server.gojo.CreateAnimeEpisodeDataTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to add anime serie videos & torrents to the server", err)
	}

	res := &aspbv1.CreateAnimeEpisodeDataResponse{
		Episode: convertAnimeEpisode(data.AnimeEpisode),
		Sub: &ashpbv1.AnimeSubDataResponse{
			Videos:   convertAnimeEpisodeVideos(data.AnimeEpisodeSubVideos),
			Torrents: convertAnimeEpisodeTorrents(data.AnimeEpisodeSubTorrents),
		},
		Dub: &ashpbv1.AnimeDubDataResponse{
			Videos:   convertAnimeEpisodeVideos(data.AnimeEpisodeDubVideos),
			Torrents: convertAnimeEpisodeTorrents(data.AnimeEpisodeDubTorrents),
		},
	}

	return res, nil
}

func validateCreateAnimeEpisodeDataRequest(req *aspbv1.CreateAnimeEpisodeDataRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetServerID()); err != nil {
		violations = append(violations, shv1.FieldViolation("serverID", err))
	}

	if req.Sub == nil && req.Dub == nil {
		violations = append(violations, shv1.FieldViolation("sub,dub", errors.New("add one video or torrent at least")))
	} else {
		if req.Sub != nil {
			if req.GetSub().Videos != nil {
				for _, v := range req.GetSub().GetVideos() {
					if err := utils.ValidateInt(int64(v.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateURL(v.GetAuthority(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("authority", err))
					}

					if err := utils.ValidateURL(v.GetReferer(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateURL(v.GetLink(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateQuality(v.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
			if req.GetSub().Torrents != nil {
				for _, v := range req.GetSub().GetTorrents() {
					if err := utils.ValidateInt(int64(v.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateString(v.GetFileName(), 1, 100); err != nil {
						violations = append(violations, shv1.FieldViolation("fileName", err))
					}

					if err := utils.ValidateString(v.GetTorrentHash(), 32, 64); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentHash", err))
					}

					if err := utils.ValidateString(v.GetTorrentFile(), 0, 200); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentFile", err))
					}

					if err := utils.ValidateInt(int64(v.GetSeeds() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("seeds", err))
					}

					if err := utils.ValidateInt(int64(v.GetPeers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("peers", err))
					}

					if err := utils.ValidateInt(int64(v.GetLeechers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("leechers", err))
					}

					if err := utils.ValidateInt(v.GetSizeBytes()); err != nil {
						violations = append(violations, shv1.FieldViolation("sizeBytes", err))
					}

					if err := utils.ValidateQuality(v.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
		}

		if req.Dub != nil {
			if req.GetDub().Videos != nil {
				for _, v := range req.GetDub().GetVideos() {
					if err := utils.ValidateInt(int64(v.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateURL(v.GetAuthority(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("authority", err))
					}

					if err := utils.ValidateURL(v.GetReferer(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateURL(v.GetLink(), ""); err != nil {
						violations = append(violations, shv1.FieldViolation("link", err))
					}

					if err := utils.ValidateQuality(v.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
			if req.GetDub().Torrents != nil {
				for _, v := range req.GetDub().GetTorrents() {
					if err := utils.ValidateInt(int64(v.GetLanguageID())); err != nil {
						violations = append(violations, shv1.FieldViolation("languageID", err))
					}

					if err := utils.ValidateString(v.GetFileName(), 1, 100); err != nil {
						violations = append(violations, shv1.FieldViolation("fileName", err))
					}

					if err := utils.ValidateString(v.GetTorrentHash(), 32, 64); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentHash", err))
					}

					if err := utils.ValidateString(v.GetTorrentFile(), 0, 200); err != nil {
						violations = append(violations, shv1.FieldViolation("torrentFile", err))
					}

					if err := utils.ValidateInt(int64(v.GetSeeds() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("seeds", err))
					}

					if err := utils.ValidateInt(int64(v.GetPeers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("peers", err))
					}

					if err := utils.ValidateInt(int64(v.GetLeechers() + 1)); err != nil {
						violations = append(violations, shv1.FieldViolation("leechers", err))
					}

					if err := utils.ValidateInt(v.GetSizeBytes()); err != nil {
						violations = append(violations, shv1.FieldViolation("sizeBytes", err))
					}

					if err := utils.ValidateQuality(v.GetQuality()); err != nil {
						violations = append(violations, shv1.FieldViolation("quality", err))
					}

				}
			}
		}
	}

	return violations
}
