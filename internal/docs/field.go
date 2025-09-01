package docs

import (
	"encoding/json"
	"fmt"
)

type DocField struct {
	Label       string
	Description string
	// QuarterlyValue DocValue
	// YearlyValue    DocValue
	Values []DocValue
}

func (df *DocField) TTMValue() (float64, error) {
	b, err := json.Marshal(df)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(b))

	for _, dv := range df.Values {
		duration := dv.End.Sub(dv.Start)
		months := duration.Hours() / 24 / 30
		if int(months) == 12 {
			return dv.Value, nil
		}
	}

	return 0, fmt.Errorf("%v missing yearly value", df.Label)
}


func (df *DocField) Value() (float64, error) {
	if len(df.Values) != 1 {
		return 0, fmt.Errorf("%v not one value", df.Label)
	}

	return df.Values[0].Value, nil
}
