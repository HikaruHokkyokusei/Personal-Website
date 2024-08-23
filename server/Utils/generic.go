package Utils

import (
	"os"
)

func GetDefaultEnv(key string, defaultValue string) string {
	if value, isPresent := os.LookupEnv(key); isPresent {
		return value
	} else {
		return defaultValue
	}
}
