package services

import (
	"io/ioutil"
	"log"
	"os"
)

func SearchNewFiles() ([]os.FileInfo, error) {
	filePath := os.Getenv("DOWNLOADS_PATH")
	files, err := ioutil.ReadDir(filePath)
	log.Printf("Scanning Downloads folder: [%v] for files...\n", filePath)
	if err != nil {
		log.Printf("Error Reading Directory. Scan skipped.\n")
		log.Println(err)
		return nil, err
	}
	if len(files) > 0 {
		return files, nil
	}

	return nil, nil
}
