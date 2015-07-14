#!/bin/python
# Requires the following library to install: sudo pip install websocket-client
# if you encounter errors with a Six import:
# you can try: pip remove six; pip install six 

from websocket import create_connection
ws = create_connection("ws://localhost:8080/ws")

while True :  # This constructs an infinite loop
  result =  ws.recv()
  print "Received '%s'" % result
  
ws.close()
print "Good bye!"



