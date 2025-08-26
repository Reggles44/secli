package company

import (
	"fmt"

	"github.com/Reggles44/secli/internal/cache"
	jsonutil "github.com/Reggles44/secli/internal/utils/json"
)

type Facts struct {
	CIK        int                        `json:"cik"`
	EntityName string                     `json:"entityName"`
	Data       map[string]map[string]Fact `json:"facts"`
}

type Fact struct {
	Description string             `json:"description"`
	Label       string             `json:"label"`
	Units       map[string][]Value `json:"units"`
}

type Value struct {
	// Filing Information
	ACCN         string            `json:"accn"`
	FiledDate    jsonutil.DateOnly `json:"filed"`
	FilingPeriod string            `json:"fp"`
	FilingYear   int               `json:"fy"`
	Frame        string            `json:"frame"`
	Form         string            `json:"form"`

	// Value Meta Data
	Start jsonutil.DateOnly `json:"start"`
	End   jsonutil.DateOnly `json:"end"`

	// Value
	Value float64 `json:"val"`
}

func (c Company) Facts() (Facts, error) {
	return cache.FileCache[Facts]{
		URL:      fmt.Sprintf("https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json", cik),
		Duration: 86400,
	}.Read()
}
