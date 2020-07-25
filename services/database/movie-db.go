package database

import (
	"auto-porter-go/models"
	"fmt"
	"log"
)

var movieTable = "tbl_movie"
var movieTableFile = "tbl_movie_files"

func FindMovieByName(name string, fetchChild bool) *models.MovieData {
	if db, err := connectDb(); err == nil {

		defer db.Close()
		stmt := fmt.Sprintf("SELECT * FROM %s WHERE name = ?", movieTable)
		record := models.MovieData{}
		err = db.QueryRow(stmt, name).Scan(
			&record.Id,
			&record.TMDBId,
			&record.Name,
			&record.Year,
			&record.Language,
			&record.FolderPath,
			&record.CreatedTs,
			&record.UpdatedTs,
		)
		if err != nil {
			log.Println(err)
			return nil
		}
		if fetchChild {
			stmt = fmt.Sprintf("SELECT * FROM %s WHERE parent_id == ?", movieTableFile)
			var retFiles []models.MovieFileData
			if rows, err := db.Query(stmt, record.Id); err == nil {

				for rows.Next() {
					file := models.MovieFileData{}
					rows.Scan(&file.Id, &file.ParentId, &file.ParentFolderName, &file.ParentFolderPath, &file.FilePath, &file.FileName, &file.CreatedTs, &file.UpdatedTs)
					retFiles = append(retFiles, file)
				}
				record.Files = retFiles
				return &record
			}

			log.Println(err)
			return nil
		} else {

			return &record
		}
	}
	return nil
}

func InsertMovieRecord(r *models.MovieData) error {
	db, err := connectDb()

	if err != nil {
		return err
	}

	defer db.Close()
	stmt := fmt.Sprintf("INSERT INTO %s(id, tmdb_id, name, year, language, folder_path) values(?, ?, ?, ?, ?, ?)", movieTable)

	_, err = db.Exec(stmt, r.Id, r.TMDBId, r.Name, r.Year, r.Language, r.FolderPath)
	if err != nil {
		return err
	}

	if r.Files != nil && len(r.Files) > 0 {
		stmt := fmt.Sprintf("INSERT INTO tbl_movie_files(id, parent_id, parent_folder_name, parent_folder_path, file_path, file_name) values (?, ?, ?, ?, ?, ?)")
		idList := []string{}
		for _, f := range r.Files {
			_, err = db.Exec(stmt, f.Id, f.ParentId, f.ParentFolderName, f.ParentFolderPath, f.FilePath, f.FileName)
			if err != nil {
				log.Printf("Error Inserting Movie Files! Rolling back these files: %v\n", idList)
				log.Println(err)
				rollbackFiles(idList)
				return err
			}
			idList = append(idList, f.Id)
		}
		return nil
	}
	//TODO: fix this
	return err
}

func rollbackFiles(idList []string) {
	if db, err := connectDb(); err == nil {
		defer db.Close()
		stmt := "DELETE FROM tbl_movie_files WHERE id in ?"
		db.Exec(stmt, idList)
		return
	}

}