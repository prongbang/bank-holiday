package bank_holiday_test

import (
	"testing"

	bank_holiday "github.com/prongbang/bank-holiday"
)

var ds bank_holiday.Utility

func init() {
	ds = bank_holiday.New()
}

func TestGetFinancialHolidayList(t *testing.T) {
	h := ds.GetFinancialHolidayList("2022")
	if len(h) == 0 {
		t.Errorf("Expect length of holiday list is not zero")
	}
}

func TestGetFinancialHolidayMap(t *testing.T) {
	h := ds.GetFinancialHolidayMap("2022")
	if len(h) == 0 {
		t.Errorf("Expect length of holiday map is not zero")
	}
}
