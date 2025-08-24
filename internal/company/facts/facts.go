package companyfacts

import (
	"fmt"

	"github.com/Reggles44/secli/internal/cache"
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
	url := fmt.Sprintf("https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json", cik)
	factsCache := cache.FileCache[CompanyFacts]{URL: url, Duration: 86400}
	return factsCache.Read()
}
