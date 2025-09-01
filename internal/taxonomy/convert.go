package taxonomy

import (
	"encoding/json"
)

func Convert[A any, B any](from *Taxonomy[A], to *Taxonomy[B], conv func(a *A, b *B) error) error {
	fromMap := make(map[string]map[string]A)
	toMap := make(map[string]map[string]B)

	// Make `from` into map
	bf, err := json.Marshal(from)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bf, &fromMap)
	if err != nil {
		return err
	}

	// Convert Values
	for t, values := range fromMap {
		toValueMap := make(map[string]B)

		for k, a := range values {
			var b B

			err = conv(&a, &b)
			if err != nil {
				return err
			}

			// fmt.Printf("%v %+v\n", k, a)
			// fmt.Printf("%v %+v\n", k, b)

			toValueMap[k] = b
		}

		toMap[t] = toValueMap
		// fmt.Printf("%s %+v\n", t, toValueMap)
	}

	// Make `toMap` into `to`
	bt, err := json.Marshal(toMap)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bt, to)
	if err != nil {
		return err
	}

	return err
}
