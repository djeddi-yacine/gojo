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
