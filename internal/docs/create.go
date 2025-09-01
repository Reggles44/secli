package docs

import (
	"encoding/json"
	"sort"

	"github.com/Reggles44/secli/internal/forms"
	"github.com/Reggles44/secli/internal/submissions"
)

func createDocs(cik int) (map[forms.Form][]*Doc, error) {
	var docs []*Doc

	submission, err := submissions.Latest(cik)
	if err != nil {
		return nil, err
	}

	// Assemble Doc
	var dm []map[string]any
	for i := range len(submission.Filings.Recent["form"]) {
		fields := make(map[string]any)
		for k, v := range submission.Filings.Recent {
			fields[k] = v[i]
		}
		fields["cik"] = cik
		dm = append(dm, fields)
	}

	b, err := json.Marshal(dm)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &docs)
	if err != nil {
		return nil, err
	}

	sort.Slice(docs, func(i int, j int) bool {
		return docs[i].ReportDate.Before(docs[j].ReportDate.Time)
	})

	docMap := make(map[forms.Form][]*Doc)

	for _, doc := range docs {
		docMap[doc.Form] = append(docMap[doc.Form], doc)
	}

	return docMap, nil
}
