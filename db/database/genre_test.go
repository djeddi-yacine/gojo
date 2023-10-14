package db

import (
	"context"
	"testing"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/stretchr/testify/require"
)

func createRandomGenre(t *testing.T) Genre {
	genre, err := testGojo.CreateGenre(context.Background(), utils.RandomString(10))
	require.NoError(t, err)
	require.NotEmpty(t, genre)
	return genre
}

func TestCreateGenre(t *testing.T) {
	createRandomGenre(t)
}

func TestGetGenre(t *testing.T) {
	genre1 := createRandomGenre(t)
	genre2, err := testGojo.GetGenre(context.Background(), genre1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, genre2)

	require.Equal(t, genre1.ID, genre2.ID)
	require.Equal(t, genre1.GenreName, genre2.GenreName)
	require.WithinDuration(t, genre1.CreatedAt, genre2.CreatedAt, time.Second)
}

func TestUpdateGenre(t *testing.T) {
	genre1 := createRandomGenre(t)
	require.NotEmpty(t, genre1)

	arg := UpdateGenreParams{
		ID:        genre1.ID,
		GenreName: utils.RandomString(15),
	}

	genre2, err := testGojo.UpdateGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, genre2)

	require.Equal(t, genre1.ID, genre2.ID)
	require.NotEqual(t, genre1.GenreName, genre2.GenreName)
	require.WithinDuration(t, genre1.CreatedAt, genre2.CreatedAt, time.Second)
}

func TestDeleteGenre(t *testing.T) {
	genre1 := createRandomGenre(t)
	err := testGojo.DeleteGenre(context.Background(), genre1.ID)
	require.NoError(t, err)
}

func TestListGenres(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomGenre(t)
	}

	arg := ListGenresParams{
		Limit:  3,
		Offset: 0,
	}

	Genres, err := testGojo.ListGenres(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Genres, 3)
	require.NotNil(t, Genres)

	for _, s := range Genres {
		require.NotNil(t, s.ID)
		require.NotNil(t, s.GenreName)
		require.NotNil(t, s.CreatedAt)
	}

}
