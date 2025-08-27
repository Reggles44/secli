package jsonutil

import "time"

type DateOnly struct{ time.Time }

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(time.DateOnly, string(b))
	if err != nil {
		return err
	}

	d.Time = date
	return nil
}
