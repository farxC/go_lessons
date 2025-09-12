package exercises

import (
	"time"
)

// Easy solution :(
func FindDay(date time.Time) string {
	day := date.Weekday()
	return day.String()
}

// Hard Solution (needs to implement)
func FindDayByDate(date string) string {
	return ""
}
