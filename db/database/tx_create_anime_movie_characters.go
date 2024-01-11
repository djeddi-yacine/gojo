package db

import (
	"context"
)

type AnimeCharacterActorsTxParams struct {
	ActorsIDs             []int64
	CreateAnimeCharacters CreateAnimeCharacterParams
}

type AnimeCharacterActorsTx struct {
	ActorsIDs      []int64
	AnimeCharacter AnimeCharacter
}

type CreateAnimeMovieCharactersTxParams struct {
	AnimeID                      int64
	AnimeCharacterActorsTxParams []AnimeCharacterActorsTxParams
}

type CreateAnimeMovieCharactersTxResult struct {
	AnimeMovie AnimeMovie
	Characters []AnimeCharacterActorsTx
}

func (gojo *SQLGojo) CreateAnimeMovieCharactersTx(ctx context.Context, arg CreateAnimeMovieCharactersTxParams) (CreateAnimeMovieCharactersTxResult, error) {
	var result CreateAnimeMovieCharactersTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeMovie = anime
		result.Characters = make([]AnimeCharacterActorsTx, len(arg.AnimeCharacterActorsTxParams))

		for i, x := range arg.AnimeCharacterActorsTxParams {
			c, err := q.CreateAnimeCharacter(ctx, x.CreateAnimeCharacters)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			for _, z := range x.ActorsIDs {
				err = q.CreateAnimeCharacterActor(ctx, CreateAnimeCharacterActorParams{
					CharacterID: c.ID,
					ActorID:     z,
				})
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}

			_, err = q.CreateAnimeMovieCharacter(ctx, CreateAnimeMovieCharacterParams{
				AnimeID:     anime.ID,
				CharacterID: c.ID,
			})
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result.Characters[i].AnimeCharacter = c
			result.Characters[i].ActorsIDs = x.ActorsIDs
		}

		return err
	})

	return result, err
}
