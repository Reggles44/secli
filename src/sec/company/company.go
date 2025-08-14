package company

import (
	"encoding/json"
	"fmt"

	secapi "github.com/Reggles44/secli/src/sec/api"
)

var companyFactsURL string = "https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json"

type Company struct {
	CIK        int       `json:"cik"`
	EntityName string       `json:"entityName"`
	Facts      CompanyFacts `json:"facts"`
}

type CompanyFacts struct {
	DEI    struct{} `json:"dei"`
	USGaap struct{} `json:"us-gaap"`
}

func getCompany(cik int) (*Company, error) {
	url := fmt.Sprintf(companyFactsURL, cik)
	data, err := secapi.Request("GET", url, nil, true, 86400)
	if err != nil {
		return nil, err
	}

	var company Company
	err = json.Unmarshal(*data, &company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}
