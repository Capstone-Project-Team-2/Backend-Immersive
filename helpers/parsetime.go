package helpers

import "time"

func ParseTime(val string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	// layout := "02 January 2006 15:04:05 MST"

	date, err := time.Parse(layout, val)
	return date, err
}
