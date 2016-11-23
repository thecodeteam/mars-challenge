package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	ss "github.com/codedellemc/mars-challenge/sensorsuite"
)

var homeTempl = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func serveAPIConfig(w http.ResponseWriter, r *http.Request) {

	config := struct {
		MaxTemperature float64 `json:"maxTemperature"`
		MinTemperature float64 `json:"minTemperature"`
		MaxRadiation   float64 `json:"maxRadiation"`
		MinRadiation   float64 `json:"minRadiation"`
		AutoReadings   bool    `json:"autoReadings"`
	}{
		ss.MaxTemp,
		ss.MinTemp,
		ss.MaxRadiation,
		ss.MinRadiation,
		game.autoReadings,
	}

	m, err := json.Marshal(&config)
	if err != nil {
		http.Error(w, "Error parsing JSON", 400)
		return
	}

	w.Write([]byte(m))
}

func serveAPIReadings(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Auth-Token")
	if len(token) == 0 {
		http.Error(w, "No auth token present", 400)
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error processing the body", 400)
		return
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	var reading Reading
	dec.Decode(&reading)

	req := ReadingsRequest{TokenRequest: TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}, readings: reading}
	game.readings <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte(res.message))
	} else {
		http.Error(w, res.message, 400)
	}
}

func serveAPIStart(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Auth-Token")
	if len(token) == 0 {
		http.Error(w, "No auth token present", 400)
		return
	}

	req := TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}
	game.start <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte(res.message))
	} else {
		http.Error(w, res.message, 400)
	}
}

func serveAPIStop(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Auth-Token")
	if len(token) == 0 {
		http.Error(w, "No auth token present", 400)
		return
	}

	req := TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}
	game.stop <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte(res.message))
	} else {
		http.Error(w, res.message, 400)
	}
}

func serveAPIReset(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Auth-Token")
	if len(token) == 0 {
		http.Error(w, "No auth token present", 400)
		return
	}

	req := TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}
	game.reset <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte(res.message))
	} else {
		http.Error(w, res.message, 400)
	}
}

func serveAPIJoin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	req := JoinRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, name: name}
	game.join <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte(res.message))
	} else {
		http.Error(w, res.message, 400)
	}
}

func serveAPIKick(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	token := r.Header.Get("X-Auth-Token")
	if len(token) == 0 {
		http.Error(w, "No auth token present", 400)
		return
	}

	req := KickRequest{TokenRequest: TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}, name: name}
	game.kick <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte(res.message))
	} else {
		http.Error(w, res.message, 400)
	}
}

func serveAPIEnableShield(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Auth-Token")
	if len(token) == 0 {
		http.Error(w, "No auth token present", 400)
		return
	}
	req := ShieldRequest{TokenRequest: TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}, enable: true}
	game.shield <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte("Shield successfully enabled"))
	} else {
		http.Error(w, "Could not enable shield", 400)
	}
}

func serveAPIDisableShield(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Auth-Token")
	if len(token) == 0 {
		http.Error(w, "No auth token present", 400)
		return
	}
	req := ShieldRequest{TokenRequest: TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}, enable: false}
	game.shield <- req
	res := <-req.Response
	if res.success {
		w.Write([]byte("Shield successfully disabled"))
	} else {
		http.Error(w, "Could not disable shield", 400)
	}
}
