package request

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
)

func GetZip(method string, url string, cacheDuration int) (map[string]*[]byte, error) {
	fmt.Println(url)

	zipData, err := Get(method, url, cacheDuration)
	if err != nil {
		return nil, err
	}

	br := bytes.NewReader(*zipData)
	reader, err := zip.NewReader(br, int64(len(*zipData)))
	if err != nil {
		return nil, err
	}

	files := make(map[string]*[]byte)

	fmt.Println(len(reader.File))

	for _, f := range reader.File {
		fmt.Println(f.Name)
		zippedFile, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer zippedFile.Close()

		data, err := io.ReadAll(zippedFile)
		if err != nil {
			return nil, err
		}

		files[f.Name] = &data
	}

	return files, nil
}
