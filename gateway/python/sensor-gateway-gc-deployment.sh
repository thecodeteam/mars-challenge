#!/usr/bin/env bash

# Start the Sensor-Client container

docker run -d --name=sensor -p 80:8080  emccode/mars-challenge-client

# Start the Game Server Container

docker run -d --name=controller -p 81:8080 -e ADMIN_TOKEN=1234 -e AUTO_READINGS=false  emccode/mars-challenge-controller

# Start the Game Server UI Container

docker run -d --name dashboard11 -e WS_ENDPOINT=10.0.2.15:8080/ws -p 86:80 emccode/mars-challenge-dashboard

# Start the Gateway

docker run -d --name=gateway  -e SENSOR_ENDPOINT=localhost:80 -e GC_ENDPOINT=localhost:81  pbutlerm/mars-challenge-gateway

#start the Team Container