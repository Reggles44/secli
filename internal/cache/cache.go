package cache

import (
	"path"
)

type Cache interface {
	Read() (*[]byte, error)
	Write(*[]byte) error
}

var cachedDir string = "/tmp/secli/"

func filePath(fileName string) string {
	return path.Join(cachedDir, fileName)
}
