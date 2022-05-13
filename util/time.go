package util

import "time"

const layout = "2006-01-02T15:04:05"

// StringToTime is a method to converted string in datetime format
func StringToTime(value string) time.Time {
	stringConverted, _ := time.Parse(layout, value)
	return stringConverted
}
