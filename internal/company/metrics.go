package company

import (
	"errors"
	"fmt"
)

// type Metric struct {
// 	Fields    []string
// 	StartDate time.Time
// 	EndDate   time.Time
// 	Func      func(fields ...string) float64
// }

// func (m Metric) Calculate(c Company) map[Doc]float64 {
// 	return nil
// }

func (d Doc) PeRatio() (float64, error) {
	return 0, nil
}

func (d Doc) EPS() (float64, error) {
	netIncome, ok := d.Fields["NetIncomeLoss"]
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v does not have 'Net Income' field", d))
	}
	dividends, ok := d.Fields["OtherPreferredStockDividendsAndAdjustments"]
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v does not have '' field", d))
	}

	shares, ok := d.Fields["EntityCommonStockSharesOutstanding"]
	if !ok {
		return 0, errors.New(fmt.Sprintf("%v does not have '' field", d))
	}


	return (netIncome.Values - dividends) / shares
	retu
}

func MakeMetric()
