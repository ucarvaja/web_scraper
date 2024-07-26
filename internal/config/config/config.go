package config

import (
	"os"
)

type Config struct {
	JSONFilePath string
}

func LoadConfig() Config {
	return Config{
		JSONFilePath: getEnv("JSON_FILE_PATH", "video_info.json"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
