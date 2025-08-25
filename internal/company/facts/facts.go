package facts

import (
	"fmt"

	"github.com/Reggles44/secli/internal/cache"
)

type Facts struct {
	CIK        int                        `json:"cik"`
	EntityName string                     `json:"entityName"`
	Data       map[string]map[string]Fact `json:"facts"`
}

type Fact struct {
	Description string  `json:"description"`
	Label       string  `json:"label"`
	Units       map[string][]Value `json:"units"`
}

type Value struct {
	// Filing Information
	ACCN         string `json:"accn"`
	FiledDate    string `json:"filed"`
	FilingPeriod string `json:"fp"`
	FilingYear   int    `json:"fy"`
	Frame        string `json:"frame"`
	Form         string `json:"form"`

	// Value Meta Data
	Start string `json:"start"`
	End   string `json:"end"`

	// Value
	Value float64 `json:"val"`
}

func Get(cik int) (Facts, error) {
	return cache.FileCache[Facts]{
		URL:      fmt.Sprintf("https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json", cik),
		Duration: 86400,
	}.Read()
}
