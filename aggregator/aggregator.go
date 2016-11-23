package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/viper"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
	"github.com/codedellemc/mars-challenge/sensorsuite/types"
	"github.com/codedellemc/mars-challenge/websocket/wsblaster"
)

// Message holds the final sensor information
type Message struct {
	Solarflare  bool           `json:"solarFlare"`
	Temperature float64        `json:"temperature"`
	Radiation   int32          `json:"radiation"`
	Stamp       types.JSONTime `json:"stamp"`
	sync.RWMutex
}

// UpdateSolarFlare sets the Solar value in the struct
func (m *Message) UpdateSolarFlare(b bool) {
	m.Lock()
	defer m.Unlock()
	m.Solarflare = b
}

func main() {
	viper.SetDefault("listen_address", ss.DefaultAggWSAddr)
	viper.SetDefault("sensor_flare_address", "localhost"+ss.DefaultFlareWSAddr)
	viper.SetDefault("gc_address", "localhost:8080")
	viper.SetDefault("post_gc", true)
	viper.SetDefault("admin_token", "1234")
	viper.AutomaticEnv()

	addr := viper.GetString("listen_address")
	blaster := wsblaster.GetBlaster(&addr, true)
	go blaster.Run()

	reading := &Message{}

	flareExit := make(chan bool)
	flareAddr := viper.GetString("sensor_flare_address")
	go types.FlareUpdateRoutine(reading, &flareAddr, flareExit)

	doPost := viper.GetBool("post_gc")
	gcAddr := "http://" + viper.GetString("gc_address") + "/api/readings"
	gcToken := viper.GetString("admin_token")

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			ok := averageReadings(reading, blaster)
			if !ok {
				break
			}
			reading.RLock()
			m, _ := json.Marshal(reading)
			reading.RUnlock()
			blaster.Write(m)
			if doPost {
				go postReading(m, &gcAddr, &gcToken)
			}
			// Go read all the sensor values and average them
			// POST values to Game Controller
		case <-flareExit:
			log.Fatal("Unable to connect to Solar Flare source")
		}
	}
}

// SensorMessage holds incoming sensor data from the publishers
type SensorMessage struct {
	Temperature float64 `json:"temperature"`
	Radiation   int     `json:"radiation"`
}

func averageReadings(m *Message, b *wsblaster.Blaster) bool {
	messages := b.GetReadBuffer()

	messages.Lock()
	defer messages.Unlock()

	numMessages := len(messages.Messages)
	log.Printf("received %d messages\n", numMessages)
	var avgTemp float64
	var avgRad int32

	count := 0
	for _, msg := range messages.Messages {
		reading := &SensorMessage{}
		err := json.Unmarshal(*msg, reading)
		if err != nil {
			log.Println("Unable to unmarshal sensor data")
			log.Printf("%s\n", err)
			continue
		}

		avgTemp += reading.Temperature
		avgRad += int32(reading.Radiation)
		count++
	}
	messages.Messages = nil

	if count == 0 {
		return false
	}

	m.Lock()
	defer m.Unlock()
	m.Radiation = avgRad / int32(count)
	m.Temperature = avgTemp / float64(count)
	log.Printf("Averages now: %d, %f", m.Radiation, m.Temperature)
	return true
}

func postReading(m []byte, controllerEndpoint *string, token *string) {

	req, err := http.NewRequest(
		"POST", *controllerEndpoint, bytes.NewBuffer(m))
	req.Header.Set("X-Auth-Token", *token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Unable to POST readings to game controller")
		log.Printf("err: %v\n", err)
		return
	}
	defer resp.Body.Close()

}
