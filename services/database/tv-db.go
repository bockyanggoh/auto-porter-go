package database

import (
	"auto-porter-go/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func FindAllTvSeries() error {
	db,err := sql.Open("mysql", getConnectionString())

	if err != nil {
		log.Println(err)
		return err
	}

	defer db.Close()

	return nil;
}

func FindTvSeriesByName(name string) (*models.TvSeriesData, error) {
	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	record := &models.TvSeriesData{}
	stmt := "SELECT id, name, year, language, folder_path, episodic FROM tbl_tv_series WHERE name = ?;"

	rows, err := db.Query(stmt, name)
	if err != nil {
		print(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&record.Id, &record.Name, &record.Year, &record.Language, &record.FolderPath, &record.Episodic)
		if err != nil {
			return nil, err
		}

		return record, nil
	}

	return nil , nil
}

func FindTvSeriesById(id string) (*models.TvSeriesData, error) {
	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	record := &models.TvSeriesData{}
	stmt := "SELECT id, name, year, language, folder_path, episodic FROM tbl_tv_series WHERE id = ?;"

	rows, err := db.Query(stmt, id)
	if err != nil {
		print(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&record.Id, &record.Name, &record.Year, &record.Language, &record.FolderPath, &record.Episodic)
		if err != nil {
			return nil, err
		}

		return record, nil
	}

	return nil , nil
}

func InsertRecords(data []models.TvSeriesData) error {
	for _, v := range data {
		err := InsertRecord(v);
		if err != nil {
			log.Printf("Error inserting data for %v. Rollbacking back and terminating session.", v.Id)
			rollbackData(v.Id)
			return err
		}
	}
	log.Printf("Successfully inserted %v records.", len(data))
	return nil
}

func InsertRecord(data models.TvSeriesData) error {
	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		log.Println(err)
		return err
	}

	defer db.Close()

	stmt := "INSERT INTO tbl_tv_series(id, name, year, language, folder_path, episodic) VALUES(?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(stmt, data.Id, data.Name, data.Year, data.Language, data.FolderPath, data.Episodic)

	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Records changed in tv_series\n")

	if len(data.Records) > 0 {
		log.Printf("Adding file records for Series %v with Id %v\n", data.Name, data.Id)
		fileStmt := "INSERT INTO tbl_tv_series_files(series_id, id, file_path, file_name, episode_number, season_number) Values(?, ?, ?, ?, ?, ?)"
		for _,v := range data.Records {
			_, err := db.Exec(fileStmt, v.SeriesId, v.Id, v.FilePath, v.FileName, v.EpisodeNumber, v.SeasonNumber)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

func rollbackData(id string) error {
	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	stmtFiles := "DELETE FROM porter_db.tbl_tv_series_files WHERE series_id = ?"
	stmtSeries := "DELETE FROM porter_db.tbl_tv_series WHERE id = ?"
	_, err = db.Exec(stmtFiles, id)

	if err != nil {
		return err
	}
	_, err = db.Exec(stmtSeries, id)

	if err != nil {
		return err
	}

	return nil
}

