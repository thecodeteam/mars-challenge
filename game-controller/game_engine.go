package main

import "time"

func (game *GameInfo) runEngine() {
	for i := range game.Teams {
		go handleTeam(&game.Teams[i])
	}
}

func handleTeam(team *Team) {
	for team.Life > 0 {
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
}
