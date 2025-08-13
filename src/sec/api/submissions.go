package secapi

import (
	"encoding/json"
	"fmt"
	"log"
)

const baseSubmissionURL string = "https://data.sec.gov/submissions/CIK%010d.json"

type Submission struct {
	Cik                               string   `json:"cik"`
	EntityType                        string   `json:"entityType"`
	SIC                               string   `json:"sic"`
	SICDescription                    string   `json:"sicDescription"`
	InsiderTransactionForOwnerExists  int      `json:"insiderTransactionForOwnerExists"`
	InsiderTransactionForIssuerExists int      `json:"insiderTransactionForIssuerExists"`
	Name                              string   `json:"name"`
	Tickers                           []string `json:"tickers"`
	Exchanges                         []string `json:"exchanges"`
	EIN                               string   `json:"ein"`
	Lei                               string   `json:"lei"`
	Description                       string   `json:"description"`
	Website                           string   `json:"website"`
	InvestorWebsite                   string   `json:"investorWebsite"`
	Category                          string   `json:"category"`
	FiscalYearEnd                     string   `json:"fiscalYearEnd"`
	StateOfIncorporation              string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription   string   `json:"stateOfIncorporationDescription"`
	Addresses                         struct {
		Mailing struct {
			Street1                   string `json:"street1"`
			Street2                   string `json:"street2"`
			City                      string `json:"city"`
			StateOrCountry            string `json:"stateOrCountry"`
			ZipCode                   string `json:"zipCode"`
			StateOrCountryDescription string `json:"stateOrCountryDescription"`
			IsForeignLocation         int8   `json:"isForeignLocation"`
			ForeignStateTerritory     string `json:"foreignStateTerritory"`
			Country                   string `json:"country"`
			CountryCode               string `json:"countryCode"`
		} `json:"mailing"`
		Business struct{} `json:"business"`
	} `json:"addresses"`
	Phone       string   `json:"phone"`
	Flags       string   `json:"flags"`
	FormerNames []string `json:"formerNames"`
	Filings     struct {
		Recent struct {
			AccessionNumber       []string `json:"accessionNumber"`
			FilingDate            []string `json:"filingDate"`
			ReportDate            []string `json:"reportDate"`
			AcceptanceDateTime    []string `json:"acceptanceDateTime"`
			Act                   []string `json:"act"`
			Form                  []string `json:"form"`
			FilingNumber          []string `json:"filingNumber"`
			FilmNumber            []string `json:"filmNumber"`
			Items                 []string `json:"items"`
			Core_type             []string `json:"core_type"`
			Size                  []int64  `json:"size"`
			IsXBRL                []int8   `json:"isXBRL"`
			IsInlineXBRL          []int8   `json:"isInlineXBRL"`
			PrimaryDocument       []string `json:"primaryDocument"`
			PrimaryDocDescription []string `json:"primaryDocDescription"`
		} `json:"recent"`
		Files []struct {
			Name        string `json:"name"`
			FilingCount int64  `json:"filingCount"`
			FilingFrom  string `json:"filingFrom"`
			FilingTo    string `json:"filingTo"`
		} `json:"files"`
	} `json:"filings"`
}

func GetSubmission(cik int) (*Submission, error) {
	url := fmt.Sprintf(baseSubmissionURL, cik)

	data, err := Request("GET", url, nil, false, 0)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(*data))

	var submission Submission
	err = json.Unmarshal(*data, &submission)
	if err != nil {
		log.Panic(err)
	}

	return &submission, nil
}
