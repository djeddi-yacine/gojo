package db

import (
	"context"
)

type AddAnimeSerieDataTxParams struct {
	ServerID    int64
	SubVideos   []CreateAnimeSerieVideoParams
	DubVideos   []CreateAnimeSerieVideoParams
	SubTorrents []CreateAnimeSerieTorrentParams
	DubTorrents []CreateAnimeSerieTorrentParams
}

type AddAnimeSerieDataTxResult struct {
	Episode            AnimeSerieEpisode
	AnimeSerieSubVideos   []AnimeSerieVideo
	AnimeSerieDubVideos   []AnimeSerieVideo
	AnimeSerieSubTorrents []AnimeSerieTorrent
	AnimeSerieDubTorrents []AnimeSerieTorrent
}

func (gojo *SQLGojo) AddAnimeSerieDataTx(ctx context.Context, arg AddAnimeSerieDataTxParams) (AddAnimeSerieDataTxResult, error) {
	var result AddAnimeSerieDataTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		server, err := q.GetAnimeSerieServer(ctx, arg.ServerID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.Episode, err = q.GetAnimeSerieEpisodeByEpisodeID(ctx, server.EpisodeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.SubVideos != nil {
			var videoArg CreateAnimeSerieVideoParams
			subVideos := make([]CreateAnimeSerieServerSubVideoParams, len(arg.SubVideos))
			result.AnimeSerieSubVideos = make([]AnimeSerieVideo, len(arg.SubVideos))

			for i, s := range arg.SubVideos {
				videoArg = CreateAnimeSerieVideoParams{
					LanguageID: s.LanguageID,
					Autority:   s.Autority,
					Referer:    s.Referer,
					Link:       s.Link,
					Quality:    s.Quality,
				}

				v, err := q.CreateAnimeSerieVideo(ctx, videoArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieSubVideos[i] = v
				subVideos[i].VideoID = v.ID
				subVideos[i].ServerID = server.ID
			}

			for _, esv := range subVideos {
				_, err = q.CreateAnimeSerieServerSubVideo(ctx, esv)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.DubVideos != nil {
			var videoArg CreateAnimeSerieVideoParams
			dubVideos := make([]CreateAnimeSerieServerDubVideoParams, len(arg.DubVideos))
			result.AnimeSerieDubVideos = make([]AnimeSerieVideo, len(arg.DubVideos))

			for i, d := range arg.DubVideos {
				videoArg = CreateAnimeSerieVideoParams{
					LanguageID: d.LanguageID,
					Autority:   d.Autority,
					Referer:    d.Referer,
					Link:       d.Link,
					Quality:    d.Quality,
				}

				v, err := q.CreateAnimeSerieVideo(ctx, videoArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieSubVideos[i] = v
				dubVideos[i].VideoID = v.ID
				dubVideos[i].ServerID = server.ID

			}

			for _, edv := range dubVideos {
				_, err = q.CreateAnimeSerieServerDubVideo(ctx, edv)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.SubTorrents != nil {
			var torrentArg CreateAnimeSerieTorrentParams
			subTorrents := make([]CreateAnimeSerieServerSubTorrentParams, len(arg.SubTorrents))
			result.AnimeSerieSubTorrents = make([]AnimeSerieTorrent, len(arg.SubTorrents))

			for i, s := range arg.SubTorrents {
				torrentArg = CreateAnimeSerieTorrentParams{
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

				t, err := q.CreateAnimeSerieTorrent(ctx, torrentArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieSubTorrents[i] = t
				subTorrents[i].TorrentID = t.ID
				subTorrents[i].ServerID = server.ID

			}

			for _, est := range subTorrents {
				_, err = q.CreateAnimeSerieServerSubTorrent(ctx, est)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.DubTorrents != nil {
			var torrentArg CreateAnimeSerieTorrentParams
			dubTorrents := make([]CreateAnimeSerieServerDubTorrentParams, len(arg.DubTorrents))
			result.AnimeSerieDubTorrents = make([]AnimeSerieTorrent, len(arg.DubTorrents))

			for i, d := range arg.DubTorrents {
				torrentArg = CreateAnimeSerieTorrentParams{
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

				t, err := q.CreateAnimeSerieTorrent(ctx, torrentArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieDubTorrents[i] = t
				dubTorrents[i].TorrentID = t.ID
				dubTorrents[i].ServerID = server.ID

			}

			for _, edt := range dubTorrents {
				_, err = q.CreateAnimeSerieServerDubTorrent(ctx, edt)
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
