// Package weather provides tools to track and forecast atmospheric conditions.
package weather

var (
	// CurrentCondition holds the text description of the latest weather state.
	CurrentCondition string

	// CurrentLocation holds the name of the most recently queried city.
	CurrentLocation string
)

// Forecast updates the global weather state and returns a formatted summary string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
