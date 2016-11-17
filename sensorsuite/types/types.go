package types

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
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
	x := rand.Intn(2)
	f.Lock()
	defer f.Unlock()
	if x != 0 {
		f.SolarFlare = true
	} else {
		f.SolarFlare = false
	}
	fmt.Println("solarflare is: ", f.SolarFlare)
}

//UpdateTemperature sets a new random temperature
func (t *TemperatureReading) UpdateTemperature() {
	var min float64
	var max float64

	if t.temperatureUptrend {
		max = t.Temperature + ss.VariationTemp
		min = t.Temperature
	} else {
		max = t.Temperature
		min = t.Temperature - ss.VariationTemp
	}

	temperature := (rand.Float64() * (max - min)) + min
	if temperature < ss.MinTemp {
		temperature = ss.MinTemp
	} else if temperature > ss.MaxTemp {
		temperature = ss.MaxTemp
	}
	t.Lock()
	defer t.Unlock()
	t.Temperature = temperature
}

//UpdateTemperatureTrend sets a new temperature trend value
func (t *TemperatureReading) UpdateTemperatureTrend() {
	ratio := (t.Temperature - ss.MinTemp) / (ss.MaxTemp - ss.MinTemp)
	chance := rand.Float64()
	t.Lock()
	defer t.Unlock()
	t.temperatureUptrend = chance > ratio || t.solarFlare
}

//UpdateSolarFlare sets the solar flare value in the TemperatureReading
func (t *TemperatureReading) UpdateSolarFlare(b bool) {
	t.Lock()
	defer t.Unlock()
	t.solarFlare = b
}

//UpdateRadiation sets a new random Radiation
func (r *RadiationReading) UpdateRadiation() {
	var min int
	var max int

	if r.radiationUptrend {
		max = r.Radiation + ss.VariationRadiation
		min = r.Radiation
	} else {
		max = r.Radiation
		min = r.Radiation - ss.VariationRadiation
	}

	radiation := rand.Intn(max-min) + min
	if radiation < ss.MinRadiation {
		radiation = ss.MinRadiation
	} else if radiation > ss.MaxRadiation {
		radiation = ss.MaxRadiation
	}
	r.Lock()
	defer r.Unlock()
	r.Radiation = radiation
}

//UpdateRadiationTrend sets a new Radiation trend value
func (r *RadiationReading) UpdateRadiationTrend() {
	ratio := (float64)(r.Radiation-ss.MinRadiation) /
		(float64)(ss.MaxRadiation-ss.MinRadiation)
	chance := rand.Float64()
	r.Lock()
	defer r.Unlock()
	r.radiationUptrend = chance > ratio || r.solarFlare
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
