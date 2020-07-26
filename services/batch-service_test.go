package services

import "testing"

func TestBatchDownloadsScan(t *testing.T) {
	loadTestConfig(t)
	output := BatchDownloadsScan()

	if output != nil {
		if len(output) == 0 {
			t.Error("Output should not be empty!")
		}

	} else {
		t.Error("Output is nil!")
	}
}

func TestBatchRenameFiles(t *testing.T) {
	loadTestConfig(t)
	BatchRenameFiles()
}