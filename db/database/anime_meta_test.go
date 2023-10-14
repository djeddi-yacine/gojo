package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnimeMeta(t *testing.T) AnimeMeta {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg := CreateAnimeMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	animeMeta, err := testGojo.CreateAnimeMeta(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeMeta)

	require.Equal(t, arg.AnimeID, animeMeta.AnimeID)
	require.Equal(t, arg.MetaID, animeMeta.MetaID)
	require.Equal(t, arg.LanguageID, animeMeta.LanguageID)

	require.NotZero(t, animeMeta.ID)

	return animeMeta
}

func TestCreateAnimeMeta(t *testing.T) {
	createRandomAnimeMeta(t)
}

func TestGetAnimeMeta(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta1, err := testGojo.CreateAnimeMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta1)

	arg2 := GetAnimeMetaParams{
		AnimeID:    a.ID,
		LanguageID: l.ID,
	}

	mID, err := testGojo.GetAnimeMeta(context.Background(), arg2)
	require.NoError(t, err)
	require.NotZero(t, mID)
	require.Equal(t, m.ID, mID)
}

func TestUpdateAnimeMeta(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta1, err := testGojo.CreateAnimeMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta1)

	nm := createRandomMeta(t)
	require.NotEmpty(t, nm)

	arg2 := UpdateAnimeMetaParams{
		AnimeID:    a.ID,
		LanguageID: l.ID,
		MetaID:     nm.ID,
	}

	Meta2, err := testGojo.UpdateAnimeMeta(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, Meta2)

	require.Equal(t, Meta1.ID, Meta2.ID)
	require.Equal(t, Meta1.AnimeID, Meta2.AnimeID)
	require.Equal(t, Meta1.LanguageID, Meta2.LanguageID)
	require.NotEqual(t, Meta1.MetaID, Meta2.MetaID)
}

func TestDeleteAnimeMeta(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta, err := testGojo.CreateAnimeMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta)

	arg2 := DeleteAnimeMetaParams{
		AnimeID:    Meta.AnimeID,
		LanguageID: Meta.LanguageID,
	}

	err = testGojo.DeleteAnimeMeta(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeMetas(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		m := createRandomMeta(t)
		require.NotEmpty(t, m)
		l := createRandomLanguage(t)
		require.NotEmpty(t, l)
		arg := CreateAnimeMetaParams{
			AnimeID:    a.ID,
			LanguageID: l.ID,
			MetaID:     m.ID,
		}
		testGojo.CreateAnimeMeta(context.Background(), arg)
	}

	arg := ListAnimeMetasParams{
		AnimeID: a.ID,
		Limit:   3,
		Offset:  0,
	}

	Metas, err := testGojo.ListAnimeMetas(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, Metas)
	require.Len(t, Metas, 3)
}
