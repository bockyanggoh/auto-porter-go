package models

type TvSeriesData struct {
	Id string
	Name string
	Year int
	Language string
	FolderPath string
	Episodic bool
	Records []TvRecordData
}


type TvRecordData struct {
	Id string
	SeriesId string
	FilePath string
	FileName string
	EpisodeNumber int
	SeasonNumber int
}

