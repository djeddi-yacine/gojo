package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAnimeSerieStudio(t *testing.T) AnimeSerieStudio {
	anime := createRandomAnimeSerie(t)
	studio := createRandomStudio(t)
	arg := CreateAnimeSerieStudioParams{
		AnimeID: anime.ID,
		StudioID: pgtype.Int4{
			Int32: studio.ID,
			Valid: true,
		},
	}

	animeSerieStudio, err := testGojo.CreateAnimeSerieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeSerieStudio)

	require.Equal(t, arg.AnimeID, animeSerieStudio.AnimeID)
	require.Equal(t, arg.StudioID, animeSerieStudio.StudioID)
	require.NotZero(t, animeSerieStudio.ID)
	require.NotZero(t, animeSerieStudio.StudioID.Int32)
	require.True(t, animeSerieStudio.StudioID.Valid)

	return animeSerieStudio
}

func TestCreateAnimeSerieStudio(t *testing.T) {
	createRandomAnimeSerieStudio(t)
}

func TestGetAnimeSerieStudio(t *testing.T) {
	a := createRandomAnimeSerie(t)
	s := createRandomStudio(t)
	arg := CreateAnimeSerieStudioParams{
		AnimeID: a.ID,
		StudioID: pgtype.Int4{
			Int32: s.ID,
			Valid: true,
		},
	}

	studio1, err := testGojo.CreateAnimeSerieStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studio1)

	studio2, err := testGojo.GetAnimeSerieStudio(context.Background(), studio1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, studio2)

	require.Equal(t, studio1.AnimeID, studio2.AnimeID)
	require.Equal(t, studio1.StudioID, studio2.StudioID)
}

func TestDeleteAnimeSerieStudio(t *testing.T) {
	a := createRandomAnimeSerie(t)
	s := createRandomStudio(t)
	arg1 := CreateAnimeSerieStudioParams{
		AnimeID: a.ID,
		StudioID: pgtype.Int4{
			Int32: s.ID,
			Valid: true,
		},
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
			AnimeID: a.ID,
			StudioID: pgtype.Int4{
				Int32: studio.ID,
				Valid: true,
			},
		}
		testGojo.CreateAnimeSerieStudio(context.Background(), arg)
	}

	arg := ListAnimeSerieStudiosParams{
		AnimeID: a.ID,
		Limit:   3,
		Offset:  0,
	}

	studios, err := testGojo.ListAnimeSerieStudios(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, studios)
	require.Len(t, studios, 3)
}
