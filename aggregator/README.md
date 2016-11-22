# Mars Challenge

## Aggregator

The *Aggregator* application receives temperature and radiation data from one or
more SensorSuite *Publishers*, and sends an averaged value to the Game
Controller.

The *Aggregator* must be able to connect to the same *Flare* websocket as the
sensors in the SensorSuite. It then listens for any incoming connections on its
own websocket, where it can receive sensor updates from any number of
*Publishers*.

Once per second, it averages all the values it has received, and does an HTTP
POST to the Game Controller. In order for this to succeed, the *Aggregator* must
be able to reach the Game Controller, and it also must be given the correct
Admin Token.

## WebSocket

The *Aggregator* uses websockets to receive data. The application will bind to
a local address and port, and will expose two endpoints. The home endpoint ("/")
can be visited by a web browser and you will see the averaged data and
refreshed once per second. The second endpoint is for the websocket itself, and
is hosted at "/ws".

The default bindings for are defined in `sensorsuite/sensorsuite.go`,
and are referenced here for convenience:

Sensor | Port Number
--- | ---
Flare | `0.0.0.0:9000`
Temperature | `0.0.0.0:9001`
Radiation | `0.0.0.0:9002`
Publisher | `0.0.0.0:9003`
Aggregator | `0.0.0.0:9004`

## Running the Aggregator

The Aggregator can be run directly by visiting its main `.go` file and using
`go run`. For example:

```
cd aggregator
go run aggregator.go
```

Environment variables are used to control what network addresses the app
binds to, and where the other endpoints are found. The variables supported by
Aggregator are:

- LISTEN_ADDRESS
- SENSOR_FLARE_ADDRESS
- GC_ADDRESS
- POST_GC
- ADMIN_TOKEN

## Building the Aggregator

Should you wish to compile the aggregator into an executable binary, `go build`
will produce the output. Example:

```
cd aggregator
go build aggregator.go
```

That will produce a binary by the same name, e.g. `aggregator`. This binary can
be executed directly on your host.

**Hint** Different flags to `go build` are needed if you want to produce a
binary that can be executed in a Docker image.

## Dependencies

Building and running the aggregator requires the following:

- a working `go` installation
- a copy of this code: `go get github.com/codedellemc/mars-challenge/aggregator`
- websockets: `go get github.com/gorilla/websocket`
- viper: `go get github.com/spf13/viper`
