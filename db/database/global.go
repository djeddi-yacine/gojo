package db

type AnimeMetaTxParam struct {
	LanguageID int32
	CreateMetaParams
}

type AnimeMetaTxResult struct {
	Language Language
	Meta     Meta
}

func checkLanguage(languages []Language, n int32) bool {
	for _, language := range languages {
		if language.ID == n {
			return true
		}
	}

	return false
}
