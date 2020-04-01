package space

type Planet string

var (
	earthSeconds float64 = 31557600
)

// Return earth-years old value for given seconds on specific planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / planetOrbitalPeriod(planet)
}

func planetOrbitalPeriod(planet Planet) float64 {

	// I wonder if there is a cleaner way of handling this.
	// It was switch/case or if statements in my head.
	// Both are a bit cluttered.
	switch {
	case planet == "Mercury":
		return 0.2408467 * earthSeconds
	case planet == "Venus":
		return 0.61519726 * earthSeconds
	case planet == "Earth":
		return earthSeconds
	case planet == "Mars":
		return 1.8808158 * earthSeconds
	case planet == "Jupiter":
		return 11.862615 * earthSeconds
	case planet == "Saturn":
		return 29.447498 * earthSeconds
	case planet == "Uranus":
		return 84.016846 * earthSeconds
	case planet == "Neptune":
		return 164.79132 * earthSeconds
	}

	return earthSeconds
}