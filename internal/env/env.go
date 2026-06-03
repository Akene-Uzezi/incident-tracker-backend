package env

import (
	"os"
	"strconv"
)


func getEnvString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return  defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}