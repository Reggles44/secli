package companyindex

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/Reggles44/secli/internal/cache"
)

type CompanyIndex struct {
	Fields []string            `json:"fields"`
	Data   []CompanyIndexEntry `json:"data"`
}

type CompanyIndexEntry struct {
	CIK    int
	Name   string
	Ticker string
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

var indexCache = cache.FileCache[CompanyIndex]{
	URL:      "https://www.sec.gov/files/company_tickers_exchange.json",
	FileName: "index.json",
	Duration: 86400,
}
var index CompanyIndex

func init() {
	i, err := indexCache.Read()
	if err != nil {
		panic(err)
	}
	index = i
}

func Find(search string) (CompanyIndexEntry, error) {
	for _, entry := range index.Data {
		if strings.EqualFold(entry.Name, search) || strings.EqualFold(entry.Ticker, search) {
			return entry, nil
		}
	}

	return CompanyIndexEntry{}, errors.New("no company matching " + search)
}
