#!/	/bin/python
# Library to install run: sudo pip install websocket-client
# if you encounter errors with a Six import
# pip remove six; pip install six 

from websocket import create_connection
ws = create_connection("ws://localhost:8080/ws")


var = 1
while var == 1 :  # This constructs an infinite loop
  result =  ws.recv()
  print "Received '%s'" % result

ws.close()
print "Good bye!"



