package database

import (
	"fmt"
	"log"
)

// Test Helper class. Do not write test here.

func clearDb(list []string) {
	db, _ := connectDb()
	defer db.Close()
	for _, l := range list {
		stmt := fmt.Sprintf("DELETE FROM %s", l)
		db.Exec(stmt)
	}

	log.Printf("Cleared all data in the following tables: %v\n", list)
}
