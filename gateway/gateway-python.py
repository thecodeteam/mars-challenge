#!/bin/python
# Gateway Tier: It takes sensor information, logs to the console and passes that information to the Game Contoller
# Requires the following library to install: sudo pip install websocket-client
# if you encounter errors with a Six import:
# you can try: pip remove six; pip install six

from websocket import create_connection
import requests
import json

# Global Variables
controller_url = 'http://localhost:8080/api/readings'
admin_password = '1234'
auth_header = {'X-Auth-Token': admin_password}
sensor_data = ''

ws = create_connection("ws://localhost:8085/ws")

# curl -i -H 'X-Auth-Token: 1234' -X POST -d '{"solarFlare":true,"temperature":-100,"radiation":384}' http://localhost:8080/api/readings

while True:  # This constructs an infinite loop

    # Get Data from the Sensor Websocket
    result = ws.recv()
    sensor_data = result

    # Performs the POST with form-encoded data
    response = requests.post(url=controller_url, headers=auth_header, data=sensor_data)

    if response.status_code == 200:
        print("Data has been accepted!")
        print ("Sending: {0:s}".format(result))
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)

    else:
        print("We got an Error")
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)


ws.close()
