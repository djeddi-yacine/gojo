package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAnimeMovieGenre(t *testing.T) AnimeMovieGenre {
	anime := createRandomAnimeMovie(t)
	genre := createRandomGenre(t)
	arg := CreateAnimeMovieGenreParams{
		AnimeID: anime.ID,
		GenreID: pgtype.Int4{
			Int32: genre.ID,
			Valid: true,
		},
	}

	animeMovieGenre, err := testGojo.CreateAnimeMovieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeMovieGenre)

	require.Equal(t, arg.AnimeID, animeMovieGenre.AnimeID)
	require.Equal(t, arg.GenreID, animeMovieGenre.GenreID)
	require.NotZero(t, animeMovieGenre.ID)
	require.NotZero(t, animeMovieGenre.GenreID.Int32)
	require.True(t, animeMovieGenre.GenreID.Valid)

	return animeMovieGenre
}

func TestCreateAnimeMovieGenre(t *testing.T) {
	createRandomAnimeMovieGenre(t)
}

func TestGetAnimeMovieGenre(t *testing.T) {
	a := createRandomAnimeMovie(t)
	g := createRandomGenre(t)
	arg := CreateAnimeMovieGenreParams{
		AnimeID: a.ID,
		GenreID: pgtype.Int4{
			Int32: g.ID,
			Valid: true,
		},
	}

	Genre1, err := testGojo.CreateAnimeMovieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Genre1)

	Genre2, err := testGojo.GetAnimeMovieGenre(context.Background(), Genre1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Genre2)

	require.Equal(t, Genre1.AnimeID, Genre2.AnimeID)
	require.Equal(t, Genre1.GenreID, Genre2.GenreID)
}

func TestDeleteAnimeMovieGenre(t *testing.T) {
	a := createRandomAnimeMovie(t)
	g := createRandomGenre(t)
	arg1 := CreateAnimeMovieGenreParams{
		AnimeID: a.ID,
		GenreID: pgtype.Int4{
			Int32: g.ID,
			Valid: true,
		},
	}

	genre, err := testGojo.CreateAnimeMovieGenre(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, genre)

	arg2 := DeleteAnimeMovieGenreParams{
		AnimeID: genre.AnimeID,
		GenreID: genre.GenreID,
	}

	err = testGojo.DeleteAnimeMovieGenre(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeMovieGenres(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		genre := createRandomGenre(t)
		arg := CreateAnimeMovieGenreParams{
			AnimeID: a.ID,
			GenreID: pgtype.Int4{
				Int32: genre.ID,
				Valid: true,
			},
		}
		testGojo.CreateAnimeMovieGenre(context.Background(), arg)
	}

	arg := ListAnimeMovieGenresParams{
		AnimeID: a.ID,
		Limit:   3,
		Offset:  0,
	}

	genres, err := testGojo.ListAnimeMovieGenres(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, genres)
	require.Len(t, genres, 3)
}
