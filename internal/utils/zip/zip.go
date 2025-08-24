package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
)

func Unzip(data *[]byte) (map[string]*[]byte, error) {
	files := make(map[string]*[]byte)

	zipReader, err := zip.NewReader(
		bytes.NewReader(*data),
		int64(len(*data)),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range zipReader.File {
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
