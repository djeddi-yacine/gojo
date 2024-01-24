package db

import (
	"context"
	"testing"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomMeta(t *testing.T) Meta {
	arg := CreateMetaParams{
		Title:    utils.RandomString(10),
		Overview: utils.RandomString(100),
	}
	meta, err := testGojo.CreateMeta(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, meta)
	return meta
}

func TestCreateMeta(t *testing.T) {
	createRandomMeta(t)
}

func TestGetMeta(t *testing.T) {
	Meta1 := createRandomMeta(t)
	Meta2, err := testGojo.GetMeta(context.Background(), Meta1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Meta2)

	require.Equal(t, Meta1.ID, Meta2.ID)
	require.Equal(t, Meta1.Title, Meta2.Title)
	require.Equal(t, Meta1.Overview, Meta2.Overview)
	require.WithinDuration(t, Meta1.CreatedAt, Meta2.CreatedAt, time.Second)
}

func TestUpdateMeta(t *testing.T) {
	Meta1 := createRandomMeta(t)
	require.NotEmpty(t, Meta1)

	arg := UpdateMetaParams{
		ID: Meta1.ID,
		Title: pgtype.Text{
			String: utils.RandomString(10),
			Valid:  true,
		},
		Overview: pgtype.Text{
			String: utils.RandomString(100),
			Valid:  true,
		},
	}

	Meta2, err := testGojo.UpdateMeta(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Meta2)

	require.Equal(t, Meta1.ID, Meta2.ID)

	require.NotEqual(t, Meta1.Title, Meta2.Title)
	require.NotEqual(t, Meta1.Overview, Meta2.Overview)
	require.WithinDuration(t, Meta1.CreatedAt, Meta2.CreatedAt, time.Second)
}

func TestDeleteMeta(t *testing.T) {
	Meta1 := createRandomMeta(t)
	err := testGojo.DeleteMeta(context.Background(), Meta1.ID)
	require.NoError(t, err)
}
