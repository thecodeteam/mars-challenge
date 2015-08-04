package main

import (
	"encoding/json"
	"log"
	"time"
)

const (
	initialTemperature = -53.5
	initialRadiation   = 500
	initialSolarFlare  = false
	maxTemperature     = 35.00
	minTemperature     = -142.00
	maxRadiation       = 1000
	minRadiation       = 0
)

var reading = Reading{
	SolarFlare:  initialSolarFlare,
	Temperature: initialTemperature,
	Radiation:   initialRadiation,
}

func (reading *Reading) run() {
	go reading.getReadings()
	ticker := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-ticker.C:
			m, err := json.Marshal(&reading)
			if err != nil {
				log.Println("Error parsing to JSON.", err)
				continue
			}
			log.Println(string(m))
			h.broadcast <- []byte(m)
		}
	}
}

func (reading *Reading) getReadings() {
	go solarFlareRoutine(reading)
	go temperatureRoutine(reading)
	radiationRoutine(reading)
}
