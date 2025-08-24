package facts

import (
	"errors"
	"fmt"
	"time"

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
	Units       []Value `json:"units"`
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

func (f Facts) Find(field string, form string) ([]Value, error) {
	var values []Value

	for _, fmap := range f.Data {
		for f, vlist := range fmap {
			if f == field {

				for _, v := range vlist.Units {
					start, _ := time.Parse("2006-01-02", v.Start)
					end, _ := time.Parse("2006-01-02", v.Start)
					
					duration := end.Sub(start)
					if v.Form == form{
						values = append(values, v)
					}
				}

				return values, nil
			}
		}
	}

	return values, errors.New("could not find field")
}
