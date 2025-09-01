package company

import (
	"github.com/Reggles44/secli/internal/docs"
	"github.com/Reggles44/secli/internal/forms"
)

type Company struct {
	CIK    int
	Name   string
	Ticker string
}

func (c *Company) Docs() (map[forms.Form][]*docs.Doc, error) {
	return docs.Docs(c.CIK)
}
