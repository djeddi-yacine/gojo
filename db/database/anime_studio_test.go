package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomAnimeStudio(t *testing.T) AnimeStudio {
	anime := createRandomAnimeMovie(t)
	studio := createRandomStudio(t)
	arg := CreateAnimeStudioParams{
		AnimeID: anime.ID,
		StudioID: pgtype.Int4{
			Int32: studio.ID,
			Valid: true,
		},
	}

	animeStudio, err := testGojo.CreateAnimeStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeStudio)

	require.Equal(t, arg.AnimeID, animeStudio.AnimeID)
	require.Equal(t, arg.StudioID, animeStudio.StudioID)
	require.NotZero(t, animeStudio.ID)
	require.NotZero(t, animeStudio.StudioID.Int32)
	require.True(t, animeStudio.StudioID.Valid)

	return animeStudio
}

func TestCreateAnimeStudio(t *testing.T) {
	createRandomAnimeStudio(t)
}

func TestGetAnimeStudio(t *testing.T) {
	a := createRandomAnimeMovie(t)
	s := createRandomStudio(t)
	arg := CreateAnimeStudioParams{
		AnimeID: a.ID,
		StudioID: pgtype.Int4{
			Int32: s.ID,
			Valid: true,
		},
	}

	studio1, err := testGojo.CreateAnimeStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studio1)

	studio2, err := testGojo.GetAnimeStudio(context.Background(), studio1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, studio2)

	require.Equal(t, studio1.AnimeID, studio2.AnimeID)
	require.Equal(t, studio1.StudioID, studio2.StudioID)
}

func TestDeleteAnimeStudio(t *testing.T) {
	a := createRandomAnimeMovie(t)
	s := createRandomStudio(t)
	arg1 := CreateAnimeStudioParams{
		AnimeID: a.ID,
		StudioID: pgtype.Int4{
			Int32: s.ID,
			Valid: true,
		},
	}

	studio, err := testGojo.CreateAnimeStudio(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, studio)

	arg2 := DeleteAnimeStudioParams{
		AnimeID:  studio.AnimeID,
		StudioID: studio.StudioID,
	}

	err = testGojo.DeleteAnimeStudio(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeStudios(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		studio := createRandomStudio(t)
		arg := CreateAnimeStudioParams{
			AnimeID: a.ID,
			StudioID: pgtype.Int4{
				Int32: studio.ID,
				Valid: true,
			},
		}
		testGojo.CreateAnimeStudio(context.Background(), arg)
	}

	arg := ListAnimeStudiosParams{
		AnimeID: a.ID,
		Limit:   3,
		Offset:  0,
	}

	studios, err := testGojo.ListAnimeStudios(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, studios)
	require.Len(t, studios, 3)
}
