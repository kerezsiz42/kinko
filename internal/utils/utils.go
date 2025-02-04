package utils

import (
	"fmt"
	"math/rand"
	"os"
)

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateID(length int) string {
	id := make([]byte, length)

	for i := range id {
		id[i] = CHARSET[rand.Intn(len(CHARSET))]
	}

	return string(id)
}

func GetEnv(key string, defaultValue *string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	if defaultValue != nil {
		return *defaultValue
	}

	msg := fmt.Sprintf("Environment variable %s is not set and no default value was provided", key)
	panic(msg)
}
