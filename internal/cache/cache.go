package cache

import (
	"errors"
	"os"
	"path"
	"time"
)

var cachedDir string = "/tmp/secli/"

func CacheFilePath(fileName string) string {
	return path.Join(cachedDir, fileName)
}

func Write(fileName string, content *[]byte) error {
	// File exists
	filePath := CacheFilePath(fileName)
	_, err := os.Stat(filePath)
	if !errors.Is(err, os.ErrNotExist) {
		return nil
	}

	// Create File
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write to File
	_, err = file.Write(*content)
	if err != nil {
		return err
	}

	return nil
}

func Read(fileName string, cacheDuration int) (*[]byte, error) {
	// File does not exit
	filePath := CacheFilePath(fileName)
	fileInfo, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	if time.Until(fileInfo.ModTime()) >= time.Duration(cacheDuration) {
		os.Remove(filePath)
		return nil, errors.New("cache expired")
	}

	// Read File
	data, err := os.ReadFile(filePath)
	return &data, err
}
