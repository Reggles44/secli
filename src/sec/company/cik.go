package company

import (
	"errors"

	secapi "github.com/Reggles44/secli/src/sec/api"
)

func Find(search *string) (*Company, error) {
	cik, ok := secapi.TickerIndex[*search]
	if !ok {
		cik, ok = secapi.CompanyNameIndex[*search]
	}

	if !ok {
		return nil, errors.New("company not found")
	}

	return getCompany(cik)
}
