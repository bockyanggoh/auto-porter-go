package services

import (
	"github.com/joho/godotenv"
	"testing"
)

func TestFileSearch(t *testing.T) {
	loadConfig(t)
	SearchNewFiles()
}

func loadConfig(t *testing.T) {
	if err := godotenv.Load("../test.env"); err != nil {
		t.Error("Error loading .env file")
		return
	}
}