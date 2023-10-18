package db

import (
	"context"
	"testing"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAnimeMovie(t *testing.T) AnimeMovie {
	arg := CreateAnimeMovieParams{
		OriginalTitle: utils.RandomString(10),
		Aired:         time.Now(),
		ReleaseYear:   utils.RandomInt(1900, 2020),
		Duration:      time.Duration(100000),
	}

	anime, err := testGojo.CreateAnimeMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime)

	require.Equal(t, arg.OriginalTitle, anime.OriginalTitle)
	require.Equal(t, arg.ReleaseYear, anime.ReleaseYear)
	require.Equal(t, arg.Duration, anime.Duration)
	require.WithinDuration(t, arg.Aired, anime.Aired, time.Second)
	require.NotZero(t, anime.ID)
	require.NotZero(t, anime.CreatedAt)

	return anime
}

func TestCreateAnimeMovie(t *testing.T) {
	createRandomAnimeMovie(t)
}

func TestGetAnimeMovie(t *testing.T) {
	anime1 := createRandomAnimeMovie(t)
	anime2, err := testGojo.GetAnimeMovie(context.Background(), anime1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, anime2)

	require.Equal(t, anime1.OriginalTitle, anime2.OriginalTitle)
	require.Equal(t, anime1.ReleaseYear, anime2.ReleaseYear)
	require.Equal(t, anime1.Duration, anime2.Duration)
	require.Equal(t, anime1.ID, anime2.ID)
	require.WithinDuration(t, anime1.Aired, anime2.Aired, time.Second)
	require.WithinDuration(t, anime1.CreatedAt, anime2.CreatedAt, time.Second)
}

func TestUpdateAnimeMovie(t *testing.T) {
	anime1 := createRandomAnimeMovie(t)
	arg := UpdateAnimeMovieParams{
		ID: anime1.ID,
		OriginalTitle: pgtype.Text{
			String: createRandomAnimeMovie(t).OriginalTitle,
			Valid:  true,
		},
		Aired: pgtype.Timestamptz{
			Time:  createRandomAnimeMovie(t).Aired,
			Valid: true,
		},
		ReleaseYear: pgtype.Int4{
			Int32: createRandomAnimeMovie(t).ReleaseYear,
			Valid: true,
		},
		Duration: pgtype.Interval{
			Microseconds: createRandomAnimeMovie(t).Duration.Microseconds(),
			Valid:        true,
		},
	}
	anime2, err := testGojo.UpdateAnimeMovie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime2)

	require.NotEqual(t, anime1.OriginalTitle, anime2.OriginalTitle)
	require.NotEqual(t, anime1.ReleaseYear, anime2.ReleaseYear)
	require.Equal(t, anime1.Duration, anime2.Duration)
	require.Equal(t, anime1.ID, anime2.ID)
	require.WithinDuration(t, anime1.CreatedAt, anime2.CreatedAt, time.Second)
}

func TestDeleteAnimeMovie(t *testing.T) {
	anime1 := createRandomAnimeMovie(t)
	err := testGojo.DeleteAnimeMovie(context.Background(), anime1.ID)
	require.NoError(t, err)
}
