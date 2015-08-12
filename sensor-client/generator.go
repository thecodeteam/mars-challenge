package main

import (
	"math/rand"
	"time"
	"fmt"
)
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
type JSONTime time.Time

func (t *JSONTime)MarshalJSON() ([]byte, error) {
    //do your serializing here
    stamp := fmt.Sprintf("\"%s\"",time.Now().Format("2006-01-02T15:04:05Z07:00"))
    return []byte(stamp), nil
}

const (
	variationTemperature = 5.00
	variationRadiation   = 20
	maxTrendSeconds      = 20
	minTrendSeconds      = 5
)

// Reading contains the current sensor readings
type Reading struct {
	SolarFlare         bool    `json:"solarFlare"`
	Temperature        float64 `json:"temperature"`
	Radiation          int     `json:"radiation"`
	temperatureUptrend bool
	radiationUptrend   bool
	Stamp		  JSONTime `json:"stamp"`
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
	var min float64
	var max float64

	if s.temperatureUptrend {
		max = s.Temperature + variationTemperature
		min = s.Temperature
	} else {
		max = s.Temperature
		min = s.Temperature - variationTemperature
	}

	temperature := (rand.Float64() * (max - min)) + min
	if temperature < minTemperature {
		temperature = minTemperature
	} else if temperature > maxTemperature {
		temperature = maxTemperature
	}
	s.Temperature = temperature
}

func (s *Reading) updateTemperatureTrend() {
	ratio := (s.Temperature - minTemperature) / (maxTemperature - minTemperature)
	chance := rand.Float64()
	s.temperatureUptrend = chance > ratio || s.SolarFlare
	//log.Printf("[Temperature] Ratio: %.2f, Change: %.2f, Uptrend: %t\n", ratio, chance, s.temperatureUptrend)
}

func (s *Reading) updateRadiation() {
	var min int
	var max int

	if s.radiationUptrend {
		max = s.Radiation + variationRadiation
		min = s.Radiation
	} else {
		max = s.Radiation
		min = s.Radiation - variationRadiation
	}

	radiation := rand.Intn(max-min) + min
	if radiation < minRadiation {
		radiation = minRadiation
	} else if radiation > maxRadiation {
		radiation = maxRadiation
	}
	s.Radiation = radiation
}

func (s *Reading) updateRadiationTrend() {
	ratio := (float64)(s.Radiation-minRadiation) / (float64)(maxRadiation-minRadiation)
	chance := rand.Float64()
	s.radiationUptrend = chance > ratio || s.SolarFlare
	//log.Printf("[Radiation] Ratio: %.2f, Change: %.2f, Uptrend: %t\n", ratio, chance, s.radiationUptrend)
}

func solarFlareRoutine(reading *Reading) {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(0)
	for {
		select {
		case <-timer.C:
			reading.updateSolarFlare()
			if reading.SolarFlare == true {
				timer.Reset(10 * time.Second)
			} else {
				timer.Reset(30 * time.Second)
			}
		case <-ticker.C:
			// Ticker to check exit condition
		}
	}
}

func temperatureRoutine(reading *Reading) {
	tickerUpdate := time.NewTicker(1 * time.Second)
	timerTrend := time.NewTimer(0)
	for {
		select {
		case <-tickerUpdate.C:
			reading.updateTemperature()
		case <-timerTrend.C:
			reading.updateTemperatureTrend()
			timerTrend.Reset(time.Duration(rand.Intn(maxTrendSeconds-minTrendSeconds)+minTrendSeconds) * time.Second)
		}
	}
}

func radiationRoutine(reading *Reading) {
	tickerUpdate := time.NewTicker(1 * time.Second)
	timerTrend := time.NewTimer(0)
	for {
		select {
		case <-tickerUpdate.C:
			reading.updateRadiation()
		case <-timerTrend.C:
			reading.updateRadiationTrend()
			timerTrend.Reset(time.Duration(rand.Intn(maxTrendSeconds-minTrendSeconds)+minTrendSeconds) * time.Second)
		}
	}
}
