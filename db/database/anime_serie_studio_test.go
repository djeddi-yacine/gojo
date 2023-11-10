package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnimeSerieStudio(t *testing.T) AnimeSerieStudio {
	anime := createRandomAnimeSerie(t)
	studio := createRandomStudio(t)
	arg := CreateAnimeSerieStudioParams{
		AnimeID:  anime.ID,
		StudioID: studio.ID,
	}

	animeSerieStudio, err := testGojo.CreateAnimeSerieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeSerieStudio)

	require.Equal(t, arg.AnimeID, animeSerieStudio.AnimeID)
	require.Equal(t, arg.StudioID, animeSerieStudio.StudioID)
	require.NotZero(t, animeSerieStudio.ID)
	require.NotZero(t, animeSerieStudio.StudioID)

	return animeSerieStudio
}

func TestCreateAnimeSerieStudio(t *testing.T) {
	createRandomAnimeSerieStudio(t)
}

func TestGetAnimeSerieStudio(t *testing.T) {
	a := createRandomAnimeSerie(t)
	s := createRandomStudio(t)
	arg := CreateAnimeSerieStudioParams{
		AnimeID:  a.ID,
		StudioID: s.ID,
	}

	studio1, err := testGojo.CreateAnimeSerieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studio1)

	studio2, err := testGojo.GetAnimeSerieStudio(context.Background(), GetAnimeSerieStudioParams{
		AnimeID:  a.ID,
		StudioID: studio1.StudioID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, studio2)

	require.Equal(t, studio1.AnimeID, studio2.AnimeID)
	require.Equal(t, studio1.StudioID, studio2.StudioID)
}

func TestDeleteAnimeSerieStudio(t *testing.T) {
	a := createRandomAnimeSerie(t)
	s := createRandomStudio(t)
	arg1 := CreateAnimeSerieStudioParams{
		AnimeID:  a.ID,
		StudioID: s.ID,
	}

	studio, err := testGojo.CreateAnimeSerieStudio(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, studio)

	arg2 := DeleteAnimeSerieStudioParams{
		AnimeID:  studio.AnimeID,
		StudioID: studio.StudioID,
	}

	err = testGojo.DeleteAnimeSerieStudio(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeSerieStudios(t *testing.T) {
	a := createRandomAnimeSerie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		studio := createRandomStudio(t)
		arg := CreateAnimeSerieStudioParams{
			AnimeID:  a.ID,
			StudioID: studio.ID,
		}
		testGojo.CreateAnimeSerieStudio(context.Background(), arg)
	}

	studios, err := testGojo.ListAnimeSerieStudios(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotNil(t, studios)
}
