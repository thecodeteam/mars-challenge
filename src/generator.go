package main

import (
	"math/rand"
	"time"
)

const (
	maxTemperature       = 35.00
	minTemperature       = -142.00
	maxRadiation         = 1000
	minRadiation         = 0
	variationTemperature = 5.00
	variationRadiation   = 20
)

// Reading contains the current sensor readings
type Reading struct {
	SolarFlare  bool    `json:"solarFlare"`
	Temperature float64 `json:"temperature"`
	Radiation   int     `json:"radiation"`
}

func (s *Reading) updateSolarFlare() {
	x := rand.Intn(2)
	if x != 0 {
		s.SolarFlare = true
	} else {
		s.SolarFlare = false
	}
}

func (s *Reading) updateTemperature() {
	//TODO: consider solar Flare

	temperature := (rand.Float64() * ((s.Temperature + variationTemperature) - (s.Temperature - variationTemperature))) + (s.Temperature - variationTemperature)
	if temperature < minTemperature {
		temperature = minTemperature
	} else if temperature > maxTemperature {
		temperature = maxTemperature
	}
	s.Temperature = temperature
}

func (s *Reading) updateRadiation() {
	//TODO: consider solar Flare
	radiation := rand.Intn((s.Radiation+variationRadiation)-(s.Radiation-variationRadiation)) + (s.Radiation - variationRadiation)
	if radiation < minRadiation {
		radiation = minRadiation
	} else if radiation > maxRadiation {
		radiation = maxRadiation
	}
	s.Radiation = radiation
}

func solarFlareRoutine(s *Reading) {
	for {
		s.updateSolarFlare()
		// fmt.Println("Flare Status:", s.SolarFlare)
		if s.SolarFlare == true {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(30 * time.Second)
		}
	}
}

func temperatureRoutine(s *Reading) {
	for {
		s.updateTemperature()
		// fmt.Println("Temperature:", s.Temperature)

		time.Sleep(time.Second)
	}
}

func radiationRoutine(s *Reading) {
	for {
		s.updateRadiation()
		// fmt.Println("Radiation:", s.Radiation)

		time.Sleep(time.Second)
	}
}
