package companyfacts

import (
	"encoding/json"
	"fmt"

	"github.com/Reggles44/secli/internal/utils/request"
)

var (
	companyFactsURL           = "https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json"
	companyFactsCacheDuration = 86400
)

type CompanyFacts struct {
	DEI    map[string]Fact `json:"dei"`
	USGaap map[string]Fact `json:"us-gaap"`
}

type Fact struct {
	Description string             `json:"description"`
	Label       string             `json:"label"`
	Units       map[string][]Value `json:"units"`
}

type Value struct {
	ACCN         string  `json:"accn"`
	PeriodEnd    string  `json:"end"`
	FiledDate    string  `json:"filed"`
	Form         string  `json:"form"`
	FilingPeriod string  `json:"fp"`
	FilingYear   int     `json:"fy"`
	Value        float64 `json:"val"`
}

func Get(cik int) (CompanyFacts, error) {
	url := fmt.Sprintf(companyFactsURL, cik)
	data, err := request.Get("GET", url, companyFactsCacheDuration)
	if err != nil {
		return CompanyFacts{}, nil
	}

	var company CompanyFacts
	err = json.Unmarshal(*data, &company)
	if err != nil {
		return CompanyFacts{}, nil
	}

	return company, nil
}
