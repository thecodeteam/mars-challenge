package main

import (
	"encoding/json"
	"log"
	"time"
)

const (
	initialEnergy      = 100
	initialLife        = 100
	initialTemperature = 0
	initialRadiation   = 500
	initialSolarFlare  = false
)

// GameInfo contains information about the state of the game
type GameInfo struct {
	Running   bool      `json:"running"`
	StartedAt time.Time `json:"startedAt"`
	Reading   Reading   `json:"readings"`
	Teams     []Team    `json:"teams"`
	start     chan GameRequest
	stop      chan GameRequest
	exit      chan []byte
}

// Team contains information about a team
type Team struct {
	name   string
	token  string
	energy int64
	life   int64
}

// GameRequest is used to interact with the game controller and get a reply back
type GameRequest struct {
	Response chan bool
}

var game = GameInfo{
	Running: false,
	start:   make(chan GameRequest),
	stop:    make(chan GameRequest),
	exit:    make(chan []byte),
	Reading: Reading{SolarFlare: initialSolarFlare, Temperature: initialTemperature, Radiation: initialRadiation},
	Teams:   make([]Team, 10, 10),
}

func (game *GameInfo) run() {
	var req GameRequest
	ticker := time.NewTicker(time.Second * 1).C
	for {
		select {
		case req = <-game.start:
			if !game.Running {
				game.Running = true
				game.StartedAt = time.Now()
				req.Response <- true
				log.Println("Game started!")
			} else {
				req.Response <- false
				log.Println("Game is already started, not doing anything...")
			}
			close(req.Response)
		case req = <-game.stop:
			if game.Running {
				game.Running = false
				req.Response <- true
				log.Println("Game stopped!")
			} else {
				req.Response <- false
				log.Println("Game is already stopped, not doing anything...")
			}
			close(req.Response)
		case <-ticker:
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
	go solarFlareRoutine(&game.Reading)
	go temperatureRoutine(&game.Reading)
	radiationRoutine(&game.Reading)
}
