package feiertage

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCheckIfIsBankHoliday(t *testing.T) {
	isBankHoliday := CheckIfIsBankHoliday(time.Date(2022, 10, 31, 0, 0, 0, 0, time.UTC), Niedersachsen(2022))
	assert.True(t, isBankHoliday)
	isBankHoliday = CheckIfIsBankHoliday(time.Date(2022, 12, 25, 0, 0, 0, 0, time.UTC), Niedersachsen(2022))
	assert.True(t, isBankHoliday)
	isBankHoliday = CheckIfIsBankHoliday(time.Date(2022, 12, 26, 0, 0, 0, 0, time.UTC), Niedersachsen(2022))
	assert.True(t, isBankHoliday)
	isBankHoliday = CheckIfIsBankHoliday(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Niedersachsen(2023))
	assert.True(t, isBankHoliday)
	isBankHoliday = CheckIfIsBankHoliday(time.Date(2022, 10, 30, 0, 0, 0, 0, time.UTC), Niedersachsen(2022))
	assert.False(t, isBankHoliday)
}
