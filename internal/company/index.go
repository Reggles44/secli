package company

import (
	"errors"
	"slices"
	"strconv"

	"github.com/Reggles44/secli/internal/cache"
)

type CompanyIndex struct {
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
}

var index CompanyIndex

func init() {
	index, _ = cache.FileCache[CompanyIndex]{
		URL:      "https://www.sec.gov/files/company_tickers_exchange.json",
		FileName: "index.json",
		Duration: 86400,
	}.Read()
}

func Find(search string) (Company, error) {
	for _, entry := range index.Data {
		if slices.Contains(entry, search) {

			// Convert CIK
			cik, err := strconv.Atoi(entry[0])
			if err != nil {
				return Company{}, err
			}

			return Company{cik, entry[1], entry[2]}, nil
		}
	}

	return Company{}, errors.New("no company matching " + search)
}
