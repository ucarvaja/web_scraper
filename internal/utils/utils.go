package utils

import (
	"os"
)

// GetEnv gets the environment variable or returns a default value if not set
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
