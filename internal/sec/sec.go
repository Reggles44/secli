package sec

import (
	"encoding/json"

	"github.com/Reggles44/secli/internal/request"
)

var (
	TickerIndex      map[string]int
	CompanyNameIndex map[string]int
)

var (
	indexUrl      = "https://www.sec.gov/files/company_tickers_exchange.json"
	cacheDuration = 86400
)

func init() {
	TickerIndex = make(map[string]int)
	CompanyNameIndex = make(map[string]int)

	resp, err := request.Get("GET", indexUrl, cacheDuration)
	if err != nil {
		panic(err)
	}

	// Write to Companies
	var jsonData interface{}
	err = json.Unmarshal(*resp, &jsonData)
	if err != nil {
		panic(err)
	}

	data, ok := jsonData["data"]

	for _, companyData := range indexFile.Data {
		cik := int(companyData[0].(float64))
		name := companyData[1].(string)
		ticker := companyData[2].(string)

		TickerIndex[ticker] = cik
		CompanyNameIndex[name] = cik
	}
}

