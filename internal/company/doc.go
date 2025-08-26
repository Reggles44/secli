package company

import (
	"encoding/json"

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
}

func (c Company) Docs() ([]Doc, error) {
	var docs []Doc

	submission, err := c.LatestSubmission()
	if err != nil {
		return nil, err
	}

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
	return docs, err
}
