package db

import (
	"context"
	"testing"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/stretchr/testify/require"
)

func createRandomStudio(t *testing.T) Studio {
	studio, err := testGojo.CreateStudio(context.Background(), utils.RandomString(10))
	require.NoError(t, err)
	require.NotEmpty(t, studio)
	return studio
}

func TestCreateStudio(t *testing.T) {
	createRandomStudio(t)
}

func TestGetStudio(t *testing.T) {
	studio1 := createRandomStudio(t)
	studio2, err := testGojo.GetStudio(context.Background(), studio1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, studio2)

	require.Equal(t, studio1.ID, studio2.ID)
	require.Equal(t, studio1.StudioName, studio2.StudioName)
	require.WithinDuration(t, studio1.CreatedAt, studio2.CreatedAt, time.Second)
}

func TestUpdateStudio(t *testing.T) {
	studio1 := createRandomStudio(t)
	require.NotEmpty(t, studio1)

	arg := UpdateStudioParams{
		ID:         studio1.ID,
		StudioName: utils.RandomString(15),
	}

	studio2, err := testGojo.UpdateStudio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, studio2)

	require.Equal(t, studio1.ID, studio2.ID)
	require.NotEqual(t, studio1.StudioName, studio2.StudioName)
	require.WithinDuration(t, studio1.CreatedAt, studio2.CreatedAt, time.Second)
}

func TestDeleteStudio(t *testing.T) {
	studio1 := createRandomStudio(t)
	err := testGojo.DeleteStudio(context.Background(), studio1.ID)
	require.NoError(t, err)
}

func TestListStudios(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomStudio(t)
	}

	arg := ListStudiosParams{
		Limit:  3,
		Offset: 0,
	}

	studios, err := testGojo.ListStudios(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, studios, 3)
	require.NotNil(t, studios)

	for _, s := range studios {
		require.NotNil(t, s)
	}

}
