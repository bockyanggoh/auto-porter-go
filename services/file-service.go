package services

import (
	"auto-porter-go/models"
	"os"
	"path/filepath"
	"regexp"
)

var videoTypes = []string{".mkv", ".avi", "mp4", ".webM", ".rmvb"}

func searchVideoFiles(folder string) ([]*models.FileData, error) {
	r, _ := regexp.Compile(`^(\.mp4|\.mkv|\.|\.avi|\.webM|\.rmvb)$`)
	var fileList []*models.FileData
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(filepath.Join(path, info.Name()))
		if !info.IsDir() && r.MatchString(ext) {
			fData := models.FileData{
				FullPath:  path,
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

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}