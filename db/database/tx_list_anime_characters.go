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
		var cache ping.SegmentKey

		result = make([]AnimeCharactersAndActors, len(arg))
		for i, x := range arg {
			cache = ping.SegmentKey(x)
			if err = gojo.ping.Handle(ctx, cache.CHR(ping.Anime), &result[i].Character, func() error {
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
				cache = ping.SegmentKey(z.ActorID)
				if err = gojo.ping.Handle(ctx, cache.ACT(), &result[i].Actor[j], func() error {
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
