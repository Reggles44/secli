package docs

import "time"

type DocValue struct {
	Unit       string
	FilingDate time.Time
	Start      time.Time
	End        time.Time
	Value      float64
}

