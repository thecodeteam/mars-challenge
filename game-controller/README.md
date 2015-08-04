# Mars Challenge

## Game Controller

Execute the controller:

    go run *.go

### Contents

- [Dependencies](#dependencies)
- [Docker container](#docker-container)
- [Websocket](#websocket)
- [API](#api)

## Docker container

The Docker image for the Game Controller is located in [Docker hub](https://registry.hub.docker.com/u/emccode/mars-challenge-controller/). To get the image just run `docker pull emccode/mars-challenge-controller`.

You can provide an admin token of your choice to perform some privileged requests. To do so, just set the `ADMIN_TOKEN` environment variable to whatever you want. If no admin token is provided, a random token will be generated and displayed in the logs.

##### Example

In the following case we are going to use "1234" as admin token and listen on the port 80 on our host machine that will bind to the container's exposed port 8080.

    docker run -d --name=controller -p 80:8080 -e ADMIN_TOKEN=1234 emccode/mars-challenge-controller


## Dependencies

The game controller has the following dependencies:

  - mux: `go get github.com/gorilla/mux`
  - websockets: `go get github.com/gorilla/websockets`

## Websocket

The websocket endpoint is located in `/ws`. It sends the following JSON structure every second.

| Field | Description |
|---|---|
| `running` | Whether the game is running or not. Values are either `true` or `false` |
| `startedAt` | Date and time when the game was started. |
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

    {"running":false,"startedAt":"0001-01-01T00:00:00Z","readings":{"solarFlare":false,"temperature":-53.5,"radiation":500},"teams":[]}

Stopped game with two teams ready to play:

    {"running":false,"startedAt":"0001-01-01T00:00:00Z","readings":{"solarFlare":false,"temperature":-53.5,"radiation":500},"teams":[{"name":"bar","energy":100,"life":100,"shield":false},{"name":"foo","energy":100,"life":100,"shield":true}]}

Game in progress with two teams playing:

    {"running":true,"startedAt":"2015-07-31T12:21:49.511228099+02:00","readings":{"solarFlare":false,"temperature":-3.43,"radiation":785},"teams":[{"name":"bar","energy":45,"life":11,"shield":false},{"name":"foo","energy":7,"life":57,"shield":true}]}


Game ended with a winner team:

    {"running":false,"startedAt":"2015-07-31T12:21:49.511228099+02:00","readings":{"solarFlare":false,"temperature":-53.5,"radiation":500},"teams":[{"name":"bar","energy":0,"life":8,"shield":false},{"name":"foo","energy":4,"life":0,"shield":false}]}

## API

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
