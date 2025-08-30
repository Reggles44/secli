package company

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/Reggles44/secli/internal/taxonomy"
	jsonutil "github.com/Reggles44/secli/internal/utils/json"
)

type Doc struct {
	AccessionNumber       string            `json:"accessionNumber"`
	FilingDate            jsonutil.DateOnly `json:"filingDate"`
	ReportDate            jsonutil.DateOnly `json:"reportDate"`
	AcceptanceDateTime    jsonutil.DateOnly `json:"acceptanceDateTime"`
	Form                  string            `json:"form"`
	FileNumber            string            `json:"fileNumber"`
	FilmNumber            string            `json:"filmNumber"`
	PrimaryDocument       string            `json:"primaryDocument"`
	PrimaryDocDescription string            `json:"primaryDocDescription"`

	DEI    taxonomy.DEI[DocField]    `json:"dei"`
	Invest taxonomy.Invest[DocField] `json:"invest"`
	SRT    taxonomy.SRT[DocField]    `json:"srt"`
	USGaap taxonomy.USGaap[DocField] `json:"us-gaap"`
}

type DocField struct {
	Label          string
	Description    string
	QuarterlyValue DocValue
	YearlyValue    DocValue
}

type DocValue struct {
	Unit       string
	FilingDate time.Time
	Value      float64
}

func (c Company) Docs() ([]Doc, error) {
	docs, err := createDocs(c)
	if err != nil {
		return nil, err
	}

	err = populateDocs(c, &docs)
	return docs, err
}

func createDocs(c Company) ([]Doc, error) {
	var docs []Doc

	submission, err := c.LatestSubmission()
	if err != nil {
		return nil, err
	}

	// Assemble Doc
	var dm []map[string]string
	for i := range len(submission.Filings.Recent["form"]) {
		fields := make(map[string]string)
		for k, v := range submission.Filings.Recent {
			fields[k] = v[i]
		}
		dm = append(dm, fields)
	}

	b, err := json.Marshal(dm)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &docs)

	sort.Slice(docs, func(i int, j int) bool {
		return docs[i].ReportDate.Before(docs[j].ReportDate.Time)
	})
	return docs, err
}

func populateDocs(c Company, docs *[]Doc) error {
	docAccnMap := make(map[string]*Doc)
	for _, doc := range *docs {
		docAccnMap[doc.AccessionNumber] = &doc
	}

	facts, err := c.Facts()
	if err != nil {
		return err
	}

	for _, doc := range *docs {
		doc.DEI, err = taxonomy.ConvertDEI(facts.DEI, FactFieldToDocField)
	}


	return nil
}

func FactFieldToDocField(from FactsField) (DocField, error) {

}
