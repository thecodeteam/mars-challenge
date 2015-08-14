# Base image
FROM golang:1.4.2

# Maintainer info
MAINTAINER Adrian Moreno <adrian.moreno@emc.com>

# Get the project
RUN go get github.com/emccode/mars-challenge/game-controller
RUN go get -u -v golang.org/x/net/websocket

# Set working directory
WORKDIR /go/src/github.com/emccode/mars-challenge/gateway/go

# Expose port
EXPOSE 8080

# Run Game Controller
CMD go run *.go
