package database

import (
	"auto-porter-go/models"
	"log"
)

func InsertBatchLog(data *models.BatchLogData) error {
	if db, err := connectDb(); err == nil {
		stmt := "INSERT INTO tbl_batch_log(start_execution_time, end_execution_time, files_detected, folders_detected, job_type) VALUES (?, ?, ?, ?, ?);"

		if _, err := db.Exec(stmt, data.StartExecutionTime, data.EndExecutionTime, data.FilesDetected, data.FoldersDetected, data.Type); err == nil {
			return nil
		} else {
			return err;
		}
	} else {
		return err
	}
}

func FindBatchLogByTypeAndLastExecuted(jobType string) *models.BatchLogData {
	if db, err := connectDb(); err == nil {
		stmt := "SELECT * FROM tbl_batch_log WHERE job_type = ? order by end_execution_time desc limit 1;"
		record := models.BatchLogData{}
		err := db.QueryRow(stmt, jobType).Scan(&record.Id, &record.StartExecutionTime, &record.EndExecutionTime, &record.FilesDetected, &record.FoldersDetected, &record.Type)
		if err != nil {
			log.Println(err)
			return nil
		}
		return &record
	} else {
		return nil
	}
}
