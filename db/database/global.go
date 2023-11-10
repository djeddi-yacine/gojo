package db

type AnimeMetaTxParam struct {
	LanguageID int32
	CreateMetaParams
}

type AnimeMetaTxResult struct {
	LanguageID int32
	Meta       Meta
}
