package companysubmission

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/Reggles44/secli/internal/company/filing"
	"github.com/Reggles44/secli/internal/utils/request"
)

var (
	companySubmissionUrl           = "https://data.sec.gov/submissions/CIK%010d.json"
	companySubmissionCacheDuration = 0
)

type CompanySubmissions struct {
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
			FileNumber            []string `json:"fileNumber"`
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

func Get(cik int) (CompanySubmissions, error) {
	url := fmt.Sprintf(companySubmissionUrl, cik)
	data, err := request.Get("GET", url, companySubmissionCacheDuration)
	if err != nil {
		return CompanySubmissions{}, err
	}

	var submission CompanySubmissions
	json.Unmarshal(*data, &submission)

	return submission, nil
}

func (s CompanySubmissions) GetFilings(formSearch string) []filing.FilingMeta {
	var filings []filing.FilingMeta

	for i, form := range s.Filings.Recent.Form {
		if form == formSearch {
			filingDate, err := time.Parse("2006-01-02", s.Filings.Recent.FilingDate[i])
			if err != nil {
				continue
			}

			reportDate, err := time.Parse("2006-01-02", s.Filings.Recent.ReportDate[i])
			if err != nil {
				continue
			}

			acceptanceDateTime, err := time.Parse("2006-01-02T15:04:05.000Z", s.Filings.Recent.AcceptanceDateTime[i])
			if err != nil {
				continue
			}

			f := filing.FilingMeta{
				CIK:                   s.Cik,
				AccessionNumber:       s.Filings.Recent.AccessionNumber[i],
				FilingDate:            filingDate,
				ReportDate:            reportDate,
				AcceptanceDateTime:    acceptanceDateTime,
				Act:                   s.Filings.Recent.Act[i],
				Form:                  form,
				FileNumber:            s.Filings.Recent.FileNumber[i],
				FilmNumber:            s.Filings.Recent.FilmNumber[i],
				Items:                 s.Filings.Recent.Items[i],
				Core_type:             s.Filings.Recent.Core_type[i],
				Size:                  s.Filings.Recent.Size[i],
				IsXBRL:                s.Filings.Recent.IsXBRL[i],
				IsInlineXBRL:          s.Filings.Recent.IsInlineXBRL[i],
				PrimaryDocument:       s.Filings.Recent.PrimaryDocument[i],
				PrimaryDocDescription: s.Filings.Recent.PrimaryDocDescription[i],
			}
			filings = append(filings, f)

		}
	}

	sort.Slice(filings, func(i int, j int) bool {
		return filings[i].FilingDate.Before(filings[j].FilingDate)
	})

	return filings
}
