package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/spf13/viper"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
	"github.com/codedellemc/mars-challenge/sensorsuite/types"
	"github.com/codedellemc/mars-challenge/websocket/wsblaster"
)

func temperatureRoutine(r *types.TemperatureReading) {
	tickerUpdate := time.NewTicker(1 * time.Second)
	timerTrend := time.NewTimer(0)

	for {
		select {
		case <-tickerUpdate.C:
			r.UpdateTemperature()
		case <-timerTrend.C:
			r.UpdateTemperatureTrend()
			timerTrend.Reset(time.Duration(rand.Intn(
				ss.MaxTrendSec-ss.MinTrendSec+1)+
				ss.MinTrendSec) * time.Second)
		}
	}
}

func main() {
	viper.SetEnvPrefix("SENSOR")
	viper.SetDefault("listen_address", ss.DefaultTempWSAddr)
	viper.SetDefault("flare_address", "localhost"+ss.DefaultFlareWSAddr)
	viper.AutomaticEnv()

	rand.Seed(time.Now().UTC().UnixNano())
	reading := &types.TemperatureReading{
		Temperature: ss.InitTemp,
	}

	flareExit := make(chan bool)
	flareAddr := viper.GetString("flare_address")
	go types.FlareUpdateRoutine(reading, &flareAddr, flareExit)

	go temperatureRoutine(reading)

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
		case <-flareExit:
			log.Fatal("Unable to connect to Solar Flare source")
		}
	}
}
