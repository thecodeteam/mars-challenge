package main

import (
	"encoding/json"
	"log"
	"sync"
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
	// Running defines whether the game is running or not
	Running   bool      `json:"running"`
	StartedAt time.Time `json:"startedAt"`
	Reading   Reading   `json:"readings"`
	Teams     []Team    `json:"teams"`
	start     chan GameRequest
	stop      chan GameRequest
	join      chan JoinRequest
	exit      chan []byte
}

// Team contains information about a team
type Team struct {
	Name   string `json:"name"`
	token  string
	Energy int64 `json:"energy"`
	Life   int64 `json:"life"`
}

// GameRequest is used to interact with the game controller and get a reply back
type GameRequest struct {
	Response chan bool
}

// JoinResponse is used when a team wants to join the game
type JoinResponse struct {
	success bool
	token   string
}

// JoinRequest is used when a team wants to join the game
type JoinRequest struct {
	Response chan JoinResponse
	name     string
}

var game = GameInfo{
	Running: false,
	start:   make(chan GameRequest),
	stop:    make(chan GameRequest),
	join:    make(chan JoinRequest),
	exit:    make(chan []byte),
	Reading: Reading{SolarFlare: initialSolarFlare, Temperature: initialTemperature, Radiation: initialRadiation},
	Teams:   []Team{},
}

func (game *GameInfo) run() {
	var wg sync.WaitGroup
	var req GameRequest
	var joinReq JoinRequest
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case req = <-game.start:
			if !game.Running {
				game.Running = true
				game.StartedAt = time.Now()
				wg.Add(3)
				go game.getReadings(&wg)
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
				wg.Wait()
				req.Response <- true
				log.Println("Game stopped!")
			} else {
				req.Response <- false
				log.Println("Game is already stopped, not doing anything...")
			}
			close(req.Response)
		case joinReq = <-game.join:
			if game.Running {
				joinReq.Response <- JoinResponse{success: false}
				log.Printf("Team '%s' cannot join the game because it's already running\n", joinReq.name)
			} else {
				if teamExists(game.Teams, joinReq.name) {
					joinReq.Response <- JoinResponse{success: false}
					log.Printf("Team '%s' already exists.\n", joinReq.name)
				} else {
					team := Team{Name: joinReq.name, Life: initialLife, Energy: initialEnergy, token: randToken()}
					game.Teams = append(game.Teams, team)
					joinReq.Response <- JoinResponse{success: true, token: team.token}
					log.Printf("Team '%s' joined the game\n", joinReq.name)
				}
			}
			close(joinReq.Response)
		case <-ticker.C:
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

func (game *GameInfo) getReadings(wg *sync.WaitGroup) {
	go solarFlareRoutine(wg, game)
	go temperatureRoutine(wg, game)
	radiationRoutine(wg, game)
}
