package preload

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	zipFilePath  string = "/tmp/companyfacts.zip"
	bulkFactsUrl string = "https://www.sec.gov/Archives/edgar/daily-index/bulkdata/submissions.zip"
)

var PreloadCmd = &cobra.Command{
	Use:  "preload",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("GET", bulkFactsUrl, nil)
		req.Header.Add("User-Agent", fmt.Sprintf("%s %s", "SEC CLI Tool", "reginaldbeakes@gmail.com"))

		if err != nil {
			panic(err)
		}

		fmt.Printf("Downloading %v\n", bulkFactsUrl)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		zipFile, err := os.OpenFile(zipFilePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer zipFile.Close()

		io.Copy(zipFile, resp.Body)

		zipReader, err := zip.OpenReader(zipFilePath)
		if err != nil {
			panic(err)
		}
		defer zipReader.Close()

		for _, file := range zipReader.File {
			writeFilePath := path.Join("/home/r/Downloads/", file.Name)
			fmt.Printf("Writing %v\n", writeFilePath)

			writeFile, err := os.OpenFile(writeFilePath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}

			content, err := file.Open()
			if err != nil {
				panic(err)
			}

			io.Copy(writeFile, content)
		}
	},
}
