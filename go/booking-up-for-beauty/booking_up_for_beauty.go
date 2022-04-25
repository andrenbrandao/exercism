package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"

	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}

	return parsedTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}

	return parsedTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}

	currentHour := parsedTime.Hour()
	return currentHour >= 12 && currentHour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const layout = "1/2/2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}

	formattedTime := parsedTime.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprint("You have an appointment on ", formattedTime, ".")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	anniversaryDate := time.Date(currentYear, 9, 15, 0, 0, 0, 0, time.UTC)

	return anniversaryDate
}
