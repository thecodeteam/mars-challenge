package main

import (
	"html/template"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
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
	req := GameRequest{Response: make(chan bool)}
	game.start <- req
	res := <-req.Response
	if res {
		w.Write([]byte("Game started"))
	} else {
		http.Error(w, "Game already in progress. Not doing anything", 400)
	}
}

func serveAPIStop(w http.ResponseWriter, r *http.Request) {
	req := GameRequest{Response: make(chan bool)}
	game.stop <- req
	res := <-req.Response
	if res {
		w.Write([]byte("Game stopped"))
	} else {
		http.Error(w, "Game already stopped. Not doing anything", 400)
	}
}
