package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2010-10-10T10:10:10")

	if convertedTime.Year() != 2010 {
		t.Errorf("Expect the year to be 2010")
	}

	if convertedTime.Month() != 10 {
		t.Errorf("Expect the month to be 10")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Expect the hour to be 10")
	}
}
