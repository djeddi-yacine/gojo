package amapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAnimeMovie(v db.AnimeMovie) *ampbv1.AnimeMovieResponse {
	if v.ID != 0 {
		return &ampbv1.AnimeMovieResponse{
			ID:                v.ID,
			OriginalTitle:     v.OriginalTitle,
			UniqueID:          v.UniqueID.String(),
			Aired:             timestamppb.New(v.Aired),
			ReleaseYear:       v.ReleaseYear,
			Rating:            v.Rating,
			Duration:          durationpb.New(v.Duration),
			PortraitPoster:    v.PortraitPoster,
			PortraitBlurHash:  v.PortraitBlurHash,
			LandscapePoster:   v.LandscapePoster,
			LandscapeBlurHash: v.LandscapeBlurHash,
			CreatedAt:         timestamppb.New(v.CreatedAt),
		}
	}

	return nil
}

func convertAnimeMovieVideos(v []db.AnimeMovieVideo) []*ampbv1.AnimeMovieVideoResponse {
	if len(v) > 0 {
		videos := make([]*ampbv1.AnimeMovieVideoResponse, len(v))

		for i, x := range v {
			videos[i] = &ampbv1.AnimeMovieVideoResponse{
				ID:         x.ID,
				LanguageID: x.LanguageID,
				Authority:  x.Authority,
				Referer:    x.Referer,
				Link:       x.Link,
				Quality:    x.Quality,
				CreatedAt:  timestamppb.New(x.CreatedAt),
			}
		}
		return videos
	} else {
		return nil
	}
}

func convertAnimeMovieTorrents(v []db.AnimeMovieTorrent) []*ampbv1.AnimeMovieTorrentResponse {
	if len(v) > 0 {
		torrents := make([]*ampbv1.AnimeMovieTorrentResponse, len(v))

		for i, x := range v {
			torrents[i] = &ampbv1.AnimeMovieTorrentResponse{
				ID:          x.ID,
				LanguageID:  x.LanguageID,
				FileName:    x.FileName,
				TorrentHash: x.TorrentHash,
				TorrentFile: x.TorrentFile,
				Seeds:       x.Seeds,
				Peers:       x.Peers,
				Leechers:    x.Leechers,
				SizeBytes:   x.SizeBytes,
				Quality:     x.Quality,
				CreatedAt:   timestamppb.New(x.CreatedAt),
			}
		}

		return torrents
	} else {
		return nil
	}
}
