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
		for _, f := range files {
			moveAndRename(f)
		}
	}

}

func BatchDownloadsScan() []*models.FileData  {
	filePath := fmt.Sprintf("%s/%s/downloads", os.Getenv("PROJECT_RELATIVE_FOLDER"), os.Getenv("BASE_FOLDER"))
	log.Printf("Scanning folder for downloads: %v\n", filePath)
	var targetList []*models.FileData
	if fileList, err := searchVideoFiles(filePath); err == nil {
		for _, file := range fileList {
			if strings.Count(file.Name, ".") > 1 {
				file.FormattedName = generateNewFileName(file)
				targetList = append(targetList, file)
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
	log.Println("Renaming files...")
	path := fmt.Sprintf("%s/%s/apptemp", os.Getenv("PROJECT_RELATIVE_FOLDER"), os.Getenv("BASE_FOLDER"))

	oldPath := fmt.Sprintf("%s", info.FullPath)
	newPath := fmt.Sprintf("%s/%s%s", path, info.FormattedName, info.Extension)

	log.Printf("oldPath: %v", oldPath)
	log.Printf("newPath: %v", newPath)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Println(err)
		log.Println("Failed to move files.")
		return err
	}
	moviePath := fmt.Sprintf("%s/%s/movies/%s", os.Getenv("PROJECT_RELATIVE_FOLDER"), os.Getenv("BASE_FOLDER"), info.FormattedName)
	log.Println(moviePath)
	if _, err := os.Stat(moviePath); os.IsNotExist(err) {
		//Path does not exist, lets continue.
		if err = os.Mkdir(moviePath, 0777); err != nil {
			log.Println(err)
			return nil
		}
		finalLoc := fmt.Sprintf("%s/%s%s", moviePath, info.FormattedName, info.Extension)
		log.Println(finalLoc)
		err = os.Rename(newPath, finalLoc)
		if err != nil {
			log.Println(err)
			log.Println("Failed to move file. Moving file to others for manual intervention.")
			moveFailureToOthers()
		} else {
			log.Println("Completed rename job. Exiting.")
			return nil
		}
	}

	return nil
}

func moveFailureToOthers() {

}

func deleteShellFolders() {
}

func generateNewFileName(file *models.FileData) string {
	newFileName := ""
	for _, part := range strings.Split(file.Name, ".") {
		matched, _ := regexp.MatchString(`^\d{4}$`, part)
		if matched {
			newFileName = fmt.Sprintf("%s (%s)", newFileName, part)
			newFileName = strings.Trim(newFileName, " ")
			return newFileName
		} else {
			newFileName = fmt.Sprintf("%s %s", newFileName, part)
		}
	}
	return ""
}