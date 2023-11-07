package shared

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/pb/shpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertGenre(genre db.Genre) *nfpb.Genre {
	return &nfpb.Genre{
		GenreID:   genre.ID,
		GenreName: genre.GenreName,
		CreatedAt: timestamppb.New(genre.CreatedAt),
	}
}

func ConvertStudio(studio db.Studio) *nfpb.Studio {
	return &nfpb.Studio{
		StudioID:   studio.ID,
		StudioName: studio.StudioName,
		CreatedAt:  timestamppb.New(studio.CreatedAt),
	}
}

func ConvertLanguage(language db.Language) *nfpb.LanguageResponse {
	return &nfpb.LanguageResponse{
		LanguageID:   language.ID,
		LanguageCode: language.LanguageCode,
		LanguageName: language.LanguageName,
		CreatedAt:    timestamppb.New(language.CreatedAt),
	}
}

func ConvertMeta(meta db.Meta) *nfpb.MetaResponse {
	return &nfpb.MetaResponse{
		MetaID:   meta.ID,
		Title:    meta.Title,
		Overview: meta.Overview,
	}
}

func ConvertAnimeMovie(a db.AnimeMovie) *ampb.AnimeMovieResponse {
	return &ampb.AnimeMovieResponse{
		ID:            a.ID,
		OriginalTitle: a.OriginalTitle,
		Aired:         timestamppb.New(a.Aired),
		ReleaseYear:   a.ReleaseYear,
		Rating:        a.Rating,
		Duration:      durationpb.New(a.Duration),
		CreatedAt:     timestamppb.New(a.CreatedAt),
	}
}

func ConvertAnimeSerie(a db.AnimeSerie) *aspb.AnimeSerieResponse {
	return &aspb.AnimeSerieResponse{
		ID:            a.ID,
		OriginalTitle: a.OriginalTitle,
		Aired:         timestamppb.New(a.Aired),
		ReleaseYear:   a.ReleaseYear,
		Rating:        a.Rating,
		Duration:      durationpb.New(a.Duration),
		CreatedAt:     timestamppb.New(a.CreatedAt),
	}
}

func ConvertAnimeResource(r db.AnimeResource) *shpb.AnimeResourceResponse {
	return &shpb.AnimeResourceResponse{
		ID:              r.ID,
		TMDbID:          r.TmdbID,
		IMDbID:          r.ImdbID,
		WikipediaUrl:    r.WikipediaUrl,
		OfficialWebsite: r.OfficialWebsite,
		CrunchyrollUrl:  r.CrunchyrollUrl,
		SocialMedia:     r.SocialMedia,
		CreatedAt:       timestamppb.New(r.CreatedAt),
	}
}

func ConvertAnimeMovieVideos(amv []db.AnimeMovieVideo) []*ampb.AnimeMovieVideoResponse {
	if amv != nil {
		Videos := make([]*ampb.AnimeMovieVideoResponse, len(amv))

		for i, v := range amv {
			Videos[i] = &ampb.AnimeMovieVideoResponse{
				ID:         v.ID,
				LanguageID: v.LanguageID,
				Autority:   v.Autority,
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

func ConvertAnimeMovieTorrents(amt []db.AnimeMovieTorrent) []*ampb.AnimeMovieTorrentResponse {
	if amt != nil {
		Torrents := make([]*ampb.AnimeMovieTorrentResponse, len(amt))

		for i, t := range amt {
			Torrents[i] = &ampb.AnimeMovieTorrentResponse{
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

func ConvertGenres(gg []db.Genre) *nfpb.AnimeGenresResponse {
	if gg != nil {
		Genres := make([]*nfpb.Genre, len(gg))

		for i, g := range gg {
			Genres[i] = &nfpb.Genre{
				GenreID:   g.ID,
				GenreName: g.GenreName,
				CreatedAt: timestamppb.New(g.CreatedAt),
			}
		}

		return &nfpb.AnimeGenresResponse{
			Genres: Genres,
		}
	} else {
		return nil
	}
}

func ConvertStudios(ss []db.Studio) *nfpb.AnimeStudiosResponse {
	if ss != nil {
		Studios := make([]*nfpb.Studio, len(ss))

		for i, s := range ss {
			Studios[i] = &nfpb.Studio{
				StudioID:   s.ID,
				StudioName: s.StudioName,
				CreatedAt:  timestamppb.New(s.CreatedAt),
			}
		}

		return &nfpb.AnimeStudiosResponse{
			Studios: Studios,
		}
	} else {
		return nil
	}
}
