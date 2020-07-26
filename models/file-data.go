package models

import "time"

type FileData struct {
	FullPath string
	Name string
	FormattedName string
	Extension string
	Folder string
	ModTime time.Time
	Size int64
}