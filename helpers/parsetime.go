package helpers

import "time"

func ParseStringToTime(val string) time.Time {
	layout := "02-01-2006 15:04:05"
	date, _ := time.Parse(layout, val)
	return date
}

func ParseTimeMidtrans(val string) time.Time {
	layout := "2006-01-02 15:04:05"
	date, _ := time.Parse(layout, val)
	return date
}
func ParseTimeToString(val time.Time) string {
	layout := "02-01-2006 15:04:05"
	date := val.Format(layout)
	return date
}
