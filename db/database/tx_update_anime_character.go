package db

import (
	"context"
)

type UpdateAnimeCharacterTxParams struct {
	AnimeCharacter UpdateAnimeCharacterParams
	ActorsIDs      []int64
}

type UpdateAnimeCharacterTxResult struct {
	AnimeCharacter AnimeCharacter
	ActorsIDs      []int64
}

func (gojo *SQLGojo) UpdateAnimeCharacterTx(ctx context.Context, arg UpdateAnimeCharacterTxParams) (UpdateAnimeCharacterTxResult, error) {
	var result UpdateAnimeCharacterTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		c, err := q.UpdateAnimeCharacter(ctx, arg.AnimeCharacter)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		actorsIDs := make([]int64, len(arg.ActorsIDs))

		for i, x := range arg.ActorsIDs {
			if x == 0 {
				continue
			}

			err = q.CreateAnimeCharacterActor(ctx, CreateAnimeCharacterActorParams{
				CharacterID: c.ID,
				ActorID:     x,
			})
			if err != nil {
				ErrorSQL(err)
				return err
			}
			actorsIDs[i] = x
		}

		result.AnimeCharacter = c
		result.ActorsIDs = actorsIDs

		return err
	})

	return result, err
}
