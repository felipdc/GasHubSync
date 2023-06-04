package utils

import (
	"log"
	"os"
)

// ReadToken reads a token from a file and returns it as a string
func ReadToken(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return string(data)
}
