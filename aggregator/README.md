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

## Docker container

The Docker image for the Aggregator is located in
[Docker hub](https://registry.hub.docker.com/u/emccode/mars-challenge-aggregator/).
To get the image just run `docker pull emccode/mars-challenge-aggregator`.

You will need to provide the address of the Solar Flare sensor, and for the
Game Controller. In order to successfully POST to the game controller, you must
also provide the same `ADMIN_TOKEN` that was used to start the Game Controller.

The Aggregator can be started *before* the Game Controller, and it will connect
when it is available. However, the Aggregator will not run without being in
contact with the Solar Flare sensor.

### Example

In the following case we are going to use "1234" as admin token and connect to
the Flare sensor running in a container named `flare` at port `9000`. We also
map port `9004` to the same port on our local host so we can point our
webrowser at `http://localhost:9004` and see what data the Aggregator has.

    docker run -d --name=aggregator --link flare -p 9004:9004 -e SENSOR_FLARE_ADDRESS=flare:9000 -e ADMIN_TOKEN=1234 -e GC_ADDRESS=<some ip>:8080 emccode/mars-challenge-aggregator

The IP address or hostname that you use for the Flare Sensor and the Game
Controller will depend on how you are running the container (e.g. `docker run`
vs `docker-compose`) and whether you are connecting to your own sensor and
controller or one run by the Hackathon judges.

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
