package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/spf13/viper"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
	"github.com/codedellemc/mars-challenge/sensorsuite/types"
	"github.com/codedellemc/mars-challenge/websocket/wsblaster"
)

const (
	defaultAddr = ":9000"
)

func solarFlareRoutine(r *types.FlareReading) {
	timer := time.NewTimer(0)

	for {
		select {
		case <-timer.C:
			r.UpdateSolarFlare()
			if r.SolarFlare {
				timer.Reset(10 * time.Second)
			} else {
				timer.Reset(30 * time.Second)
			}
		}
	}
}

func main() {
	viper.SetEnvPrefix("SENSOR")
	viper.SetDefault("listen_address", ss.DefaultFlareWSAddr)
	viper.AutomaticEnv()

	rand.Seed(time.Now().UTC().UnixNano())
	reading := &types.FlareReading{
		SolarFlare: ss.InitSolarFlare,
	}

	go solarFlareRoutine(reading)

	addr := viper.GetString("listen_address")
	blaster := wsblaster.GetBlaster(&addr, false)
	go blaster.Run()
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			reading.RLock()
			m, _ := json.Marshal(reading)
			reading.RUnlock()
			blaster.Write(m)
		}
	}
}
