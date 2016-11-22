package wswriter

import (
	"errors"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

// WSWriter holds info about the WebSocket writer
type WSWriter struct {
	conn *websocket.Conn
	Exit chan bool
}

// GetWSWriter returns a new WSWriter
func GetWSWriter(u *url.URL) (*WSWriter, error) {
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, errors.New("unable to connect")
	}

	ch := make(chan bool)

	go func() {
		defer c.Close()
		defer close(ch)
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				ch <- true
				log.Printf("error! %s", err)
				return
			}
		}
	}()

	return &WSWriter{
		conn: c,
		Exit: ch,
	}, nil
}

// Write writes the message to the websocket
func (w *WSWriter) Write(m []byte) error {
	if w.conn == nil {
		log.Print("Lost WS connection")
		return errors.New("Lost connection")
	}
	err := w.conn.WriteMessage(websocket.TextMessage, m)
	if err != nil {
		return err
	}
	return nil
}
