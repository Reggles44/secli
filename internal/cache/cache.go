package cache

import (
	"errors"
	"os"
	"path"
	"time"
)

var cachedDir string = "/tmp/secli/"

func Expired(fileName string, cache int) bool {
	filePath := path.Join(cachedDir, fileName)
	info, err := os.Stat(filePath)
	// File does not exit
	if err != nil {
		return true
	}

	// Cache can be set to -1 to make it permanent
	if cache < 0 {
		return false
	}

	// Difference between ModTime and Now compared to cacheDuration
	return time.Until(info.ModTime()) >= time.Duration(cache)
}

func Write(fileName string, content *[]byte) error {
	// File exists
	filePath := path.Join(cachedDir, fileName)
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

func Read(fileName string) (*[]byte, error) {
	// File does not exit
	filePath := path.Join(cachedDir, fileName)
	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	// Read File
	data, err := os.ReadFile(filePath)
	return &data, err
}
