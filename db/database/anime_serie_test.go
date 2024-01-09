package db

import (
	"context"
	"testing"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAnimeSerie(t *testing.T) AnimeSerie {
	arg := CreateAnimeSerieParams{
		OriginalTitle: utils.RandomString(10),
		UniqueID:      uuid.New(),
		FirstYear:     utils.RandomInt(1900, 2020),
		MalID:         utils.RandomInt(01, 999999),
		TvdbID:        utils.RandomInt(01, 999999),
		TmdbID:        utils.RandomInt(01, 999999),
	}

	anime, err := testGojo.CreateAnimeSerie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime)

	require.Equal(t, arg.OriginalTitle, anime.OriginalTitle)
	require.Equal(t, arg.FirstYear, anime.FirstYear)
	require.NotZero(t, anime.ID)
	require.NotZero(t, anime.CreatedAt)

	return anime
}

func TestCreateAnimeSerie(t *testing.T) {
	createRandomAnimeSerie(t)
}

func TestGetAnimeSerie(t *testing.T) {
	anime1 := createRandomAnimeSerie(t)
	anime2, err := testGojo.GetAnimeSerie(context.Background(), anime1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, anime2)

	require.Equal(t, anime1.OriginalTitle, anime2.OriginalTitle)
	require.Equal(t, anime1.FirstYear, anime2.FirstYear)
	require.Equal(t, anime1.ID, anime2.ID)
	require.WithinDuration(t, anime1.CreatedAt, anime2.CreatedAt, time.Second)
}

func TestUpdateAnimeSerie(t *testing.T) {
	anime1 := createRandomAnimeSerie(t)
	arg := UpdateAnimeSerieParams{
		ID: anime1.ID,
		OriginalTitle: pgtype.Text{
			String: createRandomAnimeSerie(t).OriginalTitle,
			Valid:  true,
		},
		FirstYear: pgtype.Int4{
			Int32: createRandomAnimeSerie(t).FirstYear,
			Valid: true,
		},
	}
	anime2, err := testGojo.UpdateAnimeSerie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime2)

	require.NotEqual(t, anime1.OriginalTitle, anime2.OriginalTitle)
	require.NotEqual(t, anime1.FirstYear, anime2.FirstYear)
	require.Equal(t, anime1.ID, anime2.ID)
	require.WithinDuration(t, anime1.CreatedAt, anime2.CreatedAt, time.Second)
}

func TestDeleteAnimeSerie(t *testing.T) {
	anime1 := createRandomAnimeSerie(t)
	err := testGojo.DeleteAnimeSerie(context.Background(), anime1.ID)
	require.NoError(t, err)
}
