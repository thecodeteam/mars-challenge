#!/bin/python
# Gateway Tier: It takes sensor information, logs to the console and passes that information to the Game Contoller
# Requires the following library to install: sudo pip install websocket-client
# if you encounter errors with a Six import:
# you can try: pip remove six; pip install six
import time
import logging
import logging.config
import os

import websocket
import requests

import settings

# Logging Initialization
logging.config.dictConfig(settings.GATEWAY_LOGGING)
logger = logging.getLogger("root")



# Global Variables
sensor_endpoint = os.environ['SENSOR_ENDPOINT']
controller_endpoint = os.environ['GC_ENDPOINT']

controller_url = 'http://' + controller_endpoint + '/api/readings'  #'http://localhost:8080/api/readings'
admin_password = '1234'
auth_header = {'X-Auth-Token': admin_password}
sensor_socket = 'ws://' + sensor_endpoint +'/ws'  # "ws://localhost:8085/ws"
sensor_data = ''

logger.info("Sensor EndPoint: {0}.".format(controller_url))
logger.info("Game Controller EndPoint: {0}.".format(sensor_socket))

#Check that the Sensor is ready to serve weather data
while True:
    try:
        logger.info("Trying to connect to Sensor Socket: {0:s}".format(sensor_socket))
        ws = websocket.create_connection(sensor_socket)
        break
    except websocket.WebSocketException as serror:
        logger.info("Error (Sensor Socket): {}".format(str(serror.message)))
        time.sleep(5)
        continue


# curl -i -H 'X-Auth-Token: 1234' -X POST -d '{"solarFlare":true,"temperature":-100,"radiation":384}' http://localhost:8080/api/readings

# Start collecting the Sensor Data and sending the data to the Game Server
while True:

    # Get Data from the Sensor Websocket
    result = ws.recv()
    sensor_data = result

    # Performs the POST with form-encoded data
    response = requests.post(url=controller_url, headers=auth_header, data=sensor_data)

    if response.status_code == 200:
        logger.info("Data has been accepted!")
        logger.info("Sending: {0:s}".format(result))
        logger.info("HTTP Code: {0}  | Response: {1}".format(str(response.status_code), response.text))

    else:
        logger.error("We got an Error")
        logger.error("HTTP Code: {0}  | Response: {1}".format(str(response.status_code), response.text))


ws.close()
