package main

import (
	"log"
	"math"
	"sync"
	"time"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
)

const (
	maxEnergyGain = 5
	maxEnergyLoss = 5
	maxLifeLoss   = 5
)

func (game *GameInfo) runEngine(wg *sync.WaitGroup) {
	wg.Add(len(game.Teams))
	for i := range game.Teams {
		go game.handleTeam(&game.Teams[i], wg)
	}
}

func (game *GameInfo) handleTeam(team *Team, wg *sync.WaitGroup) {
	var temperatureRatio float64
	var radiationRatio float64
	var energyGain float64
	var energyLoss float64
	var lifeLoss float64

	for team.Life > 0 && game.Running {
		time.Sleep(1 * time.Second)
		radiationRatio = (float64)(
			game.Reading.Radiation-ss.MinRadiation) / (float64)(
			ss.MaxRadiation-ss.MinRadiation)

		energyLoss = radiationRatio * maxEnergyLoss
		if float64(team.Energy)-energyLoss <= 0 {
			team.Shield = false
		}

		if team.Shield {
			team.Energy = int64(math.Max(float64(team.Energy)-math.Ceil(energyLoss), 0))
			log.Printf("Team %s: Energy -%.2f\n", team.Name, energyLoss)
			continue
		}

		lifeLoss = radiationRatio * maxLifeLoss
		team.Life = int64(math.Max(float64(team.Life)-math.Ceil(lifeLoss), 0))

		temperatureRatio = (game.Reading.Temperature - ss.MinTemp) /
			(ss.MaxTemp - ss.MinTemp)
		energyGain = temperatureRatio * maxEnergyGain
		team.Energy = int64(math.Min(float64(team.Energy)+math.Ceil(energyGain), 100))

		log.Printf("Team %s: Life -%.2f, Energy +%.2f\n", team.Name, lifeLoss, energyGain)
	}

	log.Println("Exiting goroutine for team", team.Name)
	wg.Done()
}
