package cache

import (
	"encoding/json"
	"os"
	"path"

	"github.com/Reggles44/secli/internal/utils/request"
)

type FileCache[T any] struct {
	URL      string
	FileName string
	Duration int
}

func (fc FileCache[T]) Read() (T, error) {
	fileName := path.Base(fc.URL)
	if fc.FileName != "" {
		fileName = fc.FileName
	}

	filePath := filePath(fileName)
	var obj T

	// Read File
	data, err := os.ReadFile(filePath)
	if err == nil {
		err = json.Unmarshal(data, &obj)
		if err == nil {
			return obj, nil
		}
	}

	// Fetch Data
	data, err = request.Do("GET", fc.URL)
	if err == nil {
		err = json.Unmarshal(data, &obj)
		if err == nil {
			go os.WriteFile(filePath, data, 0644)
			return obj, nil
		}
	}

	return *new(T), err
}
