package metrics

import (
	"fmt"

	"github.com/Reggles44/secli/internal/company"
)



func EBITDA(c company.Company) ([]float32, error) {
	var results []float32

	facts, err := c.Facts()
	if err != nil {
		return nil, err
	}

	operatingIncome, err := facts.Find("OperatingIncomeLoss", "10-Q")
	if err != nil {
		return nil, err
	}

	depreciationAndAmortization, err := facts.Find("DepreciationAndAmortization", "10-Q")
	if err != nil {
		return nil, err
	}

	fmt.Printf("ACCN=%v (%v-%v) %v", )


	

	
	return results, nil

}
