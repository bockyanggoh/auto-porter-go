package models

import "time"

type MovieData struct {
	Id, TMDBId, Name, Language, FolderPath string
	Year int
	CreatedTs, UpdatedTs time.Time
	Files []MovieFileData
}

type MovieFileData struct {
	Id, ParentId, ParentFolderName, ParentFolderPath, FilePath, FileName string
	CreatedTs, UpdatedTs time.Time
}