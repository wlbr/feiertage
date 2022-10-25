package feiertage

import (
	"github.com/pkg/errors"
	"time"
)

func CheckIfIsBankHoliday(date time.Time, region Region) bool {

	for _, feiertag := range region.Feiertage {
		if datesAreEqual(feiertag.Time, date) {
			return true
		}
	}
	return false

}

func datesAreEqual(date1, date2 time.Time) bool {
	if date1.Year() != date2.Year() {
		return false
	}
	if date1.Month() != date2.Month() {
		return false
	}
	if date1.Day() != date2.Day() {
		return false
	}
	return true
}
func CheckIfIsBankHolidayIn(date time.Time, state string) (bool, error) {

	region, err := GetRegionFromString(state, date.Year())
	if err != nil {
		return false, errors.Wrap(err, "Failed to parse region from state string")
	}
	return CheckIfIsBankHoliday(date, region), nil
}
