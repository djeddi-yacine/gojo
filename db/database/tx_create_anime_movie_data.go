package db

import (
	"context"
)

type CreateAnimeMovieDataTxParams struct {
	ServerID    int64
	SubVideos   []CreateAnimeMovieVideoParams
	DubVideos   []CreateAnimeMovieVideoParams
	SubTorrents []CreateAnimeMovieTorrentParams
	DubTorrents []CreateAnimeMovieTorrentParams
}

type CreateAnimeMovieDataTxResult struct {
	AnimeMovie            AnimeMovie
	AnimeMovieSubVideos   []AnimeMovieVideo
	AnimeMovieDubVideos   []AnimeMovieVideo
	AnimeMovieSubTorrents []AnimeMovieTorrent
	AnimeMovieDubTorrents []AnimeMovieTorrent
}

func (gojo *SQLGojo) CreateAnimeMovieDataTx(ctx context.Context, arg CreateAnimeMovieDataTxParams) (CreateAnimeMovieDataTxResult, error) {
	var result CreateAnimeMovieDataTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		server, err := q.GetAnimeMovieServer(ctx, arg.ServerID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeMovie, err = q.GetAnimeMovie(ctx, server.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.SubVideos != nil {
			var videoArg CreateAnimeMovieVideoParams
			subVideos := make([]CreateAnimeMovieServerSubVideoParams, len(arg.SubVideos))
			result.AnimeMovieSubVideos = make([]AnimeMovieVideo, len(arg.SubVideos))

			for i, s := range arg.SubVideos {
				videoArg = CreateAnimeMovieVideoParams{
					LanguageID: s.LanguageID,
					Authority:  s.Authority,
					Referer:    s.Referer,
					Link:       s.Link,
					Quality:    s.Quality,
				}

				v, err := q.CreateAnimeMovieVideo(ctx, videoArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMovieSubVideos[i] = v
				subVideos[i].VideoID = v.ID
				subVideos[i].ServerID = server.ID
			}

			for _, amsv := range subVideos {
				_, err = q.CreateAnimeMovieServerSubVideo(ctx, amsv)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.DubVideos != nil {
			var videoArg CreateAnimeMovieVideoParams
			dubVideos := make([]CreateAnimeMovieServerDubVideoParams, len(arg.DubVideos))
			result.AnimeMovieDubVideos = make([]AnimeMovieVideo, len(arg.DubVideos))

			for i, d := range arg.DubVideos {
				videoArg = CreateAnimeMovieVideoParams{
					LanguageID: d.LanguageID,
					Authority:  d.Authority,
					Referer:    d.Referer,
					Link:       d.Link,
					Quality:    d.Quality,
				}

				v, err := q.CreateAnimeMovieVideo(ctx, videoArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMovieSubVideos[i] = v
				dubVideos[i].VideoID = v.ID
				dubVideos[i].ServerID = server.ID

			}

			for _, amdv := range dubVideos {
				_, err = q.CreateAnimeMovieServerDubVideo(ctx, amdv)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.SubTorrents != nil {
			var torrentArg CreateAnimeMovieTorrentParams
			subTorrents := make([]CreateAnimeMovieServerSubTorrentParams, len(arg.SubTorrents))
			result.AnimeMovieSubTorrents = make([]AnimeMovieTorrent, len(arg.SubTorrents))

			for i, s := range arg.SubTorrents {
				torrentArg = CreateAnimeMovieTorrentParams{
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

				t, err := q.CreateAnimeMovieTorrent(ctx, torrentArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMovieSubTorrents[i] = t
				subTorrents[i].TorrentID = t.ID
				subTorrents[i].ServerID = server.ID

			}

			for _, amst := range subTorrents {
				_, err = q.CreateAnimeMovieServerSubTorrent(ctx, amst)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.DubTorrents != nil {
			var torrentArg CreateAnimeMovieTorrentParams
			dubTorrents := make([]CreateAnimeMovieServerDubTorrentParams, len(arg.DubTorrents))
			result.AnimeMovieDubTorrents = make([]AnimeMovieTorrent, len(arg.DubTorrents))

			for i, d := range arg.DubTorrents {
				torrentArg = CreateAnimeMovieTorrentParams{
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

				t, err := q.CreateAnimeMovieTorrent(ctx, torrentArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMovieDubTorrents[i] = t
				dubTorrents[i].TorrentID = t.ID
				dubTorrents[i].ServerID = server.ID

			}

			for _, amdt := range dubTorrents {
				_, err = q.CreateAnimeMovieServerDubTorrent(ctx, amdt)
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
