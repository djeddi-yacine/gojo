package shv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertGenre(genre db.Genre) *nfpbv1.Genre {
	return &nfpbv1.Genre{
		GenreID:   genre.ID,
		GenreName: genre.GenreName,
		CreatedAt: timestamppb.New(genre.CreatedAt),
	}
}

func ConvertStudio(studio db.Studio) *nfpbv1.Studio {
	return &nfpbv1.Studio{
		StudioID:   studio.ID,
		StudioName: studio.StudioName,
		CreatedAt:  timestamppb.New(studio.CreatedAt),
	}
}

func ConvertLanguage(language db.Language) *nfpbv1.LanguageResponse {
	return &nfpbv1.LanguageResponse{
		LanguageID:   language.ID,
		LanguageCode: language.LanguageCode,
		LanguageName: language.LanguageName,
		CreatedAt:    timestamppb.New(language.CreatedAt),
	}
}

func ConvertMeta(meta db.Meta) *nfpbv1.MetaResponse {
	return &nfpbv1.MetaResponse{
		MetaID:   meta.ID,
		Title:    meta.Title,
		Overview: meta.Overview,
	}
}

func ConvertGenres(gg []db.Genre) []*nfpbv1.Genre {
	if len(gg) > 0 {
		Genres := make([]*nfpbv1.Genre, len(gg))

		for i, g := range gg {
			Genres[i] = &nfpbv1.Genre{
				GenreID:   g.ID,
				GenreName: g.GenreName,
				CreatedAt: timestamppb.New(g.CreatedAt),
			}
		}

		return Genres
	} else {
		return nil
	}
}

func ConvertStudios(ss []db.Studio) []*nfpbv1.Studio {
	if len(ss) > 0 {
		Studios := make([]*nfpbv1.Studio, len(ss))

		for i, s := range ss {
			Studios[i] = &nfpbv1.Studio{
				StudioID:   s.ID,
				StudioName: s.StudioName,
				CreatedAt:  timestamppb.New(s.CreatedAt),
			}
		}

		return Studios
	} else {
		return nil
	}
}
