package company

import (
	"fmt"

	"github.com/Reggles44/secli/internal/cache"
	jsonutil "github.com/Reggles44/secli/internal/utils/json"
)

type Facts struct {
	CIK        int    `json:"cik"`
	EntityName string `json:"entityName"`

	// Taxonomy -> Field -> Field Info
	Facts map[string]map[string]struct {
		Description string `json:"description"`
		Label       string `json:"label"`

		// Unit Type -> Slice of Value Objects
		// Unit Type can be things like "shares" or "USD"
		Units map[string][]struct {
			ACCN         string            `json:"accn"`
			FiledDate    jsonutil.DateOnly `json:"filed"`
			FilingPeriod string            `json:"fp"`
			Form         string            `json:"form"`
			Start        jsonutil.DateOnly `json:"start"`
			End          jsonutil.DateOnly `json:"end"`
			Value        float64           `json:"val"`
		} `json:"units"`
	} `json:"facts"`
}

func (c Company) Facts() (Facts, error) {
	return cache.FileCache[Facts]{
		URL:      fmt.Sprintf("https://data.sec.gov/api/xbrl/companyfacts/CIK%010d.json", c.CIK),
		Duration: 86400,
	}.Read()
}
