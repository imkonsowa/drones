package env

import (
	"os"
)

func String(key string, fallback ...string) string {
	val := os.Getenv(key)
	if len(val) > 0 {
		return val
	}

	if len(fallback) > 0 {
		return fallback[0]
	}

	return ""
}
