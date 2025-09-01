package docs

import (
	"fmt"

	"github.com/Reggles44/secli/internal/facts"
	"github.com/Reggles44/secli/internal/taxonomy"
)

func (d *Doc) Fill() error {
	factData, err := facts.Get(d.CIK)
	if err != nil {
		return err
	}

	fmt.Println(d)

	err = taxonomy.Convert(
		&factData.Taxonomy,
		&d.Taxonomy,
		func(factField *facts.FactsField, docField *DocField) error {
			docField.Label = factField.Label
			docField.Description = factField.Description

			for unit, values := range factField.Units {
				for _, value := range values {

					// Find matching value
					// fmt.Printf("%v == %v (%v)", value.ACCN, d.AccessionNumber, value.ACCN == d.AccessionNumber)
					if value.ACCN == d.AccessionNumber {
						docField.Value.Unit = unit
						docField.Value.FilingDate = value.Filed.Time
						docField.Value.Value = value.Value
					}
				}
			}


			// return fmt.Errorf("%v no field found for %v", d.AccessionNumber, factField.Label)
			fmt.Printf("%v no field found for %v\n", d.AccessionNumber, factField.Label)
			return nil
		})

	return err
}
