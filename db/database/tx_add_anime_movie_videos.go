package db

import (
	"context"
	"fmt"
	"math"
)

type AddAnimeMovieDataTxParams struct {
	ServerID    int64
	SubVideos   []CreateAnimeMovieVideoParams
	DubVideos   []CreateAnimeMovieVideoParams
	SubTorrents []CreateAnimeMovieTorrentParams
	DubTorrents []CreateAnimeMovieTorrentParams
}

type AddAnimeMovieDataTxResult struct {
	AnimeMovie            AnimeMovie
	AnimeMovieSubVideos   []AnimeMovieVideo
	AnimeMovieDubVideos   []AnimeMovieVideo
	AnimeMovieSubTorrents []AnimeMovieTorrent
	AnimeMovieDubTorrents []AnimeMovieTorrent
}

func checkLanguage(languages []Language, n int32) bool {
	for _, language := range languages {
		if language.ID == n {
			return true
		}
	}

	return false
}

func (gojo *SQLGojo) AddAnimeMovieDataTx(ctx context.Context, arg AddAnimeMovieDataTxParams) (AddAnimeMovieDataTxResult, error) {
	var result AddAnimeMovieDataTxResult

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

		languages, err := q.ListLanguages(ctx, ListLanguagesParams{
			Limit:  math.MaxInt32,
			Offset: 0,
		})
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.SubVideos != nil {
			var videoArg CreateAnimeMovieVideoParams
			subVideos := make([]CreateAnimeMovieServerSubVideoParams, len(arg.SubVideos))
			result.AnimeMovieSubVideos = make([]AnimeMovieVideo, len(arg.SubVideos))

			for i, s := range arg.SubVideos {
				if checkLanguage(languages, s.LanguageID) {
					videoArg = CreateAnimeMovieVideoParams{
						LanguageID: s.LanguageID,
						Autority:   s.Autority,
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
				} else {
					return fmt.Errorf("sub > videos: there is no language with ID : %d", s.LanguageID)
				}
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
				if checkLanguage(languages, d.LanguageID) {
					videoArg = CreateAnimeMovieVideoParams{
						LanguageID: d.LanguageID,
						Autority:   d.Autority,
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
				} else {
					return fmt.Errorf("dub > videos: there is no language with ID : %d", d.LanguageID)
				}
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
			subTorrents := make([]CreateAnimeMovieServerTorrentParams, len(arg.SubTorrents))
			result.AnimeMovieSubTorrents = make([]AnimeMovieTorrent, len(arg.SubTorrents))

			for i, s := range arg.SubTorrents {
				if checkLanguage(languages, s.LanguageID) {
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
				} else {
					return fmt.Errorf("sub > torrent: there is no language with ID : %d", s.LanguageID)
				}
			}

			for _, amst := range subTorrents {
				_, err = q.CreateAnimeMovieServerTorrent(ctx, amst)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.DubTorrents != nil {
			var torrentArg CreateAnimeMovieTorrentParams
			dubTorrents := make([]CreateAnimeMovieServerTorrentParams, len(arg.DubTorrents))
			result.AnimeMovieDubTorrents = make([]AnimeMovieTorrent, len(arg.DubTorrents))

			for i, d := range arg.DubTorrents {
				if checkLanguage(languages, d.LanguageID) {
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
				} else {
					return fmt.Errorf("dub > torrent: there is no language with ID : %d", d.LanguageID)
				}
			}

			for _, amdt := range dubTorrents {
				_, err = q.CreateAnimeMovieServerTorrent(ctx, amdt)
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
