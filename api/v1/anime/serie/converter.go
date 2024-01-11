package asapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAnimeSerie(v db.AnimeSerie) *aspbv1.AnimeSerieResponse {
	if v.ID != 0 {
		return &aspbv1.AnimeSerieResponse{
			ID:                v.ID,
			OriginalTitle:     v.OriginalTitle,
			UniqueID:          v.UniqueID.String(),
			PortraitPoster:    v.PortraitPoster,
			PortraitBlurHash:  v.PortraitBlurHash,
			LandscapePoster:   v.LandscapePoster,
			LandscapeBlurHash: v.LandscapeBlurHash,
			FirstYear:         v.FirstYear,
			LastYear:          v.LastYear,
			MalID:             v.MalID,
			TvdbID:            v.TvdbID,
			TmdbID:            v.TmdbID,
			CreatedAt:         timestamppb.New(v.CreatedAt),
		}
	}

	return nil
}

func convertAnimeSeason(v db.AnimeSerieSeason) *aspbv1.AnimeSeasonResponse {
	if v.ID != 0 {
		return &aspbv1.AnimeSeasonResponse{
			ID:                  v.ID,
			AnimeID:             v.AnimeID,
			SeasonOriginalTitle: v.SeasonOriginalTitle,
			ReleaseYear:         v.ReleaseYear,
			Aired:               timestamppb.New(v.Aired),
			Rating:              v.Rating,
			PortraitPoster:      v.PortraitPoster,
			PortraitBlurHash:    v.PortraitBlurHash,
			CreatedAt:           timestamppb.New(v.CreatedAt),
		}
	}

	return nil
}

func convertAnimeEpisode(v db.AnimeSerieEpisode) *aspbv1.AnimeEpisodeResponse {
	if v.ID != 0 {
		return &aspbv1.AnimeEpisodeResponse{
			ID:                   v.ID,
			SeasonID:             v.SeasonID,
			EpisodeNumber:        uint32(v.EpisodeNumber),
			EpisodeOriginalTitle: v.EpisodeOriginalTitle,
			Aired:                timestamppb.New(v.Aired),
			Rating:               v.Rating,
			Duration:             durationpb.New(v.Duration),
			Thumbnails:           v.Thumbnails,
			ThumbnailsBlurHash:   v.ThumbnailsBlurHash,
			CreatedAt:            timestamppb.New(v.CreatedAt),
		}
	}

	return nil
}

func convertAnimeEpisodeVideos(v []db.AnimeEpisodeVideo) []*aspbv1.AnimeEpisodeVideoResponse {
	if len(v) > 0 {
		videos := make([]*aspbv1.AnimeEpisodeVideoResponse, len(v))

		for i, x := range v {
			videos[i] = &aspbv1.AnimeEpisodeVideoResponse{
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

func convertAnimeEpisodeTorrents(v []db.AnimeEpisodeTorrent) []*aspbv1.AnimeEpisodeTorrentResponse {
	if len(v) > 0 {
		torrents := make([]*aspbv1.AnimeEpisodeTorrentResponse, len(v))

		for i, x := range v {
			torrents[i] = &aspbv1.AnimeEpisodeTorrentResponse{
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
