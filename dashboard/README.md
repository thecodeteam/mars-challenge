# Mars Challenge

## Dashboard

The dashboard connects to the controller's websocket and displays the information in a more visual manner.

### Contents

- [Build and development](#build-and-development)
- [Docker container](#docker-container)

## Build and development

Run `grunt` for building and `grunt serve` for preview.


## Docker container

A Docker image for the dashboard is provided and is located in [Docker hub](https://registry.hub.docker.com/u/emccode/mars-challenge-dashboard/). To get the image just run `docker pull emccode/mars-challenge-dashboard`.

You must provide the controller's websocket endpoint for the dashboard to connect and receive game information. To do so, just set the `WS_ENDPOINT` environment variable when running the Docker container.

##### Example

If our websocket endpoint is located in `mars.cloudapp.net/ws`, the Docker run command would look like this:

    docker run -d --name dashboard -e WS_ENDPOINT=mars.cloudapp.net/ws -p 80:80 emccode/mars-challenge-dashboard


Note that the dashboard container exposes the port 80.
