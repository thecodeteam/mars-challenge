## Game Controller Engine

Execute the the engine:

`go run *.go`

To get the Web-sockets output with the sensor readings go to `http://your_host:8080`.

## Missing Packages

You may get an error regarding missing packages. Please use the following commands to add them:

- Mux: `go get github.com/gorilla/mux`
- Web-sockets: `go get github.com/gorilla/websockets`
