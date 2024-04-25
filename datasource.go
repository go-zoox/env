package env

import (
	"fmt"
	"os"
)

// DataSource is a data source that loads data from the environment.
type DataSource struct {
}

// Get returns the value of the given key.
func (DataSource) Get(path, key string) any {
	fmt.Println("path", path)
	fmt.Println("key", key)
	// value := os.Getenv(key)
	// if value != "" {
	// 	return value
	// }

	// return nil

	// @TODO: This is a hack to get the last part of the key
	// ks := strings.Split(path, ".")
	// key = ks[len(ks)-1]

	return os.Getenv(key)
}
