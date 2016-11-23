package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
	"github.com/codedellemc/mars-challenge/sensorsuite/utils"
)

// Reading contains the current sensor readings
type Reading struct {
	SolarFlare         bool    `json:"solarFlare"`
	Temperature        float64 `json:"temperature"`
	Radiation          int     `json:"radiation"`
	temperatureUptrend bool
	radiationUptrend   bool
}

func (s *Reading) updateSolarFlare() {
	s.SolarFlare = utils.GetNewFlare()
}

func (s *Reading) updateTemperature() {
	s.Temperature = utils.GetNewTemp(s.Temperature, s.temperatureUptrend)
}

func (s *Reading) updateTemperatureTrend() {
	s.temperatureUptrend = utils.GetNewTempTrend(
		s.Temperature, s.SolarFlare)
}

func (s *Reading) updateRadiation() {
	s.Radiation = utils.GetNewRadiation(s.Radiation, s.radiationUptrend)
}

func (s *Reading) updateRadiationTrend() {
	s.radiationUptrend = utils.GetNewRadiationTrend(
		s.Radiation, s.SolarFlare)
}

func solarFlareRoutine(wg *sync.WaitGroup, game *GameInfo) {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(0)
	for game.Running {
		select {
		case <-timer.C:
			game.Reading.updateSolarFlare()
			if game.Reading.SolarFlare == true {
				timer.Reset(10 * time.Second)
			} else {
				timer.Reset(30 * time.Second)
			}
		case <-ticker.C:
			// Ticker to check exit condition
		}
	}
	log.Println("Exiting solar flare goroutine")
	wg.Done()
}

func temperatureRoutine(wg *sync.WaitGroup, game *GameInfo) {
	tickerUpdate := time.NewTicker(1 * time.Second)
	timerTrend := time.NewTimer(0)
	for game.Running {
		select {
		case <-tickerUpdate.C:
			game.Reading.updateTemperature()
		case <-timerTrend.C:
			game.Reading.updateTemperatureTrend()
			timerTrend.Reset(
				time.Duration(rand.Intn(
					ss.MaxTrendSec-ss.MinTrendSec)+
					ss.MinTrendSec) * time.Second)
		}
	}
	log.Println("Exiting temperature goroutine")
	wg.Done()
}

func radiationRoutine(wg *sync.WaitGroup, game *GameInfo) {
	tickerUpdate := time.NewTicker(1 * time.Second)
	timerTrend := time.NewTimer(0)
	for game.Running {
		select {
		case <-tickerUpdate.C:
			game.Reading.updateRadiation()
		case <-timerTrend.C:
			game.Reading.updateRadiationTrend()
			timerTrend.Reset(
				time.Duration(rand.Intn(
					ss.MaxTrendSec-ss.MinTrendSec)+
					ss.MinTrendSec) * time.Second)
		}
	}
	log.Println("Exiting radiation goroutine")
	wg.Done()
}
