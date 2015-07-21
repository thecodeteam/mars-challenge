package main

import (
	"encoding/json"
	"log"
	"time"
)

const (
	initialEnergy = 100
	initialLife   = 100
)

// GameInfo contains information about the state of the game
type GameInfo struct {
	running   bool
	startedAt time.Time
	reading   Reading
	teams     Teams
	start     chan []byte
	end       chan []byte
	exit      chan []byte
}

// Team contains information about a team
type Team struct {
	name   string
	token  string
	energy int64
	life   int64
}

// Teams list
type Teams []Team

var game = GameInfo{
	running: false,
	start:   make(chan []byte),
	end:     make(chan []byte),
	exit:    make(chan []byte),
	reading: Reading{SolarFlare: false, Temperature: 30.0, Radiation: 50},
}

func (game *GameInfo) run() {
	timer := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-game.start:
			if !game.running {
				game.running = true
				game.startedAt = time.Now()
				log.Println("Game started!")
			} else {
				log.Println("Game is already started, not doing anything...")
			}
		case <-game.end:
			if game.running {
				game.running = false
				log.Println("Game stopped!")
			} else {
				log.Println("Game is already stopped, not doing anything...")
			}
		case <-timer.C:
			m, err := json.Marshal(&game)
			if err != nil {
				log.Println("Error parsing to JSON.", err)
				continue
			}
			log.Println(string(m))
			h.broadcast <- []byte(m)

		case <-game.exit:
			log.Println("Exiting game")
			return
		}
	}
}

func (game *GameInfo) getReadings() {
	go solarFlareRoutine(&game.reading)
	go temperatureRoutine(&game.reading)
	radiationRoutine(&game.reading)
}
