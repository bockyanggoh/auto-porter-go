package services

import (
	"auto-porter-go/models"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//func BatchTvSeriesScanJob() {
//}

func BatchRenameFiles() {
	if files := BatchDownloadsScan(); files != nil {

	}

}

func BatchDownloadsScan() []*models.FileData  {
	filePath := fmt.Sprintf("%s/downloads", os.Getenv("BASE_FOLDER"))
	var targetList []*models.FileData
	if fileList, err := searchVideoFiles(filePath); err == nil {
		for _, file := range fileList {
			if strings.Count(file.Name, ".") > 1 {
				file.FormattedName = generateNewFileName(file)
				targetList = append(targetList, file)
				log.Printf("Before: %v, After: %v\n",file.Name, file.FormattedName)
			} else {
				log.Printf("%v does not fulfill requirements for automatic rename.\n", file.Name)
			}
		}
		return targetList
	} else {
		//TODO: error?
		return nil
	}
}

func moveAndRename(info *models.FileData) error {
	path := fmt.Sprintf("%s/apptemp", os.Getenv("BASE_FOLDER"))
	err := os.Mkdir(fmt.Sprintf("%s/%s", path, info.FormattedName), os.ModeDir)
	if err != nil {
		return err
	}
	return nil
}

func generateNewFileName(file *models.FileData) string {
	newFileName := ""
	for _, part := range strings.Split(file.Name, ".") {
		matched, _ := regexp.MatchString(`^\d{4}$`, part)
		if matched {
			newFileName = fmt.Sprintf("%s (%s)", newFileName, part)
			return newFileName
		} else {
			newFileName = fmt.Sprintf("%s %s", newFileName, part)
		}
	}

	return ""
}
//func BatchRenameJob() {
//	files, err := SearchNewFiles()
//
//	if err != nil {
//		fmt.Println("No files found. Exiting Batch Job.")
//	}
//
//	if files != nil {
//		for _, file := range files {
//			log.Printf(file.Name())
//		}
//	}
//}
