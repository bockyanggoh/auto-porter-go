package models

type TvSeriesData struct {
	Id, Name, Language, FolderPath string
	Year int
	Episodic bool
	Records []TvRecordData
}


type TvRecordData struct {
	Id, SeriesId, FilePath, FileName string
	EpisodeNumber, SeasonNumber int
}

