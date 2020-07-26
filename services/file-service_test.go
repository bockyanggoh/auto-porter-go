package services

import (
	"fmt"
	"os"
	"testing"
)

func TestSearchNewMovieFiles(t *testing.T) {
	loadTestConfig(t)
	path := fmt.Sprintf("%s/%s/%s", os.Getenv("PROJECT_RELATIVE_FOLDER"), os.Getenv("BASE_FOLDER"), "downloads")
	if output, err := searchVideoFiles(path); err != nil {
		t.Errorf("Unexpected error %v\n", err)
	} else {
		if len(output) == 0 {
			t.Error("File Count should not be 0")
		} else {
			t.Logf("Received files: %v", output)
		}
	}
}