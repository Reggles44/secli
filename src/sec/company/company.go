package company

import (
	"encoding/json"
	"errors"
	"fmt"

	secapi "github.com/Reggles44/secli/src/sec/api"
)

var companyFactsURL string = "https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json"

type Company struct {
	CIK        int          `json:"cik"`
	EntityName string       `json:"entityName"`
	Facts      CompanyFacts `json:"facts"`
}

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

func (fact Fact) FindValue(filedDate string, form string) (float64, error) {

	for _, value := range fact.Units["USD"] {
		if value.FiledDate == filedDate && value.Form == form {
			return value.Value , nil
		}
	}

	return nil, errors.New("No value found")
	
}
