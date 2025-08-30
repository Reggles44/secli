package company

import (
	"fmt"

	"github.com/Reggles44/secli/internal/cache"
	"github.com/Reggles44/secli/internal/taxonomy"
	jsonutil "github.com/Reggles44/secli/internal/utils/json"
)

type Facts struct {
	CIK        int                         `json:"cik"`
	EntityName string                      `json:"entityName"`
	DEI        taxonomy.DEI[FactsField]    `json:"dei"`
	Invest     taxonomy.Invest[FactsField] `json:"invest"`
	SRT        taxonomy.SRT[FactsField]    `json:"srt"`
	USGaap     taxonomy.USGaap[FactsField] `json:"us-gaap"`
}

type FactsField struct {
	Label       string     `json:"label"`
	Description string     `json:"description"`
	Units       FactsUnits `json:"units"`
}

type FactsUnits struct {
	Shares []FactValue `json:"shares"`
	USD    []FactValue `json:"USD"`
}

type FactValue struct {
	ACCN         string            `json:"accn"`
	Filed        jsonutil.DateOnly `json:"filed"`
	FilingPeriod string            `json:"fp"`
	FilingYear   int               `json:"fy"`
	Form         string            `json:"form"`
	Start        jsonutil.DateOnly `json:"start"`
	End          jsonutil.DateOnly `json:"end"`
	Value        float64           `json:"val"`
}

func (c Company) Facts() (Facts, error) {
	return cache.FileCache[Facts]{
		URL:      fmt.Sprintf("https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json", c.CIK),
		Duration: 86400,
	}.Read()
}
