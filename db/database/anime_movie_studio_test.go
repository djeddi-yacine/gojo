package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnimeMovieStudio(t *testing.T) AnimeMovieStudio {
	anime := createRandomAnimeMovie(t)
	studio := createRandomStudio(t)
	arg := CreateAnimeMovieStudioParams{
		AnimeID:  anime.ID,
		StudioID: studio.ID,
	}

	animeMovieStudio, err := testGojo.CreateAnimeMovieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeMovieStudio)

	require.Equal(t, arg.AnimeID, animeMovieStudio.AnimeID)
	require.Equal(t, arg.StudioID, animeMovieStudio.StudioID)
	require.NotZero(t, animeMovieStudio.ID)
	require.NotZero(t, animeMovieStudio.StudioID)

	return animeMovieStudio
}

func TestCreateAnimeMovieStudio(t *testing.T) {
	createRandomAnimeMovieStudio(t)
}

func TestGetAnimeMovieStudio(t *testing.T) {
	a := createRandomAnimeMovie(t)
	s := createRandomStudio(t)
	arg := CreateAnimeMovieStudioParams{
		AnimeID:  a.ID,
		StudioID: s.ID,
	}

	studio1, err := testGojo.CreateAnimeMovieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studio1)

	studio2, err := testGojo.GetAnimeMovieStudio(context.Background(), GetAnimeMovieStudioParams{
		AnimeID: a.ID,
		StudioID: studio1.StudioID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, studio2)

	require.Equal(t, studio1.AnimeID, studio2.AnimeID)
	require.Equal(t, studio1.StudioID, studio2.StudioID)
}

func TestDeleteAnimeMovieStudio(t *testing.T) {
	a := createRandomAnimeMovie(t)
	s := createRandomStudio(t)
	arg1 := CreateAnimeMovieStudioParams{
		AnimeID:  a.ID,
		StudioID: s.ID,
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
			AnimeID:  a.ID,
			StudioID: studio.ID,
		}
		testGojo.CreateAnimeMovieStudio(context.Background(), arg)
	}

	studios, err := testGojo.ListAnimeMovieStudios(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotNil(t, studios)
}
