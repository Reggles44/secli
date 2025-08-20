package company

import (
	companyfacts "github.com/Reggles44/secli/internal/company/facts"
	companyindex "github.com/Reggles44/secli/internal/company/index"
	companysubmission "github.com/Reggles44/secli/internal/company/submission"
)

type Company struct {
	CIK    int
	Name   string
	Ticker string
}

func Find(search string) (Company, error) {
	entry, err := companyindex.Find(search)
	if err != nil {
		return Company{}, err
	}

	return Company(entry), nil
}

func (c Company) LatestSubmission() (companysubmission.CompanySubmissions, error) {
	return companysubmission.Get(c.CIK)
}

func (c Company) Facts() (companyfacts.CompanyFacts, error) {
	return companyfacts.Get(c.CIK)
}
