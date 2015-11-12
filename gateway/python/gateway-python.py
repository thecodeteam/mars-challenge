#!/bin/python

import json
import logging
import logging.config
import os
import time
import requests
import settings
import websocket


# Logging Initialization
logging.config.dictConfig(settings.GATEWAY_LOGGING)
logger = logging.getLogger("root")

# Global Variables
sensor_endpoints = os.environ['SENSOR_ENDPOINTS'].split(',')
logger.info("Sensor EndPoints: {0}.".format(sensor_endpoints))

controller_endpoint = os.environ['GC_ENDPOINT']
controller_url = 'http://' + controller_endpoint + '/api/readings'  #'http://localhost:8080/api/readings'

admin_password = '98765'
auth_header = {'X-Auth-Token': admin_password}

sensor_data = ''
logger.info("Controller URL: {0}.".format(controller_url))

# {"solarFlare":true,"temperature":-40.18456978907354,"radiation":564,"stamp":"2015-11-11T08:25:49Z"}
while True:
    try:
        temperature = 0
        radiation = 0
        sensor_data = {}
        for sensor in sensor_endpoints:
            sensor_socket = 'ws://' + sensor +'/ws'
            logger.info("Sensor EndPoint: {0}.".format(sensor_socket))
            logger.info("Trying to connect to Sensor Socket: {0:s}".format(sensor_socket))
            ws = websocket.create_connection(sensor_socket)
            result = ws.recv()
            sensor_dict = json.loads(result)
            temperature += sensor_dict['temperature']
            radiation += sensor_dict['radiation']
            logger.info("Sensor data for sensor: {0:s} is - temp: {1:s} radiation: {2:s}".format(sensor, str(sensor_dict['temperature']), str(sensor_dict['radiation'])))
        logger.info("Number of sensor endpoints: {}".format(len(sensor_endpoints)))
        logger.info("Temp: {}".format(temperature))
        logger.info("Radiation: {}".format(radiation))
        temperature = temperature / len(sensor_endpoints)
        radiation = radiation / len(sensor_endpoints)
        logger.info("Post math Temp: {}".format(temperature))
        logger.info("Post math Radiation: {}".format(radiation))
        sensor_data['temperature'] = temperature
        sensor_data['radiation'] = radiation
        sensor_data = json.dumps(sensor_data)

        # Performs the POST with form-encoded data
        response = requests.post(url=controller_url, headers=auth_header, data=sensor_data)

        if response.status_code == 200:
            logger.info("Data has been accepted!")
            logger.info("Sending: {0:s}".format(sensor_data))
            logger.info("HTTP Code: {0}  | Response: {1}".format(str(response.status_code), response.text))

        else:
            logger.error("We got an Error")
            break
    except websocket.WebSocketException as serror:
        logger.info("Error (Sensor Socket): {}".format(str(serror.message)))
        time.sleep(5)
        continue

ws.close()
