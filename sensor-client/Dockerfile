# Base image
FROM golang:1.4.2

# Maintainer info
MAINTAINER Adrian Moreno <adrian.moreno@emc.com>

# Get the project
RUN go get github.com/emccode/mars-challenge/sensor-client

# Set working directory
WORKDIR /go/src/github.com/emccode/mars-challenge/sensor-client

# Expose port
EXPOSE 8080

# Run Game Controller
CMD go run *.go
