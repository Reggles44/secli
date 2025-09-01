package taxonomy

import (
	"encoding/json"

	"github.com/go-viper/mapstructure/v2"
)

type Taxonomy[T any] struct {
	DEI    DEI[T]    `json:"dei"`
	Invest Invest[T] `json:"invest"`
	SRT    SRT[T]    `json:"srt"`
	USGaap USGaap[T] `json:"us-gaap"`
}

func Convert[A any, B any](from *Taxonomy[A], to *Taxonomy[B], conv func(a *A, b *B) error) error {
	b, err := json.Marshal(from)
	if err != nil {
		return err
	}
	fromMap := make(map[string]map[string]A)
	err = json.Unmarshal(b, &fromMap)
	if err != nil {
		return err
	}

	toMap := make(map[string]map[string]B)
	for t, values := range fromMap {
		if check, ok := toMap[t]; check == nil || !ok {
			toMap[t] = make(map[string]B)
		}

		for k, v := range values {
			var b B
			err = conv(&v, &b)
			if err != nil {
				return err
			}
			toMap[t][k] = b
		}
	}

	err = mapstructure.Decode(toMap, &to)
	return err
}
