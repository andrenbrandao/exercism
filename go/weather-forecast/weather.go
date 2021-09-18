// Package weather implements routines related to weather conditions
// on different locations.
package weather

var (
	// CurrentCondition stores the current weather.
	CurrentCondition string

	// CurrentLocation stores the current city.
	CurrentLocation string
)

// Log stores the current city and weather condition and returns a text
// with the current values.
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
