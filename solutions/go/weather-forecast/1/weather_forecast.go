// Package weather ...
package weather

var (
    // CurrentCondition is current temperature.
	CurrentCondition string
    // CurrentLocation is current location.
	CurrentLocation  string
)

// Forecast is a function that print the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
