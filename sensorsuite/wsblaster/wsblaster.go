package wsblaster

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// Blaster contains info about the WebSocket blaster
type Blaster struct {
	h    *Hub
	addr *string
}

//MessageBuffer holds incoming messages from clients
type MessageBuffer struct {
	Messages []*[]byte
	sync.RWMutex
}

// GetBlaster returns a new Blaster
func GetBlaster(addr *string, read bool) *Blaster {
	var buffer *MessageBuffer
	if read {
		buffer = &MessageBuffer{}
	}
	return &Blaster{
		h: &Hub{
			broadcast:  make(chan []byte),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			clients:    make(map[*Client]bool),
			read:       read,
			ReadBuffer: buffer,
		},
		addr: addr,
	}
}

// GetReadBuffer returns pointer to the ReadBuffer
func (b *Blaster) GetReadBuffer() *MessageBuffer {
	return b.h.ReadBuffer
}

//Run starts the HTTP and WS process. Run does not return
func (b *Blaster) Run() {

	go b.h.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(b.h, w, r)
	})

	fmt.Printf("Listening on: %s\n", *b.addr)
	log.Fatal(http.ListenAndServe(*b.addr, nil))
}

func (b *Blaster) Write(m []byte) {
	b.h.broadcast <- m
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, r.Host)
}

func serveWs(h *Hub, w http.ResponseWriter, r *http.Request) {
	log.Println("handling WS request")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: h, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	log.Println("new client registered")
	go client.writePump()
	go client.readPump()
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Mars Challenge :: Challenge Controller</title>
    <script src="//ajax.googleapis.com/ajax/libs/jquery/2.0.3/jquery.min.js"></script>
    <script type="text/javascript">
        $(function() {
            var conn;
            if (window["WebSocket"]) {
                conn = new WebSocket("ws://{{$}}/ws");
                conn.onclose = function(evt) {
                    $(document.body).text("Connection closed.")
                }
                conn.onmessage = function(evt) {
                    $(document.body).text(evt.data)
                }
            } else {
                $(document.body).text("Your browser does not support WebSockets.")
            }
        });
    </script>
  </head>
  <body>
  </body>
</html>
`))

var homeTemplate2 = `
<!DOCTYPE html>
<html lang="en">
<head>
<title>Mars Challenge Sensor Data</title>
<meta charset="utf-8">
<script type="text/javascript">
window.onload = function () {
    var conn;

    if (window["WebSocket"]) {
      conn = new WebSocket("ws://{{$}}/ws");
      conn.onclose = function(evt) {
        $(document.body).text("Connection closed.");
      };
      conn.onmessage = function(evt) {
        $(document.body).text(evt.data);
      };
    } else {
      $(document.body).text("Your browser does not support WebSockets.");
    }
});
</script>
</head>
<body>
</body>
</html>
`
