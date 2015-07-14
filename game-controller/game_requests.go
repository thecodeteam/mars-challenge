package main

// GameRequest is used to interact with the game controller and get a reply back
type GameRequest struct {
	Response chan GameResponse
}

// GameResponse is used to respond to game requests
type GameResponse struct {
	success bool
	message string
}

// TokenRequest is used to access token-protected resources
type TokenRequest struct {
	GameRequest
	token string
}

// JoinRequest is used when a team wants to join the game
type JoinRequest struct {
	GameRequest
	name string
}

// KickRequest is used when a team wants to leave the game
type KickRequest struct {
	TokenRequest
	name string
}

//ShieldRequest is used to enable/disable shield
type ShieldRequest struct {
	TokenRequest
	enable bool
}

// ReadingsRequest is used to provide readings from an external source
type ReadingsRequest struct {
	TokenRequest
	readings Reading
}
