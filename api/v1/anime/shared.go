package av1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertAnimeResource(v db.AnimeResource) *apbv1.AnimeResourceResponse {
	if v.ID != 0 {
		return &apbv1.AnimeResourceResponse{
			ID:            v.ID,
			TvdbID:        v.TvdbID,
			TmdbID:        v.TmdbID,
			ImdbID:        v.ImdbID,
			LivechartID:   v.LivechartID,
			AnimePlanetID: v.AnimePlanetID,
			AnisearchID:   v.AnisearchID,
			AnidbID:       v.AnidbID,
			KitsuID:       v.KitsuID,
			MalID:         v.MalID,
			NotifyMoeID:   v.NotifyMoeID,
			AnilistID:     v.AnilistID,
			CreatedAt:     timestamppb.New(v.CreatedAt),
		}
	}
	return nil
}

func ConvertAnimeLink(v db.AnimeLink) *apbv1.AnimeLinkResponse {
	if v.ID != 0 {
		return &apbv1.AnimeLinkResponse{
			ID:              v.ID,
			OfficialWebsite: v.OfficialWebsite,
			WikipediaUrl:    v.WikipediaUrl,
			CrunchyrollUrl:  v.CrunchyrollUrl,
			SocialMedia:     v.SocialMedia,
			CreatedAt:       timestamppb.New(v.CreatedAt),
		}
	}

	return nil
}

func ConvertAnimeImages(v []db.AnimeImage) []*apbv1.ImageResponse {
	if len(v) > 0 {
		images := make([]*apbv1.ImageResponse, len(v))

		for i, x := range v {
			images[i] = &apbv1.ImageResponse{
				ID:         x.ID,
				Host:       x.ImageHost,
				Url:        x.ImageUrl,
				Thumbnails: x.ImageThumbnails,
				Blurhash:   x.ImageBlurhash,
				Height:     uint32(x.ImageHeight),
				Width:      uint32(x.ImageWidth),
				CreatedAt:  timestamppb.New(x.CreatedAt),
			}
		}

		return images
	} else {
		return nil
	}
}

func ConvertAnimeTrailers(v []db.AnimeTrailer) []*apbv1.AnimeTrailerResponse {
	if len(v) > 0 {
		trailers := make([]*apbv1.AnimeTrailerResponse, len(v))

		for i, x := range v {
			trailers[i] = &apbv1.AnimeTrailerResponse{
				ID:         x.ID,
				IsOfficial: x.IsOfficial,
				HostName:   x.HostName,
				HostKey:    x.HostKey,
				CreatedAt:  timestamppb.New(x.CreatedAt),
			}
		}

		return trailers
	} else {
		return nil
	}
}

func ConvertAnimeCharacter(v db.AnimeCharacter) *apbv1.AnimeCharacterResponse {
	if v.ID != 0 {
		return &apbv1.AnimeCharacterResponse{
			ID:            v.ID,
			FullName:      v.FullName,
			About:         v.About,
			RolePlaying:   v.RolePlaying,
			Image:         v.ImageUrl,
			ImageBlurHash: v.ImageBlurHash,
			Pictures:      v.Pictures,
			CreatedAt:     timestamppb.New(v.CreatedAt),
		}
	}

	return nil
}

func ConvertAnimeTags(v []db.AnimeTag) []*apbv1.AnimeTag {
	if len(v) > 0 {
		tags := make([]*apbv1.AnimeTag, len(v))

		for i, x := range v {
			tags[i] = &apbv1.AnimeTag{
				ID:        x.ID,
				Tag:       x.Tag,
				CreatedAt: timestamppb.New(x.CreatedAt),
			}
		}

		return tags
	} else {
		return nil
	}
}
