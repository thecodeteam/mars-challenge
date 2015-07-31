package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func serveWs(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	go c.writePump()
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
