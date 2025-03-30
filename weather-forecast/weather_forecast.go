// Package weather provides weather forecasting features.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation holds the name of the current location.
var CurrentLocation string

// Forecast returns a weather forecast string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}