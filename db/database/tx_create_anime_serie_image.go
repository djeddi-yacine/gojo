package db

import (
	"context"
)

type CreateAnimeSerieImageTxParams struct {
	AnimeID        int64
	AnimePosters   []CreateAnimeImageParams
	AnimeBackdrops []CreateAnimeImageParams
	AnimeLogos     []CreateAnimeImageParams
}

type CreateAnimeSerieImageTxResult struct {
	AnimeSerie     AnimeSerie
	AnimePosters   []AnimeImage
	AnimeBackdrops []AnimeImage
	AnimeLogos     []AnimeImage
}

func (gojo *SQLGojo) CreateAnimeSerieImageTx(ctx context.Context, arg CreateAnimeSerieImageTxParams) (CreateAnimeSerieImageTxResult, error) {
	var result CreateAnimeSerieImageTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeSerie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSerie = anime

		if arg.AnimePosters != nil {
			if len(arg.AnimePosters) > 0 {
				var imgArg CreateAnimeImageParams
				postersArg := make([]CreateAnimeSeriePosterImageParams, len(arg.AnimePosters))
				result.AnimePosters = make([]AnimeImage, len(arg.AnimePosters))

				for i, p := range arg.AnimePosters {
					imgArg = CreateAnimeImageParams{
						ImageHost:       p.ImageHost,
						ImageUrl:        p.ImageUrl,
						ImageThumbnails: p.ImageThumbnails,
						ImageBlurhash:   p.ImageBlurhash,
						ImageHeight:     p.ImageHeight,
						ImageWidth:      p.ImageWidth,
					}

					img, err := q.CreateAnimeImage(ctx, imgArg)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.AnimePosters[i] = img
					postersArg[i].AnimeID = anime.ID
					postersArg[i].ImageID = img.ID
				}

				for _, asp := range postersArg {
					_, err = q.CreateAnimeSeriePosterImage(ctx, asp)
					if err != nil {
						ErrorSQL(err)
						return err
					}
				}
			}
		}

		if arg.AnimeBackdrops != nil {
			if len(arg.AnimeBackdrops) > 0 {
				var imgArg CreateAnimeImageParams
				backdropsArg := make([]CreateAnimeSerieBackdropImageParams, len(arg.AnimeBackdrops))
				result.AnimeBackdrops = make([]AnimeImage, len(arg.AnimeBackdrops))

				for i, b := range arg.AnimeBackdrops {
					imgArg = CreateAnimeImageParams{
						ImageHost:       b.ImageHost,
						ImageUrl:        b.ImageUrl,
						ImageThumbnails: b.ImageThumbnails,
						ImageBlurhash:   b.ImageBlurhash,
						ImageHeight:     b.ImageHeight,
						ImageWidth:      b.ImageWidth,
					}

					img, err := q.CreateAnimeImage(ctx, imgArg)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.AnimeBackdrops[i] = img
					backdropsArg[i].AnimeID = anime.ID
					backdropsArg[i].ImageID = img.ID
				}

				for _, asb := range backdropsArg {
					_, err = q.CreateAnimeSerieBackdropImage(ctx, asb)
					if err != nil {
						ErrorSQL(err)
						return err
					}
				}
			}
		}

		if arg.AnimeLogos != nil {
			if len(arg.AnimeLogos) > 0 {
				var imgArg CreateAnimeImageParams
				logosArg := make([]CreateAnimeSerieLogoImageParams, len(arg.AnimeLogos))
				result.AnimeLogos = make([]AnimeImage, len(arg.AnimeLogos))

				for i, l := range arg.AnimeLogos {
					imgArg = CreateAnimeImageParams{
						ImageHost:       l.ImageHost,
						ImageUrl:        l.ImageUrl,
						ImageThumbnails: l.ImageThumbnails,
						ImageBlurhash:   l.ImageBlurhash,
						ImageHeight:     l.ImageHeight,
						ImageWidth:      l.ImageWidth,
					}

					img, err := q.CreateAnimeImage(ctx, imgArg)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.AnimeLogos[i] = img
					logosArg[i].AnimeID = anime.ID
					logosArg[i].ImageID = img.ID
				}

				for _, asl := range logosArg {
					_, err = q.CreateAnimeSerieLogoImage(ctx, asl)
					if err != nil {
						ErrorSQL(err)
						return err
					}
				}
			}
		}

		return err
	})

	return result, err
}
