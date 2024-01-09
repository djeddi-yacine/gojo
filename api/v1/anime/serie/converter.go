package asapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAnimeSerie(a db.AnimeSerie) *aspbv1.AnimeSerieResponse {
	return &aspbv1.AnimeSerieResponse{
		ID:                a.ID,
		OriginalTitle:     a.OriginalTitle,
		UniqueID:          a.UniqueID.String(),
		PortraitPoster:    a.PortraitPoster,
		PortraitBlurHash:  a.PortraitBlurHash,
		LandscapePoster:   a.LandscapePoster,
		LandscapeBlurHash: a.LandscapeBlurHash,
		FirstYear:         a.FirstYear,
		LastYear:          a.LastYear,
		MalID:             a.MalID,
		TvdbID:            a.TvdbID,
		TmdbID:            a.TmdbID,
		CreatedAt:         timestamppb.New(a.CreatedAt),
	}
}

func convertAnimeSeason(s db.AnimeSerieSeason) *aspbv1.AnimeSeasonResponse {
	return &aspbv1.AnimeSeasonResponse{
		ID:                  s.ID,
		AnimeID:             s.AnimeID,
		SeasonOriginalTitle: s.SeasonOriginalTitle,
		ReleaseYear:         s.ReleaseYear,
		Aired:               timestamppb.New(s.Aired),
		Rating:              s.Rating,
		PortraitPoster:      s.PortraitPoster,
		PortraitBlurHash:    s.PortraitBlurHash,
		CreatedAt:           timestamppb.New(s.CreatedAt),
	}
}

func convertAnimeEpisode(e db.AnimeSerieEpisode) *aspbv1.AnimeEpisodeResponse {
	return &aspbv1.AnimeEpisodeResponse{
		ID:                   e.ID,
		SeasonID:             e.SeasonID,
		EpisodeNumber:        uint32(e.EpisodeNumber),
		EpisodeOriginalTitle: e.EpisodeOriginalTitle,
		Aired:                timestamppb.New(e.Aired),
		Rating:               e.Rating,
		Duration:             durationpb.New(e.Duration),
		Thumbnails:           e.Thumbnails,
		ThumbnailsBlurHash:   e.ThumbnailsBlurHash,
		CreatedAt:            timestamppb.New(e.CreatedAt),
	}
}

func convertAnimeEpisodeVideos(asv []db.AnimeEpisodeVideo) []*aspbv1.AnimeEpisodeVideoResponse {
	if len(asv) > 0 {
		Videos := make([]*aspbv1.AnimeEpisodeVideoResponse, len(asv))

		for i, v := range asv {
			Videos[i] = &aspbv1.AnimeEpisodeVideoResponse{
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

func convertAnimeEpisodeTorrents(ast []db.AnimeEpisodeTorrent) []*aspbv1.AnimeEpisodeTorrentResponse {
	if len(ast) > 0 {
		Torrents := make([]*aspbv1.AnimeEpisodeTorrentResponse, len(ast))

		for i, t := range ast {
			Torrents[i] = &aspbv1.AnimeEpisodeTorrentResponse{
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
