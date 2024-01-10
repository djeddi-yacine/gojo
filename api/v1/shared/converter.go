package shv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertGenre(v db.Genre) *nfpbv1.Genre {
	return &nfpbv1.Genre{
		GenreID:   v.ID,
		GenreName: v.GenreName,
		CreatedAt: timestamppb.New(v.CreatedAt),
	}
}

func ConvertStudio(studio db.Studio) *nfpbv1.Studio {
	return &nfpbv1.Studio{
		StudioID:   studio.ID,
		StudioName: studio.StudioName,
		CreatedAt:  timestamppb.New(studio.CreatedAt),
	}
}

func ConvertActor(v db.Actor) *nfpbv1.ActorResponse {
	return &nfpbv1.ActorResponse{
		ActorID:       v.ID,
		FullName:      v.FullName,
		Gender:        v.Gender,
		Biography:     v.Biography,
		Born:          timestamppb.New(v.Born),
		Image:         v.ImageUrl,
		ImageBlurHash: v.ImageBlurHash,
		CreatedAt:     timestamppb.New(v.CreatedAt),
	}
}

func ConvertLanguage(v db.Language) *nfpbv1.LanguageResponse {
	return &nfpbv1.LanguageResponse{
		LanguageID:   v.ID,
		LanguageCode: v.LanguageCode,
		LanguageName: v.LanguageName,
		CreatedAt:    timestamppb.New(v.CreatedAt),
	}
}

func ConvertMeta(v db.Meta) *nfpbv1.MetaResponse {
	return &nfpbv1.MetaResponse{
		MetaID:   v.ID,
		Title:    v.Title,
		Overview: v.Overview,
	}
}

func ConvertGenres(v []db.Genre) []*nfpbv1.Genre {
	if len(v) > 0 {
		genres := make([]*nfpbv1.Genre, len(v))

		for i, x := range v {
			genres[i] = &nfpbv1.Genre{
				GenreID:   x.ID,
				GenreName: x.GenreName,
				CreatedAt: timestamppb.New(x.CreatedAt),
			}
		}

		return genres
	} else {
		return nil
	}
}

func ConvertStudios(v []db.Studio) []*nfpbv1.Studio {
	if len(v) > 0 {
		studios := make([]*nfpbv1.Studio, len(v))

		for i, x := range v {
			studios[i] = &nfpbv1.Studio{
				StudioID:   x.ID,
				StudioName: x.StudioName,
				CreatedAt:  timestamppb.New(x.CreatedAt),
			}
		}

		return studios
	} else {
		return nil
	}
}

func ConvertActors(v []db.Actor) []*nfpbv1.ActorResponse {
	if len(v) > 0 {
		actors := make([]*nfpbv1.ActorResponse, len(v))

		for i, x := range v {
			actors[i] = &nfpbv1.ActorResponse{
				ActorID:       x.ID,
				FullName:      x.FullName,
				Gender:        x.Gender,
				Biography:     x.Biography,
				Born:          timestamppb.New(x.Born),
				Image:         x.ImageUrl,
				ImageBlurHash: x.ImageBlurHash,
				CreatedAt:     timestamppb.New(x.CreatedAt),
			}
		}

		return actors
	} else {
		return nil
	}
}
