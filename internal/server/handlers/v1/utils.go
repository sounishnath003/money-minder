package v1

import (
	"fmt"
	"time"
)

// CheckDateRange - parses and check the validity of range
// return are parsed date
//
// Returns fromDate, toDate, error
func CheckDateRange(fromDate, toDate string) (time.Time, time.Time, error) {
	fromDateParsed, err := time.Parse("2006-01-02T15:04:00.000Z", fromDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	toDateParsed, err := time.Parse("2006-01-02T15:04:00.000Z", toDate)
	toDateParsed = toDateParsed.AddDate(0, 0, 1)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if fromDateParsed.After(toDateParsed) {
		return time.Time{}, time.Time{}, fmt.Errorf("from date cannot be after to date")
	}

	return fromDateParsed, toDateParsed, nil
}
