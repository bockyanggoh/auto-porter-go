package services

import (
	"github.com/joho/godotenv"
	"testing"
)

// Test Helper class. Do not write test here.

func loadTestConfig(t *testing.T) {
	if err := godotenv.Load("../test.env"); err != nil {
		t.Error("Error loading .env file")
		return
	}
}