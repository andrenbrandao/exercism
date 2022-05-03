package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var result []Record

	for _, v := range in {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	predicate := func(record Record) bool {
		return p.From <= record.Day && record.Day <= p.To
	}

	return predicate
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	predicate := func(record Record) bool {
		return record.Category == c
	}

	return predicate
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	records := Filter(in, ByDaysPeriod(p))

	total := 0.0

	for _, v := range records {
		total += v.Amount
	}

	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	records_with_category := Filter(in, ByCategory(c))

	if len(records_with_category) == 0 {
		return 0, errors.New(fmt.Sprintf("unknown category %s", c))
	}

	return TotalByPeriod(records_with_category, p), nil
}
