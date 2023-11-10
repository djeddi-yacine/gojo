package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnimeMovieMeta(t *testing.T) AnimeMovieMeta {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg := CreateAnimeMovieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	animeMovieMeta, err := testGojo.CreateAnimeMovieMeta(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeMovieMeta)

	require.Equal(t, arg.AnimeID, animeMovieMeta.AnimeID)
	require.Equal(t, arg.MetaID, animeMovieMeta.MetaID)
	require.Equal(t, arg.LanguageID, animeMovieMeta.LanguageID)

	require.NotZero(t, animeMovieMeta.ID)

	return animeMovieMeta
}

func TestCreateAnimeMovieMeta(t *testing.T) {
	createRandomAnimeMovieMeta(t)
}

func TestGetAnimeMovieMeta(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeMovieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta1, err := testGojo.CreateAnimeMovieMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta1)

	arg2 := GetAnimeMovieMetaParams{
		AnimeID:    a.ID,
		LanguageID: l.ID,
	}

	mID, err := testGojo.GetAnimeMovieMeta(context.Background(), arg2)
	require.NoError(t, err)
	require.NotZero(t, mID)
	require.Equal(t, m.ID, mID)
}

func TestUpdateAnimeMovieMeta(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeMovieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta1, err := testGojo.CreateAnimeMovieMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta1)

	nm := createRandomMeta(t)
	require.NotEmpty(t, nm)

	arg2 := UpdateAnimeMovieMetaParams{
		AnimeID:    a.ID,
		LanguageID: l.ID,
		MetaID:     nm.ID,
	}

	Meta2, err := testGojo.UpdateAnimeMovieMeta(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, Meta2)

	require.Equal(t, Meta1.ID, Meta2.ID)
	require.Equal(t, Meta1.AnimeID, Meta2.AnimeID)
	require.Equal(t, Meta1.LanguageID, Meta2.LanguageID)
	require.NotEqual(t, Meta1.MetaID, Meta2.MetaID)
}

func TestDeleteAnimeMovieMeta(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)
	m := createRandomMeta(t)
	require.NotEmpty(t, m)
	l := createRandomLanguage(t)
	require.NotEmpty(t, l)
	arg1 := CreateAnimeMovieMetaParams{
		AnimeID:    a.ID,
		MetaID:     m.ID,
		LanguageID: l.ID,
	}

	Meta, err := testGojo.CreateAnimeMovieMeta(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, Meta)

	arg2 := DeleteAnimeMovieMetaParams{
		AnimeID:    Meta.AnimeID,
		LanguageID: Meta.LanguageID,
	}

	err = testGojo.DeleteAnimeMovieMeta(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeMovieMetas(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		m := createRandomMeta(t)
		require.NotEmpty(t, m)
		l := createRandomLanguage(t)
		require.NotEmpty(t, l)
		arg := CreateAnimeMovieMetaParams{
			AnimeID:    a.ID,
			LanguageID: l.ID,
			MetaID:     m.ID,
		}
		testGojo.CreateAnimeMovieMeta(context.Background(), arg)
	}

	Metas, err := testGojo.ListAnimeMovieMetas(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotNil(t, Metas)
}
