package taxonomy

import (
	"encoding/json"

	"github.com/go-viper/mapstructure/v2"
)

type Taxonomy[T any] interface {
	DEI[T] | Invest[T] | SRT[T] | USGaap[T]
}

func convertTaxonomy[A any, B any, F Taxonomy[A], T Taxonomy[B]](from F, c func(from A) (B, error)) (T, error) {
	var to T

	b, err := json.Marshal(from)
	if err != nil {
		return to, err
	}
	var fromMap map[string]A
	err = json.Unmarshal(b, &fromMap)
	if err != nil {
		return to, err
	}

	toMap := make(map[string]B)
	for k, v := range fromMap {
		tov, err := c(v)
		if err != nil {
			return to, err
		}
		toMap[k] = tov
	}

	err = mapstructure.Decode(toMap, &to)
	return to, err
}

func ConvertDEI[A any, B any](from DEI[A], convert func(from A) (B, error)) (DEI[B], error) {
	return convertTaxonomy[A, B, DEI[A], DEI[B]](from, convert)
}

func ConvertInvest[A any, B any](from Invest[A], convert func(from A) (B, error)) (Invest[B], error) {
	return convertTaxonomy[A, B, Invest[A], Invest[B]](from, convert)
}

func ConvertSRT[A any, B any](from SRT[A], convert func(from A) (B, error)) (SRT[B], error) {
	return convertTaxonomy[A, B, SRT[A], SRT[B]](from, convert)
}

func ConvertUSGaap[A any, B any](from USGaap[A], convert func(from A) (B, error)) (USGaap[B], error) {
	return convertTaxonomy[A, B, USGaap[A], USGaap[B]](from, convert)
}
