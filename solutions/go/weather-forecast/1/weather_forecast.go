// Package weather  provides functionality to get current weather condition for a given location.
package weather

var (
	// CurrentCondition holds the current weather condition as a string.
	CurrentCondition string
    // CurrentLocation holds the name of the current location as a string.
	CurrentLocation  string
)

// Forecast returns a formatted string with the current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
