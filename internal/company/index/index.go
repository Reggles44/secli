package companyindex

import (
	"encoding/json"
	"errors"

	"github.com/Reggles44/secli/internal/utils/request"
)

var (
	companyIndexUrl           = "https://www.sec.gov/files/company_tickers_exchange.json"
	companyIndexCacheDuration = 86400
)

type CompanyIndex struct {
	Fields []string            `json:"fields"`
	Data   []CompanyIndexEntry `json"data"`
}

type CompanyIndexEntry struct {
	CIK      int
	Name     string
	Ticker   string
	// Exchange string
}

func (cie *CompanyIndexEntry) UnmarshalJSON(data []byte) error {
	var d []any
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}

	cie.CIK = int(d[0].(float64))
	cie.Name = d[1].(string)
	cie.Ticker = d[2].(string)
	// cie.Exchange = d[3].(string)

	return nil
}

func Find(search string) (CompanyIndexEntry, error) {
	resp, err := request.Get("GET", companyIndexUrl, companyIndexCacheDuration)
	if err != nil {
		panic(err)
	}

	var index CompanyIndex
	err = json.Unmarshal(*resp, &index)
	if err != nil {
		panic(err)
	}

	for _, entry := range index.Data {
		if entry.Name == search || entry.Ticker == search {
			return entry, nil
		}
	}

	return CompanyIndexEntry{}, errors.New("no company matching " + search)
}
