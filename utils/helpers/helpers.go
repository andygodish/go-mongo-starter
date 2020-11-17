package helpers

import "os"

// GetEnv looks up an environment variable. If key is not set, a default value is returned
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
