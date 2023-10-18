package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAnimeMovieStudio(t *testing.T) AnimeMovieStudio {
	anime := createRandomAnimeMovie(t)
	studio := createRandomStudio(t)
	arg := CreateAnimeMovieStudioParams{
		AnimeID: anime.ID,
		StudioID: pgtype.Int4{
			Int32: studio.ID,
			Valid: true,
		},
	}

	animeMovieStudio, err := testGojo.CreateAnimeMovieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeMovieStudio)

	require.Equal(t, arg.AnimeID, animeMovieStudio.AnimeID)
	require.Equal(t, arg.StudioID, animeMovieStudio.StudioID)
	require.NotZero(t, animeMovieStudio.ID)
	require.NotZero(t, animeMovieStudio.StudioID.Int32)
	require.True(t, animeMovieStudio.StudioID.Valid)

	return animeMovieStudio
}

func TestCreateAnimeMovieStudio(t *testing.T) {
	createRandomAnimeMovieStudio(t)
}

func TestGetAnimeMovieStudio(t *testing.T) {
	a := createRandomAnimeMovie(t)
	s := createRandomStudio(t)
	arg := CreateAnimeMovieStudioParams{
		AnimeID: a.ID,
		StudioID: pgtype.Int4{
			Int32: s.ID,
			Valid: true,
		},
	}

	studio1, err := testGojo.CreateAnimeMovieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studio1)

	studio2, err := testGojo.GetAnimeMovieStudio(context.Background(), studio1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, studio2)

	require.Equal(t, studio1.AnimeID, studio2.AnimeID)
	require.Equal(t, studio1.StudioID, studio2.StudioID)
}

func TestDeleteAnimeMovieStudio(t *testing.T) {
	a := createRandomAnimeMovie(t)
	s := createRandomStudio(t)
	arg1 := CreateAnimeMovieStudioParams{
		AnimeID: a.ID,
		StudioID: pgtype.Int4{
			Int32: s.ID,
			Valid: true,
		},
	}

	studio, err := testGojo.CreateAnimeMovieStudio(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, studio)

	arg2 := DeleteAnimeMovieStudioParams{
		AnimeID:  studio.AnimeID,
		StudioID: studio.StudioID,
	}

	err = testGojo.DeleteAnimeMovieStudio(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeMovieStudios(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		studio := createRandomStudio(t)
		arg := CreateAnimeMovieStudioParams{
			AnimeID: a.ID,
			StudioID: pgtype.Int4{
				Int32: studio.ID,
				Valid: true,
			},
		}
		testGojo.CreateAnimeMovieStudio(context.Background(), arg)
	}

	arg := ListAnimeMovieStudiosParams{
		AnimeID: a.ID,
		Limit:   3,
		Offset:  0,
	}

	studios, err := testGojo.ListAnimeMovieStudios(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, studios)
	require.Len(t, studios, 3)
}
