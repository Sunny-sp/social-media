package config

import (
	"os"
	"strconv"
)

func getEnvString(key string, fallback string) string {

	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}
	return val
}

func getEnvInt(key string, fallback int) int {

	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

// func getEnvBool(key string, fallback bool) bool {

// 	val, ok := os.LookupEnv(key)

// 	if !ok {
// 		return fallback
// 	}

// 	valAsInt, err := strconv.ParseBool(val)
// 	if err != nil {
// 		return fallback
// 	}

// 	return valAsInt
// }
