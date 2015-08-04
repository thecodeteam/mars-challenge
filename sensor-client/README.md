# Mars Challenge

## Sensor client

The Sensor Client will run on the Raspberry PI and will expose a Websocket endpoint for the aggregator service to obtain the sensor readings and subsequently send them to the Game Controller.

Run the sensor client:

    go run *.go

### Contents

- [Docker container](#docker-container)
- [Dependencies](#dependencies)
- [Websocket](#websocket)

## Docker container

The Docker image for the Sensor Client is located in [Docker hub](https://registry.hub.docker.com/u/emccode/mars-challenge-client/). To get the image just run `docker pull emccode/mars-challenge-client`.

##### Example

In the following case we are going to listen on the port 80 on our host machine that will bind to the container's exposed port 8080.

    docker run -d --name=sensor -p 80:8080  emccode/mars-challenge-client


## Dependencies

The game controller has the following dependencies:

  - mux: `go get github.com/gorilla/mux`
  - websockets: `go get github.com/gorilla/websockets`



## Websocket

The websocket endpoint is located in `/ws`. It sends the following JSON structure every second corresponding to the Sensor data.

| Field | Description |
|---|---|
| `solarFlare` | Whether there is a solar flare or not. Values are either `true` or `false` |
| `temperature` | Float value ranging from -142 to 35 representing the current temperature in Celsius. |
| `radiation` | Integer value ranging from 0 to 1000. |

The endpoint in `/` (root) contains a basic Javascript client that connects to the Websocket endpoint.
