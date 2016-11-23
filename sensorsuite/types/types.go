package types

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/codedellemc/mars-challenge/sensorsuite/utils"
	"github.com/codedellemc/mars-challenge/websocket/wsreader"
)

//SensorSuiteReading contains the combined readings from the sensors
type SensorSuiteReading struct {
	Temperature float64  `json:"temperature"`
	Radiation   int      `json:"radiation"`
	Stamp       JSONTime `json:"stamp"`
	sync.RWMutex
}

//JSONTime holds a timestamp that can be marshaled to JSON
type JSONTime time.Time

//MarshalJSON coverts the timestamp into a JSON format
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Now().Format("2006-01-02T15:04:05Z07:00"))
	return []byte(stamp), nil
}

// FlareReading contains SolarFlare readings
type FlareReading struct {
	SolarFlare bool     `json:"solarFlare"`
	Stamp      JSONTime `json:"stamp"`
	sync.RWMutex
}

// TemperatureReading contains Temperature sensor readings
type TemperatureReading struct {
	Temperature        float64  `json:"temperature"`
	Stamp              JSONTime `json:"stamp"`
	solarFlare         bool
	temperatureUptrend bool
	sync.RWMutex
}

// RadiationReading contains Radiation sensor readings
type RadiationReading struct {
	Radiation        int      `json:"radiation"`
	Stamp            JSONTime `json:"stamp"`
	solarFlare       bool
	radiationUptrend bool
	sync.RWMutex
}

// UpdateSolarFlare sets a new random solarflare flag
func (f *FlareReading) UpdateSolarFlare() {
	f.Lock()
	defer f.Unlock()
	f.SolarFlare = utils.GetNewFlare()
	fmt.Println("solarflare is: ", f.SolarFlare)
}

//UpdateTemperature sets a new random temperature
func (t *TemperatureReading) UpdateTemperature() {
	t.Lock()
	defer t.Unlock()
	t.Temperature = utils.GetNewTemp(t.Temperature, t.temperatureUptrend)
}

//UpdateTemperatureTrend sets a new temperature trend value
func (t *TemperatureReading) UpdateTemperatureTrend() {
	t.Lock()
	defer t.Unlock()
	t.temperatureUptrend = utils.GetNewTempTrend(
		t.Temperature, t.solarFlare)
}

//UpdateSolarFlare sets the solar flare value in the TemperatureReading
func (t *TemperatureReading) UpdateSolarFlare(b bool) {
	t.Lock()
	defer t.Unlock()
	t.solarFlare = b
}

//UpdateRadiation sets a new random Radiation
func (r *RadiationReading) UpdateRadiation() {
	r.Lock()
	defer r.Unlock()
	r.Radiation = utils.GetNewRadiation(r.Radiation, r.radiationUptrend)
}

//UpdateRadiationTrend sets a new Radiation trend value
func (r *RadiationReading) UpdateRadiationTrend() {
	r.Lock()
	defer r.Unlock()
	r.radiationUptrend = utils.GetNewRadiationTrend(
		r.Radiation, r.solarFlare)
}

//UpdateSolarFlare sets the solar flare value in the RadiationReading
func (r *RadiationReading) UpdateSolarFlare(b bool) {
	r.Lock()
	defer r.Unlock()
	r.solarFlare = b
}

//FlareUpdater have a method for updating solarflare value
type FlareUpdater interface {
	UpdateSolarFlare(b bool)
}

//FlareUpdateRoutine will update a flare value from the websocket
func FlareUpdateRoutine(f FlareUpdater, addr *string, exit chan bool) {
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
		sf := &FlareReading{}
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
