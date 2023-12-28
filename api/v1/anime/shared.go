package aapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertAnimeResource(r db.AnimeResource) *ashpbv1.AnimeResourceResponse {
	return &ashpbv1.AnimeResourceResponse{
		ID:            r.ID,
		TvdbID:        r.TvdbID,
		TmdbID:        r.TmdbID,
		ImdbID:        r.ImdbID,
		LivechartID:   r.LivechartID,
		AnimePlanetID: r.AnimePlanetID,
		AnisearchID:   r.AnisearchID,
		AnidbID:       r.AnidbID,
		KitsuID:       r.KitsuID,
		MalID:         r.MalID,
		NotifyMoeID:   r.NotifyMoeID,
		AnilistID:     r.AnilistID,
		CreatedAt:     timestamppb.New(r.CreatedAt),
	}
}

func ConvertAnimeLink(l db.AnimeLink) *ashpbv1.AnimeLinkResponse {
	return &ashpbv1.AnimeLinkResponse{
		ID:              l.ID,
		OfficialWebsite: l.OfficialWebsite,
		WikipediaUrl:    l.WikipediaUrl,
		CrunchyrollUrl:  l.CrunchyrollUrl,
		SocialMedia:     l.SocialMedia,
		CreatedAt:       timestamppb.New(l.CreatedAt),
	}
}

func ConvertAnimeImages(ai []db.AnimeImage) []*ashpbv1.ImageResponse {
	if len(ai) > 0 {
		images := make([]*ashpbv1.ImageResponse, len(ai))

		for i, g := range ai {
			images[i] = &ashpbv1.ImageResponse{
				ID:         g.ID,
				Host:       g.ImageHost,
				Url:        g.ImageUrl,
				Thumbnails: g.ImageThumbnails,
				Blurhash:   g.ImageBlurhash,
				Height:     uint32(g.ImageHeight),
				Width:      uint32(g.ImageWidth),
				CreatedAt:  timestamppb.New(g.CreatedAt),
			}
		}

		return images
	} else {
		return nil
	}
}

func ConvertAnimeTrailers(at []db.AnimeTrailer) []*ashpbv1.AnimeTrailerResponse {
	if len(at) > 0 {
		trailers := make([]*ashpbv1.AnimeTrailerResponse, len(at))

		for i, t := range at {
			trailers[i] = &ashpbv1.AnimeTrailerResponse{
				ID:         t.ID,
				IsOfficial: t.IsOfficial,
				HostName:   t.HostName,
				HostKey:    t.HostKey,
				CreatedAt:  timestamppb.New(t.CreatedAt),
			}
		}

		return trailers
	} else {
		return nil
	}
}
