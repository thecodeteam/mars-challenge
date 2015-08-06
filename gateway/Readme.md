 
 
 Start the Sensor-Client container

 	docker run -d --name=sensor -p 80:8080  emccode/mars-challenge-client

 Start the Game Server Container 

 	docker run -d --name=controller -p 81:8080 -e ADMIN_TOKEN=1234 -e AUTO_READINGS=false  emccode/mars-challenge-controller

 Start the Game Server UI Container 

 	docker run -d --name dashboard -e WS_ENDPOINT=localhost:81/ws -p 82:80 emccode/mars-challenge-dashboard

 Start the Gateway

	docker run -d --name=gateway  -e SENSOR_ENDPOINT=localhost:80 -e GC_ENDPOINT=localhost:81  pbutlerm/mars-challenge-gateway
Set SENSOR_ENDPOINT=localhost:80
	set GC_ENDPOINT=localhost:81
	
	
	
	Use as base docker pull python:2.7
	
	create a requirements.txt
	https://registry.hub.docker.com/u/library/python/ 