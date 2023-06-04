package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	PollingIntervalSeconds int `json:"polling_interval_seconds"`
}

// ReadConfig reads a configuration file and returns a Config struct
func ReadConfig(filePath string) Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	return config
}
