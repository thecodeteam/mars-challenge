package utils

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/codedellemc/mars-challenge/sensorsuite/types"
	"github.com/codedellemc/mars-challenge/sensorsuite/wsreader"
)

//FlareUpdateRoutine will update a flare value from the websocket
func FlareUpdateRoutine(f types.FlareUpdater, addr *string, exit chan bool) {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	uStr := u.String()
	reader, err := wsreader.GetWSReader(&uStr)
	if err != nil {
		log.Printf("%s", err)
		exit <- true
		return
	}
	reader.Run()
	for {
		sf := &types.FlareReading{}
		select {
		case m := <-reader.C:
			json.Unmarshal(m, sf)
			f.UpdateSolarFlare(sf.SolarFlare)
		case <-reader.Exit:
			log.Printf("Lost connection to Flare websocket")
			exit <- true
			return
		}
	}
}
