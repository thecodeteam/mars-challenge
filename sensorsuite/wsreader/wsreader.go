package wsreader

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

//WSReader holds info about the WebSocker Reader
type WSReader struct {
	Interrupt chan bool
	Exit      chan bool
	C         chan []byte
	conn      *websocket.Conn
}

//GetWSReader returns a new WSReader
func GetWSReader(url *string) (*WSReader, error) {

	log.Printf("connecting to %s\n", *url)

	c, _, err := websocket.DefaultDialer.Dial(*url, nil)
	if err != nil {
		return nil, errors.New("unable to connect")
	}

	reader := &WSReader{
		Interrupt: make(chan bool),
		Exit:      make(chan bool),
		C:         make(chan []byte),
		conn:      c,
	}

	return reader, nil
}

func (r *WSReader) readLoop() {
	done := make(chan struct{})

	go func() {
		defer r.conn.Close()
		defer close(done)

		for {
			_, message, err := r.conn.ReadMessage()
			if err != nil {
				r.Interrupt <- true
				return
			}
			r.C <- message
		}
	}()

	for {
		select {
		case <-r.Interrupt:
			err := r.conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseNormalClosure, ""),
			)
			if err != nil {
				r.Exit <- true
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			r.conn.Close()
			r.Exit <- true
			return
		}

	}
}

//Run starts a go routine to constantly read messages
func (r *WSReader) Run() {
	go r.readLoop()
}
