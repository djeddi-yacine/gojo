package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAnimeGenre(t *testing.T) AnimeGenre {
	anime := createRandomAnimeMovie(t)
	genre := createRandomGenre(t)
	arg := CreateAnimeGenreParams{
		AnimeID: anime.ID,
		GenreID: pgtype.Int4{
			Int32: genre.ID,
			Valid: true,
		},
	}

	animeGenre, err := testGojo.CreateAnimeGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeGenre)

	require.Equal(t, arg.AnimeID, animeGenre.AnimeID)
	require.Equal(t, arg.GenreID, animeGenre.GenreID)
	require.NotZero(t, animeGenre.ID)
	require.NotZero(t, animeGenre.GenreID.Int32)
	require.True(t, animeGenre.GenreID.Valid)

	return animeGenre
}

func TestCreateAnimeGenre(t *testing.T) {
	createRandomAnimeGenre(t)
}

func TestGetAnimeGenre(t *testing.T) {
	a := createRandomAnimeMovie(t)
	g := createRandomGenre(t)
	arg := CreateAnimeGenreParams{
		AnimeID: a.ID,
		GenreID: pgtype.Int4{
			Int32: g.ID,
			Valid: true,
		},
	}

	Genre1, err := testGojo.CreateAnimeGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Genre1)

	Genre2, err := testGojo.GetAnimeGenre(context.Background(), Genre1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Genre2)

	require.Equal(t, Genre1.AnimeID, Genre2.AnimeID)
	require.Equal(t, Genre1.GenreID, Genre2.GenreID)
}

func TestDeleteAnimeGenre(t *testing.T) {
	a := createRandomAnimeMovie(t)
	g := createRandomGenre(t)
	arg1 := CreateAnimeGenreParams{
		AnimeID: a.ID,
		GenreID: pgtype.Int4{
			Int32: g.ID,
			Valid: true,
		},
	}

	genre, err := testGojo.CreateAnimeGenre(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, genre)

	arg2 := DeleteAnimeGenreParams{
		AnimeID: genre.AnimeID,
		GenreID: genre.GenreID,
	}

	err = testGojo.DeleteAnimeGenre(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeGenres(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		genre := createRandomGenre(t)
		arg := CreateAnimeGenreParams{
			AnimeID: a.ID,
			GenreID: pgtype.Int4{
				Int32: genre.ID,
				Valid: true,
			},
		}
		testGojo.CreateAnimeGenre(context.Background(), arg)
	}

	arg := ListAnimeGenresParams{
		AnimeID: a.ID,
		Limit:   3,
		Offset:  0,
	}

	genres, err := testGojo.ListAnimeGenres(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, genres)
	require.Len(t, genres, 3)
}
