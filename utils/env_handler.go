package utils

import "os"

func GetEnv(key string, alternative string) string {

	val := os.Getenv(key)

	if val == "" {
		return alternative
	}

	return val
}
