package db

import (
	"context"
	"testing"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/stretchr/testify/require"
)

func createRandomLanguage(t *testing.T) Language {
	arg := CreateLanguageParams{
		LanguageName: utils.RandomString(12),
		LanguageCode: utils.RandomString(6),
	}
	language, err := testGojo.CreateLanguage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, language)
	return language
}

func createCustomLanguage(t *testing.T, name, code string) Language {
	arg := CreateLanguageParams{
		LanguageName: name,
		LanguageCode: code,
	}
	language, err := testGojo.CreateLanguage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, language)
	return language
}

func TestCreateLanguage(t *testing.T) {
	createRandomLanguage(t)
}

func TestGetLanguage(t *testing.T) {
	/// English
	LanguageEN1 := createCustomLanguage(t, "English", "EN")
	LanguageEN2, err := testGojo.GetLanguage(context.Background(), LanguageEN1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, LanguageEN2)

	require.Equal(t, LanguageEN1.ID, LanguageEN2.ID)
	require.Equal(t, LanguageEN1.LanguageName, LanguageEN2.LanguageName)
	require.WithinDuration(t, LanguageEN1.CreatedAt, LanguageEN2.CreatedAt, time.Second)
	err = testGojo.DeleteLanguage(context.Background(), LanguageEN1.ID)
	require.NoError(t, err)

	///Arabic
	LanguageAR1 := createCustomLanguage(t, "Arabic", "AR")
	LanguageAR2, err := testGojo.GetLanguage(context.Background(), LanguageAR1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, LanguageAR2)

	require.Equal(t, LanguageAR1.ID, LanguageAR2.ID)
	require.Equal(t, LanguageAR1.LanguageName, LanguageAR2.LanguageName)
	require.WithinDuration(t, LanguageAR1.CreatedAt, LanguageAR2.CreatedAt, time.Second)
	err = testGojo.DeleteLanguage(context.Background(), LanguageAR1.ID)
	require.NoError(t, err)

	///Japanese
	LanguageJP1 := createCustomLanguage(t, "Japanese", "JP")
	LanguageJP2, err := testGojo.GetLanguage(context.Background(), LanguageJP1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, LanguageJP2)

	require.Equal(t, LanguageJP1.ID, LanguageJP2.ID)
	require.Equal(t, LanguageJP1.LanguageName, LanguageJP2.LanguageName)
	require.WithinDuration(t, LanguageJP1.CreatedAt, LanguageJP2.CreatedAt, time.Second)
	err = testGojo.DeleteLanguage(context.Background(), LanguageJP1.ID)
	require.NoError(t, err)
}

func TestUpdateLanguage(t *testing.T) {
	Language1 := createRandomLanguage(t)
	require.NotEmpty(t, Language1)

	arg := UpdateLanguageParams{
		ID:           Language1.ID,
		LanguageName: utils.RandomString(15),
		LanguageCode: utils.RandomString(5),
	}

	Language2, err := testGojo.UpdateLanguage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Language2)

	require.Equal(t, Language1.ID, Language2.ID)
	require.NotEqual(t, Language1.LanguageName, Language2.LanguageName)
	require.WithinDuration(t, Language1.CreatedAt, Language2.CreatedAt, time.Second)
}

func TestDeleteLanguage(t *testing.T) {
	Language1 := createRandomLanguage(t)
	err := testGojo.DeleteLanguage(context.Background(), Language1.ID)
	require.NoError(t, err)
}

func TestListLanguages(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomLanguage(t)
	}

	arg := ListLanguagesParams{
		Limit:  3,
		Offset: 0,
	}

	Languages, err := testGojo.ListLanguages(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Languages, 3)
	require.NotNil(t, Languages)

	for _, s := range Languages {
		require.NotNil(t, s.ID)
		require.NotNil(t, s.LanguageName)
		require.NotNil(t, s.CreatedAt)
	}

}
