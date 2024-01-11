package aapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertAnimeResource(v db.AnimeResource) *ashpbv1.AnimeResourceResponse {
	if v.ID != 0 {
		return &ashpbv1.AnimeResourceResponse{
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

func ConvertAnimeLink(v db.AnimeLink) *ashpbv1.AnimeLinkResponse {
	if v.ID != 0 {
		return &ashpbv1.AnimeLinkResponse{
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

func ConvertAnimeImages(v []db.AnimeImage) []*ashpbv1.ImageResponse {
	if len(v) > 0 {
		images := make([]*ashpbv1.ImageResponse, len(v))

		for i, x := range v {
			images[i] = &ashpbv1.ImageResponse{
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

func ConvertAnimeTrailers(v []db.AnimeTrailer) []*ashpbv1.AnimeTrailerResponse {
	if len(v) > 0 {
		trailers := make([]*ashpbv1.AnimeTrailerResponse, len(v))

		for i, x := range v {
			trailers[i] = &ashpbv1.AnimeTrailerResponse{
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

func ConvertAnimeCharacter(v db.AnimeCharacter) *ashpbv1.AnimeCharacterResponse {
	if v.ID != 0 {
		return &ashpbv1.AnimeCharacterResponse{
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

func ConvertAnimeTags(v []db.AnimeTag) []*ashpbv1.AnimeTag {
	if len(v) > 0 {
		tags := make([]*ashpbv1.AnimeTag, len(v))

		for i, x := range v {
			tags[i] = &ashpbv1.AnimeTag{
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
