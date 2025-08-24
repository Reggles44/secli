package cache

import (
	"os"
	"path"
	"strings"

	"github.com/Reggles44/secli/internal/utils/request"
	"github.com/Reggles44/secli/internal/utils/zip"
)

type ZipCache struct {
	URL      string
	FileName string
}

func fileToDir(path string) string {
	return strings.Replace(path, ".zip", "", 1)
}

func (zc ZipCache) Read() (map[string]*[]byte, error) {
	data, err := zc.load()
	if err != nil {
		raw, err := request.Do("GET", zc.URL)
		if err != nil {
			return nil, err
		}

		data, err = zip.Unzip(&raw)
		if err != nil {
			return nil, err
		}
	}

	go zc.write(data)

	return data, nil
}

func (zc ZipCache) load() (map[string]*[]byte, error) {
	data := make(map[string]*[]byte)

	fileName := path.Base(zc.URL)
	if zc.FileName != "" {
		fileName = zc.FileName
	}

	zipPath := filePath(fileName)
	dir := fileToDir(zipPath)

	paths, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range paths {
		path := path.Join(dir, entry.Name())

		d, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		data[entry.Name()] = &d
	}

	return data, nil
}

func (zc ZipCache) write(data map[string]*[]byte) error {
	fileName := path.Base(zc.URL)
	if zc.FileName != "" {
		fileName = zc.FileName
	}

	zipPath := filePath(fileName)
	dir := fileToDir(zipPath)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	for name, d := range data {
		err := os.WriteFile(path.Join(dir, name), *d, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
