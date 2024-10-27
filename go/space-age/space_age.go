// Package space handles planetary conversions.
package space

const (
	secondsInAYear = 31557600
)

// Planet is a type for planet names.
type Planet string

// Age determines the age in years on the given
// planet for the given seconds.
func Age(seconds float64, planet Planet) float64 {
	var planetAge float64

	// Convert the amount of seconds into earth years.
	earthYears := seconds / secondsInAYear

	switch planet {
	case "Mercury":
		planetAge = earthYears / 0.2408467
	case "Venus":
		planetAge = earthYears / 0.61519726
	case "Mars":
		planetAge = earthYears / 1.8808158
	case "Jupiter":
		planetAge = earthYears / 11.862615
	case "Saturn":
		planetAge = earthYears / 29.447498
	case "Uranus":
		planetAge = earthYears / 84.016846
	case "Neptune":
		planetAge = earthYears / 164.79132
	case "Earth":
		planetAge = earthYears
	}

	return planetAge
}
