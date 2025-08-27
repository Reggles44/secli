package company

import (
	"encoding/json"
	"sort"
	"time"

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

	// Data
	DEI    Taxonomy
	USGaap Taxonomy
	Other map[string]Taxonomy
}

type Taxonomy struct {
	Fields map[string]Field
}

type Field struct {
	Name        string
	Label       string
	Description string
	Values      []Value
}

type Value struct {
	Unit       string
	FilingDate time.Time
	Value      float64
}

func (c Company) Docs() ([]Doc, error) {
	var docs []Doc

	submission, err := c.LatestSubmission()
	if err != nil {
		return nil, err
	}
	facts, err := c.Facts()
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

	docAccnMap := make(map[string]Doc)
	for _, doc := range docs {
		docAccnMap[doc.AccessionNumber] = doc
	}

	// Add facts to docs
	for taxonomy, fields := range facts.Facts {
		for fieldName, field := range fields {
			for units, values := range field.Units {
				for _, v := range values {

					// Access Doc
					doc, ok := docAccnMap[v.ACCN]
					if ok {
						var docTaxonomy *Taxonomy
						if taxonomy == "dei" {
							docTaxonomy = &doc.DEI
						} else if taxonomy == "us-gaap" {
							docTaxonomy = &doc.USGaap
						} else {
							docTaxonomy = &make(Taxonomy{})
							doc.Other[taxonomy] = *docTaxonomy
						}

					}
					
				}
			}
		}
	}
	for fieldType, fieldMap := range facts.Data {
		for fieldName, fact := range fieldMap {
			for unit, values := range fact.Units {
				for _, v := range values {
					doc, ok := docAccnMap[v.ACCN]
					if ok {
						if fieldType == "DEI" {
							doc.DEI[fieldName] = v
						}
					}
				}
			}
		}
	}

	sort.Slice(docs, func(i int, j int) bool {
		return docs[i].ReportDate.Before(docs[j].ReportDate.Time)
	})
	return docs, err
}
