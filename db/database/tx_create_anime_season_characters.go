package db

import (
	"context"
)

type CreateAnimeSeasonCharactersTxParams struct {
	SeasonID                     int64
	AnimeCharacterActorsTxParams []AnimeCharacterActorsTxParams
}

type CreateAnimeSeasonCharactersTxResult struct {
	Characters []AnimeCharacterActorsTx
}

func (gojo *SQLGojo) CreateAnimeSeasonCharactersTx(ctx context.Context, arg CreateAnimeSeasonCharactersTxParams) (CreateAnimeSeasonCharactersTxResult, error) {
	var result CreateAnimeSeasonCharactersTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

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

			_, err = q.CreateAnimeSeasonCharacter(ctx, CreateAnimeSeasonCharacterParams{
				SeasonID:    arg.SeasonID,
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
