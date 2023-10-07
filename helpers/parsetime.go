package helpers

import "time"

func ParseStringToTime(val string) time.Time {
	layout := "02-01-2006 15:04:05"
	date, _ := time.Parse(layout, val)
	return date
}

func ParseTimeToString(val time.Time) string {
	layout := "02-01-2006 15:04:05"
	date := val.Format(layout)
	return date
}
