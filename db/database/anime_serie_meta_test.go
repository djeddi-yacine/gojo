package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnimeSerieMeta(t *testing.T) AnimeSerieMeta {
	a := createRandomAnimeSerie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg := CreateAnimeSerieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	animeSerieMeta, err := testGojo.CreateAnimeSerieMeta(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeSerieMeta)

	require.Equal(t, arg.AnimeID, animeSerieMeta.AnimeID)
	require.Equal(t, arg.MetaID, animeSerieMeta.MetaID)
	require.Equal(t, arg.LanguageID, animeSerieMeta.LanguageID)

	require.NotZero(t, animeSerieMeta.ID)

	return animeSerieMeta
}

func TestCreateAnimeSerieMeta(t *testing.T) {
	createRandomAnimeSerieMeta(t)
}

func TestGetAnimeSerieMeta(t *testing.T) {
	a := createRandomAnimeSerie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeSerieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta1, err := testGojo.CreateAnimeSerieMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta1)

	arg2 := GetAnimeSerieMetaParams{
		AnimeID:    a.ID,
		LanguageID: l.ID,
	}

	mID, err := testGojo.GetAnimeSerieMeta(context.Background(), arg2)
	require.NoError(t, err)
	require.NotZero(t, mID)
	require.Equal(t, m.ID, mID)
}

func TestUpdateAnimeSerieMeta(t *testing.T) {
	a := createRandomAnimeSerie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeSerieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta1, err := testGojo.CreateAnimeSerieMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta1)

	nm := createRandomMeta(t)
	require.NotEmpty(t, nm)

	arg2 := UpdateAnimeSerieMetaParams{
		AnimeID:    a.ID,
		LanguageID: l.ID,
		MetaID:     nm.ID,
	}

	Meta2, err := testGojo.UpdateAnimeSerieMeta(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, Meta2)

	require.Equal(t, Meta1.ID, Meta2.ID)
	require.Equal(t, Meta1.AnimeID, Meta2.AnimeID)
	require.Equal(t, Meta1.LanguageID, Meta2.LanguageID)
	require.NotEqual(t, Meta1.MetaID, Meta2.MetaID)
}

func TestDeleteAnimeSerieMeta(t *testing.T) {
	a := createRandomAnimeSerie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeSerieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta, err := testGojo.CreateAnimeSerieMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta)

	arg2 := DeleteAnimeSerieMetaParams{
		AnimeID:    Meta.AnimeID,
		LanguageID: Meta.LanguageID,
	}

	err = testGojo.DeleteAnimeSerieMeta(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeSerieMetas(t *testing.T) {
	a := createRandomAnimeSerie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		m := createRandomMeta(t)
		require.NotEmpty(t, m)
		l := createRandomLanguage(t)
		require.NotEmpty(t, l)
		arg := CreateAnimeSerieMetaParams{
			AnimeID:    a.ID,
			LanguageID: l.ID,
			MetaID:     m.ID,
		}
		testGojo.CreateAnimeSerieMeta(context.Background(), arg)
	}

	Metas, err := testGojo.ListAnimeSerieMetas(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotNil(t, Metas)
}
