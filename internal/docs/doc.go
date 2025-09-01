package docs

import (
	"time"

	"github.com/Reggles44/secli/internal/forms"
	"github.com/Reggles44/secli/internal/jsonutil"
	"github.com/Reggles44/secli/internal/taxonomy"
)

type Doc struct {
	CIK                   int
	AccessionNumber       string                      `json:"accessionNumber"`
	FilingDate            jsonutil.Time               `json:"filingDate"`
	ReportDate            jsonutil.Time               `json:"reportDate"`
	AcceptanceDateTime    jsonutil.Time               `json:"acceptanceDateTime"`
	Form                  forms.Form                  `json:"form"`
	FileNumber            string                      `json:"fileNumber"`
	FilmNumber            string                      `json:"filmNumber"`
	PrimaryDocument       string                      `json:"primaryDocument"`
	PrimaryDocDescription string                      `json:"primaryDocDescription"`
	Taxonomy              taxonomy.Taxonomy[DocField] `json:"taxonomy"`
}

type DocField struct {
	Label       string
	Description string
	// QuarterlyValue DocValue
	// YearlyValue    DocValue
	Value DocValue
}

type DocValue struct {
	Unit       string
	FilingDate time.Time
	Value      float64
}

func Docs(cik int) (map[forms.Form][]*Doc, error) {
	return createDocs(cik)
}
