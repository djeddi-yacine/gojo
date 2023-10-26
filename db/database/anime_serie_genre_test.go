package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnimeSerieGenre(t *testing.T) AnimeSerieGenre {
	anime := createRandomAnimeSerie(t)
	genre := createRandomGenre(t)
	arg := CreateAnimeSerieGenreParams{
		AnimeID: anime.ID,
		GenreID: genre.ID,
	}

	animeSerieGenre, err := testGojo.CreateAnimeSerieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeSerieGenre)

	require.Equal(t, arg.AnimeID, animeSerieGenre.AnimeID)
	require.Equal(t, arg.GenreID, animeSerieGenre.GenreID)
	require.NotZero(t, animeSerieGenre.ID)
	require.NotZero(t, animeSerieGenre.GenreID)

	return animeSerieGenre
}

func TestCreateAnimeSerieGenre(t *testing.T) {
	createRandomAnimeSerieGenre(t)
}

func TestGetAnimeSerieGenre(t *testing.T) {
	a := createRandomAnimeSerie(t)
	g := createRandomGenre(t)
	arg := CreateAnimeSerieGenreParams{
		AnimeID: a.ID,
		GenreID: g.ID,
	}

	Genre1, err := testGojo.CreateAnimeSerieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Genre1)

	Genre2, err := testGojo.GetAnimeSerieGenre(context.Background(), Genre1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Genre2)

	require.Equal(t, Genre1.AnimeID, Genre2.AnimeID)
	require.Equal(t, Genre1.GenreID, Genre2.GenreID)
}

func TestDeleteAnimeSerieGenre(t *testing.T) {
	a := createRandomAnimeSerie(t)
	g := createRandomGenre(t)
	arg1 := CreateAnimeSerieGenreParams{
		AnimeID: a.ID,
		GenreID: g.ID,
	}

	genre, err := testGojo.CreateAnimeSerieGenre(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, genre)

	arg2 := DeleteAnimeSerieGenreParams{
		AnimeID: genre.AnimeID,
		GenreID: genre.GenreID,
	}

	err = testGojo.DeleteAnimeSerieGenre(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeSerieGenres(t *testing.T) {
	a := createRandomAnimeSerie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		genre := createRandomGenre(t)
		arg := CreateAnimeSerieGenreParams{
			AnimeID: a.ID,
			GenreID: genre.ID,
		}
		testGojo.CreateAnimeSerieGenre(context.Background(), arg)
	}

	arg := ListAnimeSerieGenresParams{
		AnimeID: a.ID,
		Limit:   3,
		Offset:  0,
	}

	genres, err := testGojo.ListAnimeSerieGenres(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, genres)
	require.Len(t, genres, 3)
}
