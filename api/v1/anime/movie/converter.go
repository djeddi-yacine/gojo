package amapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAnimeMovie(a db.AnimeMovie) *ampbv1.AnimeMovieResponse {
	return &ampbv1.AnimeMovieResponse{
		ID:                a.ID,
		OriginalTitle:     a.OriginalTitle,
		Aired:             timestamppb.New(a.Aired),
		ReleaseYear:       a.ReleaseYear,
		Rating:            a.Rating,
		Duration:          durationpb.New(a.Duration),
		PortraitPoster:    a.PortraitPoster,
		PortraitBlurHash:  a.PortraitBlurHash,
		LandscapePoster:   a.LandscapePoster,
		LandscapeBlurHash: a.LandscapeBlurHash,
		CreatedAt:         timestamppb.New(a.CreatedAt),
	}
}

func convertAnimeMovieVideos(amv []db.AnimeMovieVideo) []*ampbv1.AnimeMovieVideoResponse {
	if len(amv) > 0 {
		Videos := make([]*ampbv1.AnimeMovieVideoResponse, len(amv))

		for i, v := range amv {
			Videos[i] = &ampbv1.AnimeMovieVideoResponse{
				ID:         v.ID,
				LanguageID: v.LanguageID,
				Authority:  v.Authority,
				Referer:    v.Referer,
				Link:       v.Link,
				Quality:    v.Quality,
				CreatedAt:  timestamppb.New(v.CreatedAt),
			}
		}
		return Videos
	} else {
		return nil
	}
}

func convertAnimeMovieTorrents(amt []db.AnimeMovieTorrent) []*ampbv1.AnimeMovieTorrentResponse {
	if len(amt) > 0 {
		Torrents := make([]*ampbv1.AnimeMovieTorrentResponse, len(amt))

		for i, t := range amt {
			Torrents[i] = &ampbv1.AnimeMovieTorrentResponse{
				ID:          t.ID,
				LanguageID:  t.LanguageID,
				FileName:    t.FileName,
				TorrentHash: t.TorrentHash,
				TorrentFile: t.TorrentFile,
				Seeds:       t.Seeds,
				Peers:       t.Peers,
				Leechers:    t.Leechers,
				SizeBytes:   t.SizeBytes,
				Quality:     t.Quality,
				CreatedAt:   timestamppb.New(t.CreatedAt),
			}
		}

		return Torrents
	} else {
		return nil
	}
}
