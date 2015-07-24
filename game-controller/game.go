package main

import (
	"encoding/json"
	"fmt"
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
	shield    chan ShieldRequest
}

// Team contains information about a team
type Team struct {
	Name   string `json:"name"`
	token  string
	Energy int64 `json:"energy"`
	Life   int64 `json:"life"`
	Shield bool  `json:"shield"`
}

// GameRequest is used to interact with the game controller and get a reply back
type GameRequest struct {
	Response chan bool
}

// JoinResponse is used when a team wants to join the game
type JoinResponse struct {
	success bool
	token   string
	message string
}

// JoinRequest is used when a team wants to join the game
type JoinRequest struct {
	Response chan JoinResponse
	name     string
}

//ShieldRequest is used to enable/disable shield
type ShieldRequest struct {
	Response chan bool
	enable   bool
	token    string
}

var game = GameInfo{
	Running: false,
	start:   make(chan GameRequest),
	stop:    make(chan GameRequest),
	join:    make(chan JoinRequest),
	shield:  make(chan ShieldRequest),
	exit:    make(chan []byte),
	Reading: Reading{SolarFlare: initialSolarFlare, Temperature: initialTemperature, Radiation: initialRadiation},
	Teams:   []Team{},
}

func (game *GameInfo) run() {
	var wg sync.WaitGroup
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case req := <-game.start:
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
		case req := <-game.stop:
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
		case req := <-game.join:
			res := JoinResponse{success: false}
			if game.Running {
				res.message = fmt.Sprintf("Team '%s' cannot join the game because it's already running", req.name)
				log.Println(res.message)
			} else {
				if teamExists(game.Teams, req.name) {
					res.message = fmt.Sprintf("Team '%s' already exists.", req.name)
					log.Println(res.message)
				} else {
					team := Team{Name: req.name, Life: initialLife, Energy: initialEnergy, Shield: false, token: randToken()}
					game.Teams = append(game.Teams, team)
					res.success = true
					res.token = team.token
					res.message = fmt.Sprintf("Team '%s' joined the game", req.name)
					log.Printf(res.message)
				}
			}
			req.Response <- res
			close(req.Response)
		case req := <-game.shield:
			i, ok := getTeamIndexByToken(game.Teams, req.token)
			if ok {
				game.Teams[i].Shield = req.enable
				log.Printf("Team '%s' set shield to %t\n", game.Teams[i].Name, game.Teams[i].Shield)
				req.Response <- true
			} else {
				log.Println("Invalid token")
				req.Response <- false
			}
			close(req.Response)
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
