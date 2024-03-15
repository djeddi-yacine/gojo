package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

type AnimeCharactersAndActors struct {
	Character AnimeCharacter
	Actor     []Actor
}

func (gojo *SQLGojo) ListAnimeCharacetrsTx(ctx context.Context, arg []int64) ([]AnimeCharactersAndActors, error) {
	var result []AnimeCharactersAndActors

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result = make([]AnimeCharactersAndActors, len(arg))
		for i, x := range arg {
			if err = gojo.ping.Handle(ctx, ping.SegmentKey(x).CHR(ping.Anime), &result[i].Character, func() error {
				result[i].Character, err = q.GetAnimeCharacter(ctx, x)
				if err != nil {
					return err
				}

				return nil
			}); err != nil {
				ErrorSQL(err)
				return err
			}

			r, err := q.ListAnimeCharacterActors(ctx, result[i].Character.ID)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result[i].Actor = make([]Actor, len(r))
			for j, z := range r {
				if err = gojo.ping.Handle(ctx, ping.SegmentKey(z.ActorID).ACT(), &result[i].Actor[j], func() error {
					result[i].Actor[j], err = q.GetActor(ctx, z.ActorID)
					if err != nil {
						return err
					}

					return nil
				}); err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		return err
	})

	return result, err
}
