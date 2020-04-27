package space

// Planet type of string
type Planet string

const earthSeconds float64 = 31557600

var planetsOrbital = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age - Return earth-years old value for given seconds on specific planet
func Age(seconds float64, planet Planet) (age float64) {

	// convert seconds to earth years
	age = seconds / earthSeconds

	// adjust seconds based on select planet
	age /= planetsOrbital[planet]

	//round to the nearest two significant digits
	return age
}
