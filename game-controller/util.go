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

func teamExists(teams []Team, name string) bool {
	for _, v := range teams {
		if v.Name == name {
			return true
		}
	}
	return false
}

func getTeamIndexByToken(teams []Team, token string) (int, bool) {
	for i, v := range teams {
		if v.token == token {
			return i, true
		}
	}
	return -1, false
}
