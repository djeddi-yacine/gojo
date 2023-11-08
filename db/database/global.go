package db

type AnimeMetaTxParam struct {
	LanguageID int32
	CreateMetaParams
}

type AnimeMetaTxResult struct {
	Language Language
	Meta     Meta
}

