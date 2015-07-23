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

func teamExists(team []Team, name string) bool {
	for _, v := range team {
		if v.Name == name {
			return true
		}
	}
	return false
}
