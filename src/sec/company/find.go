package company

import (
	"errors"

	secapi "github.com/Reggles44/secli/src/sec/api"
)

func Find(search *string) error {
	cik, ok := secapi.TickerIndex[*search]
	if !ok {
		cik, ok = secapi.CompanyNameIndex[*search]
	}

	if !ok {
		return errors.New("Company not found")
	}

	secapi.GetSubmissions(cik)
	return nil
}
