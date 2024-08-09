package utils

import (
	"os"
	"path/filepath"
)

func GetFileName(filePath string) string {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}

	return fileInfo.Name()
}

func GetDirectoryPath(filePath string) string {
	dirPath := filepath.Dir(filePath)
	return dirPath
}
