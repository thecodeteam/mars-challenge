package main

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/spf13/viper"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
	"github.com/codedellemc/mars-challenge/sensorsuite/types"
	"github.com/codedellemc/mars-challenge/websocket/wsblaster"
	"github.com/codedellemc/mars-challenge/websocket/wsreader"
	"github.com/codedellemc/mars-challenge/websocket/wswriter"
)

func sensorUpdateRoutine(r *types.SensorSuiteReading, c chan bool) {
	tempAddr := viper.GetString("temperature_address")
	radAddr := viper.GetString("radiation_address")
	addr := viper.GetString("listen_address")

	tu := url.URL{Scheme: "ws", Host: tempAddr, Path: "/ws"}
	ru := url.URL{Scheme: "ws", Host: radAddr, Path: "/ws"}
	tuStr := tu.String()
	ruStr := ru.String()
	tempReader, err := wsreader.GetWSReader(&tuStr)
	if err != nil {
		log.Printf("%s", err)
		c <- true
		return
	}
	radReader, err := wsreader.GetWSReader(&ruStr)
	if err != nil {
		log.Printf("%s", err)
		c <- true
		return
	}
	tempReader.Run()
	radReader.Run()

	blaster := wsblaster.GetBlaster(&addr, false)
	go blaster.Run()

	tr := &types.TemperatureReading{}
	rr := &types.RadiationReading{}
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case m := <-tempReader.C:
			json.Unmarshal(m, tr)
			r.Lock()
			r.Temperature = tr.Temperature
			r.Unlock()
		case m := <-radReader.C:
			json.Unmarshal(m, rr)
			r.Lock()
			r.Radiation = rr.Radiation
			r.Unlock()
		case <-ticker.C:
			r.RLock()
			m, _ := json.Marshal(r)
			r.RUnlock()
			blaster.Write(m)
		case <-tempReader.Exit:
			log.Printf("Lost connection to Temperature websocket")
			c <- true
			return
		case <-radReader.Exit:
			log.Printf("Lost connection to Radiation websocket")
			c <- true
			return
		}
	}
}

func publishData(r *types.SensorSuiteReading, c chan bool) {
	defer func() {
		c <- true
	}()

	aggAddr := viper.GetString("aggregator_address")
	aggURL := url.URL{Scheme: "ws", Host: aggAddr, Path: "/ws"}
	for {
		writer, err := wswriter.GetWSWriter(&aggURL)
		if err != nil {
			log.Print("Unable to connect to aggregator.")
			log.Print("Will retry in 5 seconds")
			time.Sleep(5 * time.Second)
			continue
		} else {
			log.Print("Connected to aggregator")
		}

		ticker := time.NewTicker(1 * time.Second)
	R:
		for {
			select {
			case <-ticker.C:
				r.RLock()
				m, _ := json.Marshal(r)
				r.RUnlock()
				err := writer.Write(m)
				if err != nil {
					break R
				}
			case <-writer.Exit:
				log.Print("Lost connection to aggregator")
				break R
			}
		}
		ticker.Stop()
		writer = nil
	}
}

func main() {
	viper.SetEnvPrefix("SENSOR")
	viper.SetDefault("listen_address", ss.DefaultPubWSAddr)
	viper.SetDefault("temperature_address", "localhost"+ss.DefaultTempWSAddr)
	viper.SetDefault("radiation_address", "localhost"+ss.DefaultRadWSAddr)
	viper.SetDefault("aggregator_address", "localhost"+ss.DefaultAggWSAddr)
	viper.AutomaticEnv()

	reading := &types.SensorSuiteReading{}

	sensorExit := make(chan bool)
	go sensorUpdateRoutine(reading, sensorExit)

	publishExit := make(chan bool)
	go publishData(reading, publishExit)

	for {
		select {
		case <-sensorExit:
			log.Fatal("Unable to connect to one or more sensors")
		case <-publishExit:
			log.Fatal("Unrecoverable error writing to aggregator")
		}
	}
}
