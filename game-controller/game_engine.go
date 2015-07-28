package main

import (
	"log"
	"sync"
	"time"
)

func (game *GameInfo) runEngine(wg *sync.WaitGroup) {
	wg.Add(len(game.Teams))
	for i := range game.Teams {
		go game.handleTeam(&game.Teams[i], wg)
	}
}

func (game *GameInfo) handleTeam(team *Team, wg *sync.WaitGroup) {
	for team.Life > 0 && game.Running {
		time.Sleep(1 * time.Second)

		if team.Energy <= 0 {
			team.Shield = false
		}

		if team.Shield {
			team.Energy--
			continue
		}

		team.Life--
	}
	log.Println("Exiting team", team.Name)
	wg.Done()
}
