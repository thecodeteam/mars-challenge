# Base image
FROM golang:1.7-alpine

# Maintainer info
MAINTAINER Travis Rhoden <travis.rhoden@dell.com>

RUN apk add --no-cache git

# Get the project
RUN go get github.com/codedellemc/mars-challenge/game-controller

# Set working directory
WORKDIR /go/src/github.com/codedellemc/mars-challenge/game-controller

# Expose port
EXPOSE 8080

# Run Game Controller
CMD go run *.go
