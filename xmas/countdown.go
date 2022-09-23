package xmas

import (
	"errors"
	"time"
)

// Days is how many days of Xmas there are.
const Days = 12

// Date is the day of the month Xmas falls on.
const Date = 25

// DaysUntilChristmas returns the number of days left until Xmas, starting with today.
func DaysUntilChristmas() (int, error) {
	_, _, date := time.Now().Date()
	until := Date - date
	if until > Days {
		return until, errors.New("too far away from christmas")
	}
	return until, nil
}
