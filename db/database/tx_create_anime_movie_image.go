package db

import (
	"context"
)

type CreateAnimeMovieImageTxParams struct {
	AnimeID        int64
	AnimePosters   []CreateAnimeImageParams
	AnimeBackdrops []CreateAnimeImageParams
	AnimeLogos     []CreateAnimeImageParams
}

type CreateAnimeMovieImageTxResult struct {
	AnimeMovie     AnimeMovie
	AnimePosters   []AnimeImage
	AnimeBackdrops []AnimeImage
	AnimeLogos     []AnimeImage
}

func (gojo *SQLGojo) CreateAnimeMovieImageTx(ctx context.Context, arg CreateAnimeMovieImageTxParams) (CreateAnimeMovieImageTxResult, error) {
	var result CreateAnimeMovieImageTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeMovie = anime

		if arg.AnimePosters != nil {
			if len(arg.AnimePosters) > 0 {
				var imgArg CreateAnimeImageParams
				postersArg := make([]CreateAnimeMoviePosterImageParams, len(arg.AnimePosters))
				result.AnimePosters = make([]AnimeImage, len(arg.AnimePosters))

				for i, p := range arg.AnimePosters {
					imgArg = CreateAnimeImageParams{
						ImageHost:       p.ImageHost,
						ImageUrl:        p.ImageUrl,
						ImageThumbnails: p.ImageThumbnails,
						ImageBlurHash:   p.ImageBlurHash,
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

				for _, amp := range postersArg {
					_, err = q.CreateAnimeMoviePosterImage(ctx, amp)
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
				backdropsArg := make([]CreateAnimeMovieBackdropImageParams, len(arg.AnimeBackdrops))
				result.AnimeBackdrops = make([]AnimeImage, len(arg.AnimeBackdrops))

				for i, b := range arg.AnimeBackdrops {
					imgArg = CreateAnimeImageParams{
						ImageHost:       b.ImageHost,
						ImageUrl:        b.ImageUrl,
						ImageThumbnails: b.ImageThumbnails,
						ImageBlurHash:   b.ImageBlurHash,
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

				for _, amb := range backdropsArg {
					_, err = q.CreateAnimeMovieBackdropImage(ctx, amb)
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
				logosArg := make([]CreateAnimeMovieLogoImageParams, len(arg.AnimeLogos))
				result.AnimeLogos = make([]AnimeImage, len(arg.AnimeLogos))

				for i, l := range arg.AnimeLogos {
					imgArg = CreateAnimeImageParams{
						ImageHost:       l.ImageHost,
						ImageUrl:        l.ImageUrl,
						ImageThumbnails: l.ImageThumbnails,
						ImageBlurHash:   l.ImageBlurHash,
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

				for _, aml := range logosArg {
					_, err = q.CreateAnimeMovieLogoImage(ctx, aml)
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
