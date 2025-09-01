package facts

import (
	"github.com/Reggles44/secli/internal/cache"
	"github.com/Reggles44/secli/internal/jsonutil"
	"github.com/Reggles44/secli/internal/taxonomy"
)

type Facts struct {
	CIK        int                           `json:"cik"`
	EntityName string                        `json:"entityName"`
	Taxonomy   taxonomy.Taxonomy[FactsField] `json:"facts"`
}

type FactsField struct {
	Label       string                 `json:"label"`
	Description string                 `json:"description"`
	Units       map[string][]FactValue `json:"units"`
}

type FactValue struct {
	ACCN         string        `json:"accn"`
	Filed        jsonutil.Time `json:"filed"`
	FilingPeriod string        `json:"fp"`
	FilingYear   int           `json:"fy"`
	Form         string        `json:"form"`
	Start        jsonutil.Time `json:"start"`
	End          jsonutil.Time `json:"end"`
	Value        float64       `json:"val"`
}

var facts = &cache.Cache[Facts]{
	URL: "https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json",
}

func Get(cik int) (*Facts, error) {
	return facts.Read(cik)
}

// func (f *Facts) Taxonomy() taxonomy.Taxonomy[FactsField] {
// 	return taxonomy.Taxonomy[FactsField]{
// 		DEI:    f.DEI,
// 		Invest: f.Invest,
// 		SRT:    f.SRT,
// 		USGaap: f.USGaap,
// 	}
// }
