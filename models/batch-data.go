package models

import "time"

type BatchLogData struct {
	Id int
	StartExecutionTime, EndExecutionTime time.Time
	FilesDetected, FoldersDetected int
	Type string
}
