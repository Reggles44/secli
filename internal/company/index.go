package company

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Reggles44/secli/internal/cache"
)

type CompanyIndex struct {
	Fields []string       `json:"fields"`
	Data   []CompanyEntry `json:"data"`
}

type CompanyEntry struct {
	CIK    int
	Name   string
	Ticker string
	// Exchange string
}

func (ce *CompanyEntry) UnmarshalJSON(b []byte) error {
	var data []any
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	ce.CIK = int(data[0].(float64))
	ce.Name = data[1].(string)
	ce.Ticker = data[2].(string)
	// ce.Exchange = data[3].(string)
	return nil
}

var companyIndex = &cache.Cache[CompanyIndex]{
	URL: "https://www.sec.gov/files/company_tickers_exchange.json",
}

func Find(search string) (*Company, error) {
	index, err := companyIndex.Read()
	if err != nil {
		return nil, err
	}

	for _, entry := range index.Data {
		if entry.Name == search || entry.Ticker == search || strconv.Itoa(entry.CIK) == search {
			company := Company(entry)
			return &company, nil
		}
	}

	return nil, fmt.Errorf("no company matching %v", search)
}
