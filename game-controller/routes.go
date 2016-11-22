package main

import "net/http"

// Route represents the structure of a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes represents the available routes for the web service
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		serveHome,
	},
	Route{
		"Start Game",
		"POST",
		"/api/start",
		serveAPIStart,
	},
	Route{
		"Stop Game",
		"POST",
		"/api/stop",
		serveAPIStop,
	},
	Route{
		"Reset Game",
		"POST",
		"/api/reset",
		serveAPIReset,
	},
	Route{
		"Join Game",
		"POST",
		"/api/join/{name}",
		serveAPIJoin,
	},
	Route{
		"Kick Team",
		"POST",
		"/api/kick/{name}",
		serveAPIKick,
	},
	Route{
		"Enable Shield",
		"POST",
		"/api/shield/enable",
		serveAPIEnableShield,
	},
	Route{
		"Disable Shield",
		"POST",
		"/api/shield/disable",
		serveAPIDisableShield,
	},
	Route{
		"Input readings",
		"POST",
		"/api/readings",
		serveAPIReadings,
	},
	Route{
		"Get Game Config",
		"GET",
		"/api/config",
		serveAPIConfig,
	},
}
