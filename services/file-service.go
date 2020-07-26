package services

import (
	"auto-porter-go/models"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var videoTypes = []string{".mkv", ".avi", "mp4", ".webM", ".rmvb"}

func searchVideoFiles(folder string) ([]*models.FileData, error) {

	r, _ := regexp.Compile(`^(\.mp4|\.mkv|\.|\.avi|\.webM|\.rmvb)$`)
	var fileList []*models.FileData
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(filepath.Join(path, info.Name()))
		if !info.IsDir() && r.MatchString(ext) {
			fData := models.FileData{
				FullPath:  fmt.Sprintf("%s/%s", path, info.Name()),
				Name:      info.Name(),
				Extension: ext,
				Folder:    path,
				ModTime:   info.ModTime(),
				Size:      info.Size(),
			}
			fileList = append(fileList, &fData)
		}
		return nil
	})

	return fileList, nil
}

func renameTVFile(completion chan bool) {

}

func renameMovieFile(completion chan bool) {

}

func renameAnimeFile(completion chan bool) {

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}