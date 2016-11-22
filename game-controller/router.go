package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/codedellemc/mars-challenge/websocket/wsblaster"
)

// NewRouter sets the routes to the web service
func NewRouter(b *wsblaster.Blaster) *mux.Router {

	// And route for '/ws' explicitly since we need access to 'b'
	routes = append(routes, Route{
		"Websocket",
		"GET",
		"/ws",
		b.GetWSHandler(),
	},
	)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
