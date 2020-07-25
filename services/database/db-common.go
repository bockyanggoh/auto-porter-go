package database

import (
	"database/sql"
	"fmt"
	"os"
)

func getConnectionString() string {
	conString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_URL"),
		os.Getenv("DB_NAME"))

	return conString
}

func connectDb() (*sql.DB, error){
	db,err := sql.Open("mysql", getConnectionString())
	if err != nil {
		return nil, err
	}
	return db, nil
}
