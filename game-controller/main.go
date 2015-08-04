package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	adminToken := os.Getenv("ADMIN_TOKEN")
	if len(adminToken) <= 0 {
		adminToken = randToken()
		log.Printf("Admin token not defined. Using '%s'", adminToken)
	}

	autoReadings, err := strconv.ParseBool(os.Getenv("AUTO_READINGS"))
	if err != nil {
		autoReadings = true
	}
	game.autoReadings = autoReadings
	log.Printf("Automatic readings set to '%t'", autoReadings)

	go h.run()
	go game.run(adminToken)

	router := NewRouter()

	err = http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
