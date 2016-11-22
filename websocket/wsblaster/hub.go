// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wsblaster

// Hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Whether to run the writePump routine
	write bool

	// Whether to run the readPump routine
	read bool

	// Buffer holding messages from Clients
	ReadBuffer *MessageBuffer
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.clients, c)
				}
			}
		}
	}
}
