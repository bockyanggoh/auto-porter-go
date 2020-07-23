package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func BatchTvSeriesScanJob() {
	fmt.Print(os.Getenv("TV_FOLDER_PATH"))
	filepath.Walk(os.Getenv("TV_FOLDER_PATH"), func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) == ".mkv" || filepath.Ext(path) == ".mp4" || filepath.Ext(path) == ".rmvb" {
			fmt.Printf("%s", info.Name())
		}

		return nil
	})
}

func BatchRenameJob() {
	files, err := SearchNewFiles()

	if err != nil {
		fmt.Println("No files found. Exiting Batch Job.")
	}

	if files != nil {
		for _, file := range files {
			log.Printf(file.Name())
		}
	}
}



