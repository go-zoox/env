package env

import "os"

// Get gets the value of the given key from system environment.
func Get(key string, defaultValue ...string) string {
	if len(defaultValue) > 1 {
		panic("Too many arguments supplied")
	}

	value := os.Getenv(key)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return value
}
