# Mars Challenge 

## Game Controller

Run the controller:

    go run *.go

Run the controller to accept external sensor readings and with a custom administrator token:

    ADMIN_TOKEN=1234 AUTO_READINGS=false go run *.go

### Contents

- [Dependencies](#dependencies)
- [Docker container](#docker-container)
- [Game logic](#game-logic)
- [Websocket](#websocket)
- [API](#api)

## Docker container

The Docker image for the Game Controller is located in [Docker hub](https://registry.hub.docker.com/u/emccode/mars-challenge-controller/). To get the image just run `docker pull emccode/mars-challenge-controller`.

You can provide an admin token of your choice to perform some privileged requests. To do so, just set the `ADMIN_TOKEN` environment variable to whatever you want. If no admin token is provided, a random token will be generated and displayed in the logs.

If you want to provide external readings to the game controller you need to set the `AUTO_READINGS` environment variable to `false`. If it is not set it will default to `true`. Read the API specification to learn more about how to provide external readings.

##### Example

In the following case we are going to use "1234" as admin token and listen on the port 80 on our host machine that will bind to the container's exposed port 8080.

    docker run -d --name=controller -p 80:8080 -e ADMIN_TOKEN=1234 emccode/mars-challenge-controller

An example modifying the **AUTO_READINGS**: 

    docker run -d --name=controller -p 80:8080 -e ADMIN_TOKEN=1234 -e AUTO_READINGS=false  emccode/mars-challenge-controller


## Dependencies

The game controller has the following dependencies:

  - mux: `go get github.com/gorilla/mux`
  - websockets: `go get github.com/gorilla/websockets`

## Game logic

Temperature and radiation levels affect team's energy and life respectively. Every second the game engine checks the levels and applies the following algorithm to determine how energy and life are affected.

    radiationRatio = (currentRadiation - minRadiation) / (maxRadiation - minRadiation)
    temperatureRatio = (currentTemperature - minTemperature) / (maxTemperature - minTemperature)

    if shield is ON:
      energyLoss = radiationRatio * 5
      team.energy = team.energy - ceil(energyLoss)
    else:
      lifeLoss = radiationRatio * 5
      team.life = team.life - ceil(lifeLoss)

      energyGain = temperatureRatio * 5
      team.energy = team.energy + ceil(energyGain)

`temperatureRatio` and `radiationRatio` will range from 0 to 1 depending on the relative value of the current level, e.g., if we want to calculate the energy gain for a temperature of -53.5, `temperatureRatio` will be 0.5 because it is exactly between the minimum (-142) and the maximum (35). Therefore, the `energyGain` will be 0.5 * 5 = 2.5, rounded up to 3.


## Websocket

The websocket endpoint is located in `/ws`. It sends the following JSON structure every second.

| Field | Description |
|---|---|
| `running` | Whether the game is running or not. Values are either `true` or `false` |
| `startedAt` | Date and time when the game was started. |
| `timestamp` | Current server date and time. |
| `readings` | `Sensor` data structure |
| `teams` | List of `Team` data structure |

`Sensor` data structure:

| Field | Description |
|---|---|
| `solarFlare` | Whether there is a solar flare or not. Values are either `true` or `false` |
| `temperature` | Float value ranging from -142 to 35 representing the current temperature in Celsius. |
| `radiation` | Integer value ranging from 0 to 1000. |

`Team` data structure:

| Field | Description |
|---|---|
| `name` | Name of the team. Unique. |
| `energy` | Integer representing the energy level ranging from 0 to 100. |
| `life` | Integer representing the life level ranging from 0 to 100. 0 means dead. |
| `shield` | Whether the shield is ON or OFF for this team. Values are either `true` or `false` |

##### Examples

Stopped game with no teams:

    {"running":false,"startedAt":"0001-01-01T00:00:00Z", "timestamp": "2015-08-04T15:09:20.923151468+02:00","readings":{"solarFlare":false,"temperature":-53.5,"radiation":500},"teams":[]}

Stopped game with two teams ready to play:

    {"running":false,"startedAt":"0001-01-01T00:00:00Z","timestamp": "2015-08-04T15:09:20.923151468+02:00","readings":{"solarFlare":false,"temperature":-53.5,"radiation":500},"teams":[{"name":"bar","energy":100,"life":100,"shield":false},{"name":"foo","energy":100,"life":100,"shield":true}]}

Game in progress with two teams playing:

    {"running":true,"startedAt":"2015-07-31T12:21:49.511228099+02:00","timestamp": "2015-07-31T12:21:56.215288591+02:00","readings":{"solarFlare":false,"temperature":-3.43,"radiation":785},"teams":[{"name":"bar","energy":45,"life":11,"shield":false},{"name":"foo","energy":7,"life":57,"shield":true}]}


Game ended with a winner team:

    {"running":false,"startedAt":"2015-07-31T12:21:49.511228099+02:00","timestamp": "2015-07-31T12:21:56.215288591+02:00","readings":{"solarFlare":false,"temperature":-53.5,"radiation":500},"teams":[{"name":"bar","energy":0,"life":8,"shield":false},{"name":"foo","energy":4,"life":0,"shield":false}]}

## API

#### POST /api/readings

Endpoint to provide readings from an external source. **Requires administrator rights**. Only enabled if the `AUTO_READINGS` environment variable is `false` when launching the program.

##### Request Body

The `Sensor` data structure is provided as a JSON-formatted text.

##### Example

    $ curl -i -H 'X-Auth-Token: 1234' -X POST -d '{"solarFlare":true,"temperature":-100,"radiation":384}' http://localhost:8080/api/readings
    HTTP/1.1 200 OK
    Date: Tue, 04 Aug 2015 10:34:07 GMT
    Content-Length: 16
    Content-Type: text/plain; charset=utf-8

    Readings updated

If the game does not accept external readings (i.e. `AUTO_READINGS` is `false`):

    $ curl -i -H 'X-Auth-Token: 1234' -X POST -d '{"solarFlare":false,"temperature":12,"radiation":93}' http://localhost:8080/api/readings
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Tue, 04 Aug 2015 10:33:40 GMT
    Content-Length: 75

    Game running with auto generated readings, not accepting external readings

If one of the provided readings is out of bounds:

    $ curl -i -H 'X-Auth-Token: 1234' -X POST -d '{"solarFlare":true,"temperature":-150,"radiation":54}' http://localhost:8080/api/readings
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Tue, 04 Aug 2015 10:23:08 GMT
    Content-Length: 53

    Temperature not within valid range [-142.00, 35.00]


#### GET /api/config

Endpoint to obtain the game configuration.

##### Response Body

JSON-formatted text containing the following keys:

- `minTemperature`: Float value containing the minimum temperature.
- `maxTemperature`: Float value containing the maximum temperature.
- `minRadiation`: Integer value containing the minimum radiation.
- `maxRadiation`: Integer value containing the minimum radiation.
- `autoReadings`: Boolean value indicating whether the game auto-generates sensor readings or not.

##### Example

    $ curl -i -X GET http://localhost:8080/api/config
    HTTP/1.1 200 OK
    Date: Tue, 04 Aug 2015 11:35:55 GMT
    Content-Length: 100
    Content-Type: text/plain; charset=utf-8

    {"maxTemperature":35,"minTemperature":-142,"maxRadiation":1000,"minRadiation":0,"autoReadings":true}



#### POST /api/start

Starts the game. **Requires administrator rights.**

##### Example

    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
    HTTP/1.1 200 OK
    Date: Fri, 31 Jul 2015 09:30:48 GMT
    Content-Length: 12
    Content-Type: text/plain; charset=utf-8

    Game started

It will return a 400 if the game is already running:

    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 31 Jul 2015 10:21:51 GMT
    Content-Length: 44

    Game is already started, not doing anything

#### POST /api/stop

Stops the game if it is running. **Requires administrator rights.**

##### Example

    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/stop
    HTTP/1.1 200 OK
    Date: Fri, 31 Jul 2015 09:49:11 GMT
    Content-Length: 12
    Content-Type: text/plain; charset=utf-8

    Game stopped

#### POST /api/reset

Resets the game if it is not running. Cleans all teams and sets the values to their initial state. **Requires administrator rights.**

##### Example

    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/reset
    HTTP/1.1 200 OK
    Date: Fri, 31 Jul 2015 09:51:37 GMT
    Content-Length: 23
    Content-Type: text/plain; charset=utf-8

    Game reset successfully


If the game is still running:

    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/reset
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 31 Jul 2015 10:19:46 GMT
    Content-Length: 38

    Cannot reset game while it is running


#### POST /api/join/{team_name}

Join a team to the game. The given team name must be not in use. An auth token will be returned in the body if the request is processed successfully.

##### Example

    $ curl -i -X POST http://localhost:8080/api/join/foobar
    HTTP/1.1 200 OK
    Date: Fri, 31 Jul 2015 09:54:47 GMT
    Content-Length: 16
    Content-Type: text/plain; charset=utf-8

    3276918835a4a462

If the team name already exists:

    $ curl -i -X POST http://localhost:8080/api/join/foobar
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 31 Jul 2015 09:55:39 GMT
    Content-Length: 30

    Team 'foobar' already exists.



#### POST /api/kick/{team_name}

Kicks a team from the game. **Requires administrator rights.**

##### Example

    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/kick/foobar
    HTTP/1.1 200 OK
    Date: Tue, 04 Aug 2015 08:55:26 GMT
    Content-Length: 27
    Content-Type: text/plain; charset=utf-8

    Team 'foobar' left the game

If the team name does not exist:

    $ curl -i -X POST http://localhost:8080/api/kick/foobar2
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 31 Jul 2015 09:55:39 GMT
    Content-Length: 30

    Team 'foobar2' does not exist

#### POST /api/shield/{enable|disable}

Enables or disables the shield. **Requires a team token**.

##### Example Shield Enabled

    $ curl -i -H 'X-Auth-Token: 1335aa6af5d0289f' -X POST http://localhost:8080/api/shield/enable
    HTTP/1.1 200 OK
    Date: Fri, 31 Jul 2015 10:38:20 GMT
    Content-Length: 27
    Content-Type: text/plain; charset=utf-8

    Shield successfully enabled

If the token is not valid or not set:

    $ curl -i -H 'X-Auth-Token: 1111' -X POST http://localhost:8080/api/shield/enable
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 31 Jul 2015 10:44:38 GMT
    Content-Length: 24

    Could not enable shield

##### Example Shield Disabled

    $ curl -i -H 'X-Auth-Token: 1335aa6af5d0289f' -X POST http://localhost:8080/api/shield/disable

    HTTP/1.1 200 OK
    Date: Fri, 31 Jul 2015 10:38:20 GMT
    Content-Length: 27
    Content-Type: text/plain; charset=utf-8

    Shield successfully disabled

If the token is not valid or not set:

    $ curl -i -H 'X-Auth-Token: 1111' -X POST http://localhost:8080/api/shield/disable
    HTTP/1.1 400 Bad Request
    Content-Type: text/plain; charset=utf-8
    Date: Fri, 31 Jul 2015 10:44:38 GMT
    Content-Length: 24

    Could not disable shield
