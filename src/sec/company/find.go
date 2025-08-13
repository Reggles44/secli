package company

import (
	"errors"

	secapi "github.com/Reggles44/secli/src/sec/api"
)

func Find(search *string) (*secapi.Submission, error) {
	cik, ok := secapi.TickerIndex[*search]
	if !ok {
		cik, ok = secapi.CompanyNameIndex[*search]
	}

	if !ok {
		return nil, errors.New("company not found")
	}

	return secapi.GetSubmission(cik)
}
