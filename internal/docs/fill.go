package docs

import (
	"github.com/Reggles44/secli/internal/facts"
	"github.com/Reggles44/secli/internal/taxonomy"
)

func (d *Doc) Fill() error {
	factData, err := facts.Get(d.CIK)
	if err != nil {
		return err
	}

	// fmt.Println(d.CIK, d.AccessionNumber, d.FileNumber, d.FilmNumber, d.PrimaryDocument)
	// netIncome, _ := json.Marshal(factData.Taxonomy.USGaap.NetIncomeLoss)
	// fmt.Println(string(netIncome))
	// fmt.Printf("%+v", factData.Taxonomy.USGaap.NetIncomeLoss)

	err = taxonomy.Convert(
		&factData.Taxonomy,
		&d.Taxonomy,
		func(factField *facts.FactsField, docField *DocField) error {
			docField.Label = factField.Label
			docField.Description = factField.Description

			// fmt.Println(factField.Label)
			for unit, values := range factField.Units {
				for _, value := range values {
					// Find matching value
					// fmt.Printf("%v == %v (%v)\n", value.ACCN, d.AccessionNumber, value.ACCN == d.AccessionNumber)
					if value.ACCN == d.AccessionNumber {
						docField.Values = append(docField.Values, DocValue{
							Unit:       unit,
							FilingDate: value.Filed.Time,
							Start:      value.Start.Time,
							End:        value.End.Time,
							Value:      value.Value,
						})
					}
				}
			}

			// return fmt.Errorf("%v no field found for %v", d.AccessionNumber, factField.Label)
			// fmt.Printf("%v no field found for %v\n", d.AccessionNumber, factField.Label)
			return nil
		})

	return err
}
