package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	index_url       = "https://www.sec.gov/files/company_tickers_exchange.json"
	tmp_file        = "/tmp/company_tickers_exchange.json"
)

type Company struct {
	CIK      int    `json:"cik"`
	Name     string `json:"name"`
	Ticker   string `json:"ticker"`
	Exchange string `json:"exchange"`
}

type Index struct {
	Companies []Company
}

func Find(search *string, company *string, email *string) (Company, error) {
	// If the file does *NOT* exist
	if _, err := os.Stat(tmp_file); err != nil {
		err := download(company, email)
		if err != nil {
			return Company{}, err
		}
	}

	index, err := load()
	if err != nil {
		return Company{}, err
	}

	for _, company := range index.Companies {
		if company.Ticker == *search || company.Name == *search {
			return company, nil
		}
	}

	return Company{}, errors.New("Company not found")
}

func download(company *string, email *string) error {
	// Create Request
	req, err := http.NewRequest("GET", index_url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", fmt.Sprintf("%v %v", company, email))

	// Do Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Creat File
	file, err := os.Create(tmp_file)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write to File
	io.Copy(file, resp.Body)

	return nil
}

func load() (Index, error) {
	file, err := os.Open(tmp_file)
	if err != nil {
		return Index{}, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return Index{}, err
	}

	var index Index

	json.Unmarshal(data, &index)

	return index, nil
}
