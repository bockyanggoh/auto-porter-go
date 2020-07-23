package database

import (
	"auto-porter-go/models"
	"database/sql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestInsertRecordSuccess(t *testing.T) {
	loadConfig(t)

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

	loadConfig(t)
	if err := godotenv.Load("../../test.env"); err != nil {
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

func TestFindTvSeriesByName_NoResults(t *testing.T) {

	loadConfig(t)
	res , err := FindTvSeriesByName("test")
	if err == nil && res == nil {
		log.Print(err)
		return
	}

	t.Error("Expected error to be thrown for no results!")

	t.Cleanup(func() {
		cleanUp()
	})
}

func TestFindTvSeriesByName_WithSingleResult(t *testing.T) {
	name :="peacan pie"
	loadConfig(t)
	recordId := uuid.New().String()
	InsertRecord(models.TvSeriesData{
		Id:         recordId,
		Name:       name,
		Year:       1965,
		Language:   "English",
		FolderPath: "/test",
		Episodic:   true,
	})
	res, err := FindTvSeriesByName(name)

	if err != nil {
		t.Error(err)
	}

	if res.Name != name {
		t.Error("Response does not match input!")
	}

	t.Cleanup(func() {
		cleanUp()
	})
}

func TestFindTvSeriesById_Fail(t *testing.T) {
	id1 := uuid.New().String()
	loadConfig(t)

	res, err := FindTvSeriesById(id1)

	if err != nil && res != nil {
		t.Error(err)
	}

}

func TestFindTvSeriesById_SingleResult(t *testing.T) {
	id1 := uuid.New().String()
	loadConfig(t)
	InsertRecord(models.TvSeriesData{
		Id:         id1,
		Name:       "tester",
		Year:       1965,
		Language:   "English",
		FolderPath: "/test",
		Episodic:   true,
	})
	res, err := FindTvSeriesById(id1)

	if err != nil {
		t.Error(err)
	}

	if res.Id != id1 {
		t.Error("Response does not match input!")
	}

	t.Cleanup(func() {
		cleanUp()
	})
}

func cleanUp() {

	db, _ := sql.Open("mysql", getConnectionString())
	defer db.Close()

	db.Exec("delete from tbl_tv_series_files")
	db.Exec("delete from tbl_tv_series")


}

func loadConfig(t *testing.T) {
	if err := godotenv.Load("../../test.env"); err != nil {
		t.Error("Error loading .env file")
		return
	}
}