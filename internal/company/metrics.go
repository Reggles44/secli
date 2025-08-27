package company

import "time"

type Metric struct {
	Fields    []string
	StartDate time.Time
	EndDate   time.Time
	Func      func(fields ...string) float64
}

func (m Metric) Calculate(c Company) map[Doc]float64 { 

}
