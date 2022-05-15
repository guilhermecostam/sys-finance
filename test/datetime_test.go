package test

import (
	"testing"

	"github.com/guilhermecostam/sys-finance/util"
)

func TestStringToTime(test *testing.T) {
	var convertedTime = util.StringToTime("2010-10-10T10:10:10")

	if convertedTime.Year() != 2010 {
		test.Errorf("Expect the year to be 2010")
	}

	if convertedTime.Month() != 10 {
		test.Errorf("Expect the month to be 10")
	}

	if convertedTime.Hour() != 10 {
		test.Errorf("Expect the hour to be 10")
	}
}
