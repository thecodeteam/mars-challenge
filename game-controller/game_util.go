package main

import (
	"crypto/rand"
	"fmt"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
)

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (reading *Reading) validate() (bool, string) {
	if reading.Temperature < ss.MinTemp ||
		reading.Temperature > ss.MaxTemp {
		return false,
			fmt.Sprintf(
				"Temperature not within valid range [%.2f, %.2f]",
				ss.MinTemp, ss.MaxTemp)
	}
	if reading.Radiation < ss.MinRadiation ||
		reading.Radiation > ss.MaxRadiation {
		return false,
			fmt.Sprintf(
				"Radiation not within valid range [%d, %d]",
				ss.MinRadiation, ss.MaxRadiation)
	}
	return true, ""
}

func (game *GameInfo) teamExists(name string) bool {
	for _, v := range game.Teams {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (game *GameInfo) getTeamIndex(name string) (int, bool) {
	for i, v := range game.Teams {
		if v.Name == name {
			return i, true
		}
	}
	return -1, false
}

func (game *GameInfo) authorizeTeam(token string) (int, bool) {
	for i, v := range game.Teams {
		if v.token == token {
			return i, true
		}
	}
	return -1, false
}

func (game *GameInfo) authorizeAdmin(token string) bool {
	return game.adminToken == token
}

func (game *GameInfo) isOver() bool {
	remainingTeams := 0
	for _, v := range game.Teams {
		if v.Life > 0 {
			remainingTeams++
		}
	}
	return remainingTeams <= 1
}
