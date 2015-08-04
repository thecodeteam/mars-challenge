package main

import (
	"crypto/rand"
	"fmt"
)

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (reading *Reading) validate() (bool, string) {
	if reading.Temperature < minTemperature || reading.Temperature > maxTemperature {
		return false, fmt.Sprintf("Temperature not within valid range [%.2f, %.2f]", minTemperature, maxTemperature)
	}
	if reading.Radiation < minRadiation || reading.Radiation > maxRadiation {
		return false, fmt.Sprintf("Radiation not within valid range [%.2f, %.2f]", minRadiation, maxRadiation)
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
