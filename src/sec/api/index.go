package secapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	TickerIndex      map[string]int
	CompanyNameIndex map[string]int
)

var (
	index_url = "https://www.sec.gov/files/company_tickers_exchange.json"
	tmp_file  = "/tmp/company_tickers_exchange.json"
)

type IndexFile struct {
	Fields []string `json:"fields"`
	Data   [][]any  `json:"data"`
}

func init() {
	TickerIndex = make(map[string]int)
	CompanyNameIndex = make(map[string]int)

	// If the file does *NOT* exist
	if _, err := os.Stat(tmp_file); errors.Is(err, os.ErrNotExist) {
		err := download()
		if err != nil {
			os.Remove(tmp_file)
		}
	}

	err := load()
	if err != nil {
		log.Fatal(err)
	}
}

func download() error {
	fmt.Println("Downloading sec indicies")

	resp, err := Request("GET", index_url, nil)
	if err != nil {
		return err
	}

	// Creat File
	file, err := os.Create(tmp_file)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func load() error {
	// Open file
	file, err := os.Open(tmp_file)
	if err != nil {
		return err
	}

	// Read file
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Write to Companies
	var indexFile IndexFile
	err = json.Unmarshal(data, &indexFile)
	if err != nil {
		return err
	}

	for _, companyData := range indexFile.Data {
		cik := int(companyData[0].(float64))
		name := companyData[1].(string)
		ticker := companyData[2].(string)

		TickerIndex[ticker] = cik
		CompanyNameIndex[name] = cik
	}
	return nil
}
