package utils

import "os"

func CreateDirectories(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
