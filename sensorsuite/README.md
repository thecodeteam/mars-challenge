# Mars Challenge

## SensorSuite

SensorSuite implements a set of atmospheric sensors as individual applications.
The *Temperature* and *Radiation* sensors expose WebSocket endpoints that the
*Publisher* uses to gather readings and send those readings to the Sensor
Aggregator.

Both the *Temperature* and *Radiation* sensors must be able to connect to the
*Flare* websocket, as the presence of a solar flare effects the algorithm for
calculating a new temperature and radiation level.

## WebSockets

Each sensor uses websockets to send and receive data. Each sensor will bind to
a local address and port, and will expose two endpoints. The home endpoint ("/")
can be visited by a web browser and you will see the data being generated and
refreshed once per second. The second endpoint is for the websocket itself, and
is hosted at "/ws".

The default bindings for each sensor are defined in `sensorsuite/sensorsuite.go`,
and are referenced here for convenience:

Sensor | Port Number
--- | ---
Flare | `0.0.0.0:9000`
Temperature | `0.0.0.0:9001`
Radiation | `0.0.0.0:9002`
Publisher | `0.0.0.0:9003`
Aggregator | `0.0.0.0:9004`

## Running the sensors

Each sensor can be run directly by visiting its main `.go` file and using
`go run`. For example:

```
cd sensorsuite/flare
go run flare.go
```

Each sensor uses environment variables to control what network addresses they
bind to, and where the other sensors are located. The variables supported by
each sensor are:

**Flare**
- SENSOR_LISTEN_ADDRESS

**Temperature**
- SENSOR_LISTEN_ADDRESS
- SENSOR_FLARE_ADDRESS

**Radiation**
- SENSOR_LISTEN_ADDRESS
- SENSOR_FLARE_ADDRESS

**Publisher**
- SENSOR_LISTEN_ADDRESS
- SENSOR_TEMPERATURE_ADDRESS
- SENSOR_RADIATION_ADDRESS
- SENSOR_AGGREGATOR_ADDRESS

## Building the sensorsuite

Should you wish to compile the sensors into an executable binary, `go build`
will produce the output. Example:

```
cd sensorsuite/flare
go build flare.go
```

That will produce a binary by the same name, e.g. `flare`. This binary can be
executed directly, or placed in a Docker image.

## Dependencies

Building and running the sensors requires the following:

- a working `go` installation
- a copy of this code: `go get github.com/codedellemc/mars-challenge/sensorsuite`
- websockets: `go get github.com/gorilla/websocket`
- viper: `go get github.com/spf13/viper`

## Sensor Data

The websocket endpoint located in `/ws` emits the following data formatted in
JSON once per second:

**Flare**

| Field | Description |
|---|---|
| `solarFlare` | Whether there is a solar flare or not. Values are either `true` or `false` |
| `stamp` | A string timestamp |

**Temperature**

| Field | Description |
|---|---|
| `temperature` | Float value ranging from -142.0 to 35.0 representing the current temperature in Celsius.
| `stamp` | A string timestamp |

**Radiation**
| Field | Description |
|---|---|
| `radiation` | Integer value ranging from 0 to 1000. |
| `stamp` | A string timestamp |

**Publisher**

| Field | Description |
|---|---|
| `temperature` | Copied data from *Temperature* sensor |
| `radiation` | Copied data from *Radiation* sensor |
| `stamp` | A string timestamp |
