// ./utils/directories-creator.go
package utils

import "os"

// CreateDirectories creates all necessary directories for a given path.
func CreateDirectories(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
