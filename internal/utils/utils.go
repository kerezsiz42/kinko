package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

const (
	Lowercase          = "abcdefghijklmnopqrstuvwxyz"
	Uppercase          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers            = "0123456789"
	Special            = "!@#$%^&*()-_+=[]{}|;:'\",.<>?/`~"
	Alphanum           = Lowercase + Uppercase + Numbers
	LowercaseAndNumber = Lowercase + Numbers
)

func GenerateString(charset string, length int) string {
	id := make([]byte, length)

	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
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

func DeepCopy[T any](src T) (T, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return *new(T), fmt.Errorf("error marshaling to JSON: %v", err)
	}

	var copy T
	if err := json.Unmarshal(data, &copy); err != nil {
		return *new(T), fmt.Errorf("error unmarshaling from JSON: %v", err)
	}

	return copy, nil
}

func AreSlicesIdentical(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Ptr[T any](v T) *T {
	return &v
}
