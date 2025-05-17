package utility

import (
	"log"
	"os"
	"strconv"
)

// GetStringValueFromEnv helps to get the Key from the os Environment
func GetStringValueFromEnv(key, fallback string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	log.Printf("key=%s is not found in os. setting to fallback value", key)
	return fallback
}

// GetNumberValueFromEnv helps to get the Key from the os Environment
func GetNumberValueFromEnv(key string, fallback int) int {
	if value, found := os.LookupEnv(key); found {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Printf("error parsing environment variable %s: %v", key, err)
			return fallback
		}
		return int(intValue)
	}
	log.Printf("key=%s is not found in os. setting to fallback value", key)
	return fallback
}
