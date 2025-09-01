package jsonutil

import (
	"strings"
	"time"
)

const ISO8601 = "2006-01-02T15:04:05-0700Z"

var TimeLayouts = []string{
	// time.Layout     ,//01/02 03:04:05PM '06 -0700 // The reference time, in numerical order.
	// time.ANSIC      ,//Mon Jan _2 15:04:05 2006
	// time.UnixDate   ,//Mon Jan _2 15:04:05 MST 2006
	// time.RubyDate   ,//Mon Jan 02 15:04:05 -0700 2006
	// time.RFC822     ,//02 Jan 06 15:04 MST
	// time.RFC822Z    ,//02 Jan 06 15:04 -0700 // RFC822 with numeric zone
	// time.RFC850     ,//Monday, 02-Jan-06 15:04:05 MST
	// time.RFC1123    ,//Mon, 02 Jan 2006 15:04:05 MST
	// time.RFC1123Z   ,//Mon, 02 Jan 2006 15:04:05 -0700 // RFC1123 with numeric zone
	time.RFC3339,     // 2006-01-02T15:04:05Z07:00
	time.RFC3339Nano, // 2006-01-02T15:04:05.999999999Z07:00
	// time.Kitchen    ,//3:04PM
	// Handy time stamps.
	// time.Stamp     ,//Jan _2 15:04:05
	// time.StampMilli,//Jan _2 15:04:05.000
	// time.StampMicro,//Jan _2 15:04:05.000000
	// time.StampNano ,//Jan _2 15:04:05.000000000
	time.DateTime, // 2006-01-02 15:04:05
	time.DateOnly, // 2006-01-02
	// time.TimeOnly  ,//15:04:05
	ISO8601,
}

type Time struct{ time.Time }

func (d *Time) UnmarshalJSON(b []byte) error {
	var err error

	value := strings.ReplaceAll(string(b), "\"", "")
	if value == "" {
		return nil
	}

	for _, layout := range TimeLayouts {
		d.Time, err = time.Parse(layout, value)
		if err == nil {
			return nil
		}
	}
	return err
}
