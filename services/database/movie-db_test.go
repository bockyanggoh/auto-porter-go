package database

import (
	"auto-porter-go/models"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"
)



func TestInsertMovieRecord_NoChild_Single(t *testing.T) {
	LoadTestConfig(t)

	record := models.MovieData{
		Id:         uuid.New().String(),
		TMDBId:     "1234",
		Name:       "test1",
		Language:   "english",
		FolderPath: "/test/path",
		Year:       1990,
		CreatedTs:  time.Time{},
		UpdatedTs:  time.Time{},
		Files:      nil,
	}
	err :=InsertMovieRecord(&record)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}

	db, err := connectDb()

	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}

	defer db.Close()

	count := 0
	db.QueryRow("SELECT count(id) FROM tbl_movie WHERE id = ?", record.Id).Scan(&count)
	if count == 0 {
		t.Error("Count should be 1 after insertion!")
	}
	t.Cleanup(func() {
		clearDb([]string{"tbl_movie", "tbl_movie_files"})
	})

}

func TestInsertMovieRecord_WithChild_Single(t *testing.T) {
	LoadTestConfig(t)

	record := models.MovieData{
		Id:         uuid.New().String(),
		TMDBId:     "1234",
		Name:       "test1",
		Language:   "english",
		FolderPath: "/test/path",
		Year:       1990,
		CreatedTs:  time.Time{},
		UpdatedTs:  time.Time{},
		Files:      nil,
	}

	file := models.MovieFileData{
		Id:               uuid.New().String(),
		ParentId:         record.Id,
		ParentFolderName: "test",
		ParentFolderPath: "test",
		FilePath:         "test",
		FileName:         "test",
		CreatedTs:        time.Time{},
		UpdatedTs:        time.Time{},
	}
	record.Files = append(record.Files, file)
	err :=InsertMovieRecord(&record)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}

	db, err := connectDb()

	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}

	defer db.Close()

	movieCount := 0
	movieFileCount := 0
	db.QueryRow("SELECT count(id) FROM tbl_movie WHERE id = ?", record.Id).Scan(&movieCount)
	db.QueryRow("SELECT count(id) FROM tbl_movie_files WHERE id = ?", record.Files[0].Id).Scan(&movieFileCount)
	if movieCount == 0 || movieFileCount == 0 {
		t.Error("Count should be 1 after insertion!")
	}
	t.Cleanup(func() {
		clearDb([]string{"tbl_movie", "tbl_movie_files"})
	})
}



