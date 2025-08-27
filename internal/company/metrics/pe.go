package metrics

import (
	"fmt"

	"github.com/Reggles44/secli/internal/company"
)



func PE(c company.Company, form string) ([]float32, error) {
	var results []float32

	// facts, err := c.Facts()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// result := facts.Find(form, "OperatingIncomeLoss", "DepreciationAndAmortization")
	//
	// for accn, fields := range result {
	// 	operatingIncome := fields["OperatingIncomeLoss"]
	// 	depreciationAmortization := fields["DepreciationAndAmortization"]
	//
	// 	ebidta := operatingIncome.Value + depreciationAmortization.Value
	//
	// 	fmt.Printf("%v EBIDTA=%v\n", accn, ebidta)
	// }
	
	return results, nil
}

