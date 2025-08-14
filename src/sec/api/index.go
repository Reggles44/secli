package secapi

import (
	"encoding/json"
)

var (
	TickerIndex      map[string]int
	CompanyNameIndex map[string]int
)

var index_url = "https://www.sec.gov/files/company_tickers_exchange.json"

type IndexFile struct {
	Fields []string `json:"fields"`
	Data   [][]any  `json:"data"`
}

func init() {
	TickerIndex = make(map[string]int)
	CompanyNameIndex = make(map[string]int)

	data, err := Request("GET", index_url, nil, true, 86400)
	if err != nil {
		panic(err)
	}

	// Write to Companies
	var indexFile IndexFile
	err = json.Unmarshal(*data, &indexFile)
	if err != nil {
		panic(err)
	}

	for _, companyData := range indexFile.Data {
		cik := int(companyData[0].(float64))
		name := companyData[1].(string)
		ticker := companyData[2].(string)

		TickerIndex[ticker] = cik
		CompanyNameIndex[name] = cik
	}
}
