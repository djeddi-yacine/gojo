package db

import (
	"context"
)

type CreateAnimeEpisodeDataTxParams struct {
	ServerID    int64
	SubVideos   []CreateAnimeEpisodeVideoParams
	DubVideos   []CreateAnimeEpisodeVideoParams
	SubTorrents []CreateAnimeEpisodeTorrentParams
	DubTorrents []CreateAnimeEpisodeTorrentParams
}

type CreateAnimeEpisodeDataTxResult struct {
	AnimeEpisode            AnimeEpisode
	AnimeEpisodeSubVideos   []AnimeEpisodeVideo
	AnimeEpisodeDubVideos   []AnimeEpisodeVideo
	AnimeEpisodeSubTorrents []AnimeEpisodeTorrent
	AnimeEpisodeDubTorrents []AnimeEpisodeTorrent
}

func (gojo *SQLGojo) CreateAnimeEpisodeDataTx(ctx context.Context, arg CreateAnimeEpisodeDataTxParams) (CreateAnimeEpisodeDataTxResult, error) {
	var result CreateAnimeEpisodeDataTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		server, err := q.GetAnimeEpisodeServer(ctx, arg.ServerID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeEpisode, err = q.GetAnimeEpisodeByEpisodeID(ctx, server.EpisodeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.SubVideos != nil {
			var videoArg CreateAnimeEpisodeVideoParams
			subVideos := make([]CreateAnimeEpisodeServerSubVideoParams, len(arg.SubVideos))
			result.AnimeEpisodeSubVideos = make([]AnimeEpisodeVideo, len(arg.SubVideos))

			for i, s := range arg.SubVideos {
				videoArg = CreateAnimeEpisodeVideoParams{
					LanguageID: s.LanguageID,
					Authority:  s.Authority,
					Referer:    s.Referer,
					Link:       s.Link,
					Quality:    s.Quality,
				}

				v, err := q.CreateAnimeEpisodeVideo(ctx, videoArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeEpisodeSubVideos[i] = v
				subVideos[i].VideoID = v.ID
				subVideos[i].ServerID = server.ID
			}

			for _, esv := range subVideos {
				_, err = q.CreateAnimeEpisodeServerSubVideo(ctx, esv)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.DubVideos != nil {
			var videoArg CreateAnimeEpisodeVideoParams
			dubVideos := make([]CreateAnimeEpisodeServerDubVideoParams, len(arg.DubVideos))
			result.AnimeEpisodeDubVideos = make([]AnimeEpisodeVideo, len(arg.DubVideos))

			for i, d := range arg.DubVideos {
				videoArg = CreateAnimeEpisodeVideoParams{
					LanguageID: d.LanguageID,
					Authority:  d.Authority,
					Referer:    d.Referer,
					Link:       d.Link,
					Quality:    d.Quality,
				}

				v, err := q.CreateAnimeEpisodeVideo(ctx, videoArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeEpisodeSubVideos[i] = v
				dubVideos[i].VideoID = v.ID
				dubVideos[i].ServerID = server.ID

			}

			for _, edv := range dubVideos {
				_, err = q.CreateAnimeEpisodeServerDubVideo(ctx, edv)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.SubTorrents != nil {
			var torrentArg CreateAnimeEpisodeTorrentParams
			subTorrents := make([]CreateAnimeEpisodeServerSubTorrentParams, len(arg.SubTorrents))
			result.AnimeEpisodeSubTorrents = make([]AnimeEpisodeTorrent, len(arg.SubTorrents))

			for i, s := range arg.SubTorrents {
				torrentArg = CreateAnimeEpisodeTorrentParams{
					LanguageID:  s.LanguageID,
					FileName:    s.FileName,
					TorrentHash: s.TorrentHash,
					TorrentFile: s.TorrentFile,
					Seeds:       s.Seeds,
					Peers:       s.Peers,
					Leechers:    s.Leechers,
					SizeBytes:   s.SizeBytes,
					Quality:     s.Quality,
				}

				t, err := q.CreateAnimeEpisodeTorrent(ctx, torrentArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeEpisodeSubTorrents[i] = t
				subTorrents[i].TorrentID = t.ID
				subTorrents[i].ServerID = server.ID

			}

			for _, est := range subTorrents {
				_, err = q.CreateAnimeEpisodeServerSubTorrent(ctx, est)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.DubTorrents != nil {
			var torrentArg CreateAnimeEpisodeTorrentParams
			dubTorrents := make([]CreateAnimeEpisodeServerDubTorrentParams, len(arg.DubTorrents))
			result.AnimeEpisodeDubTorrents = make([]AnimeEpisodeTorrent, len(arg.DubTorrents))

			for i, d := range arg.DubTorrents {
				torrentArg = CreateAnimeEpisodeTorrentParams{
					LanguageID:  d.LanguageID,
					FileName:    d.FileName,
					TorrentHash: d.TorrentHash,
					TorrentFile: d.TorrentFile,
					Seeds:       d.Seeds,
					Peers:       d.Peers,
					Leechers:    d.Leechers,
					SizeBytes:   d.SizeBytes,
					Quality:     d.Quality,
				}

				t, err := q.CreateAnimeEpisodeTorrent(ctx, torrentArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeEpisodeDubTorrents[i] = t
				dubTorrents[i].TorrentID = t.ID
				dubTorrents[i].ServerID = server.ID

			}

			for _, edt := range dubTorrents {
				_, err = q.CreateAnimeEpisodeServerDubTorrent(ctx, edt)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		return err
	})

	return result, err
}
