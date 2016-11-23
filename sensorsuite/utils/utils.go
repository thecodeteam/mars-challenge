package utils

import (
	"math/rand"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
)

// GetNewFlare returns a new, random Solar Flare value
func GetNewFlare() bool {
	x := rand.Intn(2)
	if x != 0 {
		return true
	}
	return false
}

// GetNewTempTrend returns a new Temperature Trend value
func GetNewTempTrend(temp float64, flare bool) bool {
	ratio := (temp - ss.MinTemp) / (ss.MaxTemp - ss.MinTemp)
	chance := rand.Float64()
	return chance > ratio || flare
}

// GetNewTemp returns a new Temperature value
func GetNewTemp(currTemp float64, uptrend bool) float64 {
	var min float64
	var max float64

	if uptrend {
		max = currTemp + ss.VariationTemp
		min = currTemp
	} else {
		max = currTemp
		min = currTemp - ss.VariationTemp
	}

	temperature := (rand.Float64() * (max - min)) + min
	if temperature < ss.MinTemp {
		temperature = ss.MinTemp
	} else if temperature > ss.MaxTemp {
		temperature = ss.MaxTemp
	}
	return temperature
}

// GetNewRadiation returns a new Radiation value
func GetNewRadiation(currRad int, uptrend bool) int {
	var min int
	var max int

	if uptrend {
		max = currRad + ss.VariationRadiation
		min = currRad
	} else {
		max = currRad
		min = currRad - ss.VariationRadiation
	}

	radiation := rand.Intn(max-min) + min
	if radiation < ss.MinRadiation {
		radiation = ss.MinRadiation
	} else if radiation > ss.MaxRadiation {
		radiation = ss.MaxRadiation
	}
	return radiation
}

// GetNewRadiationTrend returns a new radiation trend value
func GetNewRadiationTrend(rad int, flare bool) bool {
	ratio := (float64)(rad-ss.MinRadiation) /
		(float64)(ss.MaxRadiation-ss.MinRadiation)
	chance := rand.Float64()
	return chance > ratio || flare
}
