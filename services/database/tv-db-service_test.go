package database

import (
	"auto-porter-go/models"
	"database/sql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"testing"
)

func TestInsertRecordSuccess(t *testing.T) {

	if err := godotenv.Load("../../.env.test"); err != nil {
		t.Error("Error loading .env file")
		return
	}
	seriesId := uuid.New().String()
	fileId1 := uuid.New().String()
	fileId2 := uuid.New().String();
	r1 := models.TvRecordData{
		Id: fileId1,
		SeriesId: seriesId,
		FilePath: "/test/path",
		FileName: "test",
		EpisodeNumber: 1,
		SeasonNumber: 1,
	}
	r2 := models.TvRecordData{
		Id: fileId2,
		SeriesId: seriesId,
		FilePath: "/test/path/2",
		FileName: "testPath",
		EpisodeNumber: 1,
		SeasonNumber: 1,
	}

	data := []models.TvRecordData {r1, r2}
	record := models.TvSeriesData{
		Id: seriesId,
		Name: "test1",
		Year: 1965,
		Language: "English",
		FolderPath: "/test",
		Episodic: true,
		Records: data,
	}

	err := InsertRecord(record)

	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		cleanUp()
	})
}

func TestInsertRecordsSuccess(t *testing.T) {

	if err := godotenv.Load("../../.env.test"); err != nil {
		t.Error("Error loading .env file")
		return
	}
	seriesId1 := uuid.New().String()
	seriesId2 := uuid.New().String()
	data := []models.TvRecordData {}
	record1 := models.TvSeriesData{
		Id: seriesId1,
		Name: "test1",
		Year: 1965,
		Language: "English",
		FolderPath: "/test",
		Episodic: true,
		Records: data,
	}
	record2 := models.TvSeriesData{
		Id: seriesId2,
		Name: "test2",
		Year: 1965,
		Language: "English",
		FolderPath: "/test",
		Episodic: true,
		Records: data,
	}

	err := InsertRecords([]models.TvSeriesData{record1, record2})

	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		cleanUp()
	})
}

func cleanUp() {

	db, _ := sql.Open("mysql", getConnectionString())
	defer db.Close()

	db.Exec("truncate tbl_tv_series_files")
	db.Exec("truncate tbl_tv_series")


}
