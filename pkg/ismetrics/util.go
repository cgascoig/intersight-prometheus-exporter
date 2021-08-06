package ismetrics

import (
	"fmt"
	"time"
)

func optStr(s string) *string {
	return &s
}

func optFloat64(f float64) *float64 {
	return &f
}

func getIntervalString(endTime time.Time, duration time.Duration) string {
	startTime := endTime.Add(-duration)

	layout := "2006-01-02T15:04:05-07:00"

	return fmt.Sprintf("%s/%s", startTime.Format(layout), endTime.Format(layout))
}
