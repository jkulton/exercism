package space

type Planet string

const EarthSecondsInYear = 31557600

var OrbitalPeriods = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the earth age an individual based on number of seconds and planet they reside on
func Age(seconds float64, planet Planet) float64 {
	if val, ok := OrbitalPeriods[planet]; ok {
		earthAge := seconds / EarthSecondsInYear
		return earthAge / val
	} else {
		return 0
	}
}
