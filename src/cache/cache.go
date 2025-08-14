package secapi

import (
	"errors"
	"net/url"
	"os"
	"path"
	"time"
)

var cachedDir string = "/tmp/secli/"

func cacheExpired(fileName string, cacheDuration int64) bool {
	filePath := path.Join(cachedDir, fileName)
	info, err := os.Stat(filePath)
	// File does not exit
	if err != nil {
		return true
	}

	// Difference between ModTime and Now compared to cacheDuration
	return info.ModTime().Sub(time.Now()) >= time.Duration(cacheDuration)
}

func writeToCache(fileName string, content *[]byte) error {
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

func readFromCache(fileName string) (*[]byte, error) {
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

