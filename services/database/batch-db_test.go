package database

import (
	"auto-porter-go/models"
	"testing"
	"time"
)

func TestInsertBatchLog(t *testing.T) {
	LoadTestConfig(t)
	startTime := time.Now()
	endTime := startTime.Local().Add(time.Minute * 2)
	record := models.BatchLogData{
		StartExecutionTime: startTime,
		EndExecutionTime:   endTime,
		FilesDetected:      10,
		FoldersDetected:    10,
		Type:               "RENAME_TV",
	}
	err := InsertBatchLog(&record)
	if err != nil {
		t.Error(err)
	}
	db ,err := connectDb()
	defer db.Close()
	count := 0
	db.QueryRow("select count(*) from tbl_batch_log").Scan(&count)

	if count == 0 {
		t.Error("Count should not be 0!")
	}
	t.Cleanup(func() {
		clearDb([]string{"tbl_batch_log"})
	})
}

func TestFindBatchLogByTypeAndLastExecuted(t *testing.T) {
	timeFormat := "StampMilli"
	LoadTestConfig(t)
	startTime1 := time.Now()
	endTime1 := time.Now().Add(time.Minute *1)
	startTime2 := endTime1
	endTime2 := startTime2.Add(time.Minute * 2)
	records := []*models.BatchLogData{}
	records = append(records, &models.BatchLogData{
		Id:                 0,
		StartExecutionTime: startTime1,
		EndExecutionTime:   endTime1,
		FilesDetected:      1,
		FoldersDetected:    1,
		Type:               "RENAME_TV",
	});
	records = append(records, &models.BatchLogData{
		Id:                 0,
		StartExecutionTime: startTime1,
		EndExecutionTime:   endTime1,
		FilesDetected:      1,
		FoldersDetected:    1,
		Type:               "RENAME_MOVIE",
	});
	records = append(records, &models.BatchLogData{
		Id:                 0,
		StartExecutionTime: startTime2,
		EndExecutionTime:   endTime2,
		FilesDetected:      1,
		FoldersDetected:    1,
		Type:               "RENAME_TV",
	});

	for _, f := range records {
		InsertBatchLog(f)
	}

	ret := FindBatchLogByTypeAndLastExecuted("RENAME_TV")

	if ret == nil {
		t.Error("No record retrieved from db!")
	}

	if ret.StartExecutionTime.Format(timeFormat) != startTime2.Format(timeFormat){
		t.Errorf("%v and %v is not equivalent!", startTime2, ret.StartExecutionTime)
	}

	if ret.EndExecutionTime.Format(timeFormat) != endTime2.Format(timeFormat) {
		t.Errorf("%v and %v is not equivalent!", endTime2.Format(timeFormat), ret.EndExecutionTime.Format(timeFormat))
	}

	t.Cleanup(func() {
		clearDb([]string{"tbl_batch_log"})
	})
}


