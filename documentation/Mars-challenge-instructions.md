# Mars Challenge Instructions 

### Contents

- [Introduction](#Introduction)
- [Goals](#goals)
- [Requirements](#requirements)
- [Winning](#winning)
- [Challenge Assets](#challenge-assets)
- [Participant's Implementation](#participants-implementation)
- [Testing the Command and Control Center](#testing-the-command-and-control-center)
- [System Tiers](#system-tiers)
- [System Deployment](#system-deployment)
- [Points Table](#points-table)


## Introduction 

Welcome to the Mars Challenge, you and your team are about to embark into a 7 hour space adventure that will challenge your **developer, operational, data analysis and infrastructure** skills. The Challenge puts you in the arid lands of Mars: 


*You and your team just landed in Mars. As you prepare your base of operations, massive Sun storms are coming your way. That was the last communication you had from Earth. Now connectivity to earth has been cut, and although the base has enough energy, your protective shields can only be running for a few minutes at the time, without recharging.* 

*Your only chance of survival is to monitor the current temperature and radiation levels in the planet atmosphere to detect solar flares and activate your base shields for protection.*

*You only have a few hours, 7 of them, to implement a sensor array, build and deploy the monitoring application to engage/disengage your shields, then fine tune an algorithm based on your data analysis that decides when to charge your shields and when to engage them for protection. Will you and your team survive?*

You and your team will have at your disposal the necessary tools to survive and win the challenge, however you will need all wits and skills to work together and implement a solution that allows you to survive and compete against other teams to find out who would be the last survivor.


## Goals 

The Mars challenge was build by people passionate about technology for people as passionate or more about technology. The goals for the challenge are the following: 

- Network with people that share the same interests as you do
- Gather on teams to accomplish something fun and learn something while doing it
- Build and deploy a modern distributed application using best practices.
- Deploy a distributed application using Containers. 
- Practice container configuration, execution and debugging. 
- Deploy a distributed application using modern deployment technologies like  Docker Swarm, Compose, Mesos, and Kubernetes).
- Develop an application on the Internet of Things (IoT) space.

## Requirements

This Challenge requires the following for their participants:

- Each Participant needs to bring their own computer/laptop
- Have a good attitude and plan to have serious technical fun
- Be a Team player 
- Adhere to the LinuxCon/ContainterCon/MesosCon code of Conduct

## Winning

The Challenge is designed to be completed in 7 hours. Each team will have to choose which features [(tasks)](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "List of Tasks") wants to implement and deploy as part of their solution. **The only condition** is to have a working solution that satisfies task **[CC-1](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md)** (demonstrate participation in an official Game). Refer to the [Testing the Command and Control Center](#testing-the-command-and-control-center) section to get this part running.

**[Each task your team accomplishes has a point value](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")**

For each task, functionality needs to be successfully demoed to a Judge, then points will be assigned. Point count for each team will be visible for the entire length of the contest.

**The team that has the most amount of points by the end of the seven hours wins the challenge.** 


## Challenge Assets

For this challenge we are providing most of the services, the codebases, and associated containers for the tiers that make the solution. 


### Services and Codebases locations

These are the locations for code samples and information for each one of the services:


|Service Name|Folder Location|
|----|----|
|Sensor Client Service|[https://github.com/emccode/mars-challenge/tree/master/sensor-client](https://github.com/emccode/mars-challenge/tree/master/sensor-client)|
|Gateway and Aggregator Service|[https://github.com/emccode/mars-challenge/tree/master/gateway](https://github.com/emccode/mars-challenge/tree/master/gateway)|
|Data Repository| Refer to the [System Tiers](#system-tiers) section for more information . |
|Data Analysis | Refer to the [System Tiers](#system-tiers) for more information.|
|Team Command & Control|[https://github.com/emccode/mars-challenge/tree/master/clients/python](https://github.com/emccode/mars-challenge/tree/master/clients/python)|
|Game Controller|[https://github.com/emccode/mars-challenge/tree/master/game-controller](https://github.com/emccode/mars-challenge/tree/master/game-controller)|
|Game Controller Dashboard|[https://github.com/emccode/mars-challenge/tree/master/dashboard](https://github.com/emccode/mars-challenge/tree/master/dashboard)|
|RaspBerry Pi Sensor Setup|[https://github.com/emccode/mars-challenge/blob/master/documentation/Raspberry-Go-Weather-Simulator-Setup.md](https://github.com/emccode/mars-challenge/blob/master/documentation/Raspberry-Go-Weather-Simulator-Setup.md)|

### Containers

These are the containers available for all the teams to use: 

|Service Name|Container Location|
|----|----|
|Sensor Client Service (Tier 1)| [https://hub.docker.com/r/emccode/mars-challenge-client/](https://hub.docker.com/r/emccode/mars-challenge-client/ "Sensor Client Service")|
|Gateway and Aggregator Service (Tier 2)| [https://hub.docker.com/r/emccode/mars-challenge-gateway-py/](https://hub.docker.com/r/emccode/mars-challenge-gateway-py/ "Gateway and Aggregator Service")|
|Game Challenge Controller|[https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard|[https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

**Note:** You can use the code implementations and containers provided or you can choose to build your own. There are [points](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table") associated for each decision that you make. 

## Participant's Implementation

 Participants have an opportunity to build multiple tiers of functionality. These are been organized on the following groups:
 
 - Section 1: Getting Data From the Sensors to the Command and Control Service
 - Section 2: Perform Analytic and Data saving operations 
 - Section 3: DevOps and Application Management 


###  Section 1: Getting Data From the Sensors to the Command and Control Service

This section focuses on getting the sensor data from the Sensor services to the Command and Control Tier.  

 ![Mars Challenge Participants Stage 1](https://github.com/emccode/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage1.jpg).

In order to implement section 1 you will need the following components: 

|Service Name|Container Location|
|----|----|
|Sensor Client Service (Tier 1)| [https://hub.docker.com/r/emccode/mars-challenge-client/](https://hub.docker.com/r/emccode/mars-challenge-client/ "Sensor Client Service")|
|Gateway and Aggregator Service (Tier 2)| [https://hub.docker.com/r/emccode/mars-challenge-gateway-py/](https://hub.docker.com/r/emccode/mars-challenge-gateway-py/ "Gateway and Aggregator Service")|
|Game Challenge Controller|[https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard|[https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

A Demo implementations of a Command and Control (Tier 5) Service, implemented in Python, is located in the following folder: [https://github.com/emccode/mars-challenge/tree/master/clients/python](https://github.com/emccode/mars-challenge/tree/master/clients/python "Command and Control Demo")

Once you have the components in place, you and your team will have to create a deployment and orchestrate how services are deployed in a specific order:

1. Sensor Service
2. Game Controller
3. Game Controller UI
4. Gateway and Aggregation service
5. Team Control and Command Center service

You will get [points](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table") for automating the deployment and orchestration of this section. 


### Section 2: Perform Analytic and Data Saving Operations 

**Section 2:** This section focuses on getting the rest of services that work with the Team Command and Control center:

![Mars Challenge Participants Stage 2](https://github.com/emccode/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage2.jpg)


The first step would be make sure you have a Command and Control service working. You can test your implementation of the Command & Control (C&C) service using the information in [Testing the Command and Control Center](#testing-the-command-and-control-center) section. 


Once you have the C&C service ready and on a container, you can start adding the other pieces that you may need. These are some of the options of the services you can implement: 

- Data Repository Tier
- Data Analytic Tier 
- Backup Service

Please Refer to the **[System Tiers](#system-tiers)** section for more details on these. Remember, each one of these tier will get you more [points](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table") and at the same time it will increase the complexity of the deployments.


### Section 3: DevOps and Application Management 

Even on Mars, you will need to deploy your application, in your infrastructure to get to the functionality you and your team have implemented. The goal is for you and your team to get experience managing and deploying micro services applications. Some important aspects of managing micro services are:

- **Service Discovery:** Implement Service Discovery between all the containers deployed in the solution.
- **Service Monitoring:** Implement Service Monitoring of all containers deployed in the solution.
- **Service Configuration:** Implement Service configuration for all the Tiers/containers.
- **Service Orchestration:** Implement Service Orchestration for all the Tiers/containers.
- **Automated Service Deployment:** Deployment of the implemented system using Docker Tooling, Kubernetes, Me
- sos, Puppet/Ansible/Chef/Saltstack or other Tooling.
- **Logging Router:** Deploy, connect and route the logs of the application using a logging router.

**Note:** We are providing VMs on a Cloud provider for you to do deploy your system. You can use any tooling to automate this process. 


## Testing the Command and Control Center

In order to Test your Team Command and Control center you can run the Mar's Game Controller in testing mode. You will need the to run the Game Controller and Game Controller's Dashboard. 

![Testing the Command and Control Center](https://github.com/emccode/mars-challenge/blob/master/documentation/images/Mars-challenge-testing-control-center.JPG)


The Game containers are located here: 

|Service Name|Container Location|
|----|----|
|Game Challenge Controller|[https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard|[https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

The provided containers are Docker containers. You need to have docker installed on the host(s)  that you plan to run these on. you will be start up the containers on the following order: 

1. Game Controller
2. Game Controller UI
3. Command and Control Service

Please use the following commands to setup the containers: 
 
First, start by executing the Game Controller Container. You can define both the listening port: `-p 80:8080` and the Admin Token: **`-e ADMIN_TOKEN=1234`** to what you may need:

    docker run -d --name=controller -p 80:8080 -e ADMIN_TOKEN=1234 emccode/mars-challenge-controller

you can check the Game Controller service running by running a browser against the host: `http://<host ip address>:80`.

Second, execute the Game Controller Dashboard and map it to the endpoint used by the Game Controller: 

    docker run -d --name dashboard -e WS_ENDPOINT=localhost:80/ws -p 82:80 emccode/mars-challenge-dashboard

Where `-e WS_ENDPOINT=localhost:80/ws` is the location of the Game Controller and  `-p 82:80` is the endpoint of the Game Controller Dashboard. You can check the Game Controller Dashboard service running on the specified address . In our example, the address would be: `http://<host ip address>:82`.

The last step is to execute the Command and Control Center code. You can run the code directly or run it as a container. 

A Demo of a Command and Control (Tier 5) Center, implemented in Python, is provided and  is located in the following folder: [https://github.com/emccode/mars-challenge/tree/master/clients/python](https://github.com/emccode/mars-challenge/tree/master/clients/python "Command and Control Demo")

The Demo implements all the interfaces provides by the Game Controller and implements Shield operations based on sensor data. 

For our example we will run the code directly from command line/terminal: 

1. Get the code from the code repository: `https://github.com/emccode/mars-challenge/`. You can perform a git clone: `git clone https://github.com/emccode/mars-challenge.git`
2. Navigate to the `/clients/python` folder
3. Edit the Team-client.py file. On lines 16 and 17, change the Values of the **IP address and ports** to the ones that you used for the Game Controller. These are the variables:
	- Line 16: `server_url = 'http://192.168.59.103:8080/api'`   # URL of the SERVER API 
	- Line 17: `server_ws = 'ws://192.168.59.103:8080/ws'`       # URL of the Sensors Websocket 
4. Save the File
5. Execute the following command: `sudo python Team-client.py'. The Program will wait until the game starts to register. 
6. To Start a Game in the Game Controller. This will require for you to send a curl command to the Game Controller API. I will require administrator rights. Use the authentication token that you defined when you executed the docker Game Controller docker container. In our example: the Admin Token is: **`-e ADMIN_TOKEN=1234`**
		
	    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
	    HTTP/1.1 200 OK
	    Date: Fri, 31 Jul 2015 09:30:48 GMT
	    Content-Length: 12
	    Content-Type: text/plain; charset=utf-8
	
	    Game started
	
	It will return a 400 if the game is already running:
	
	    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
	    HTTP/1.1 400 Bad Request
	    Content-Type: text/plain; charset=utf-8
	    Date: Fri, 31 Jul 2015 10:21:51 GMT
	    Content-Length: 44
	
	    Game is already started, not doing anything

7. To Reset a Game in the Game Controller. This will require for you to send a curl command to the Game Controller API. It requires administrator rights. Use the authentication token that you defined when you executed the docker Game Controller docker container. In our example: the Admin Token is: **`-e ADMIN_TOKEN=1234`**

	    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/reset
	    HTTP/1.1 200 OK
	    Date: Fri, 31 Jul 2015 09:51:37 GMT
	    Content-Length: 23
	    Content-Type: text/plain; charset=utf-8
	
	    Game reset successfully
	
	If the game is still running:
	
	    $ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/reset
	    HTTP/1.1 400 Bad Request
	    Content-Type: text/plain; charset=utf-8
	    Date: Fri, 31 Jul 2015 10:19:46 GMT
	    Content-Length: 38
	
	    Cannot reset game while it is running 


**Note:** Refer to the [Game controller Information page](https://github.com/emccode/mars-challenge/tree/master/game-controller "Game Controller information page") for more details on the Game Controller API and configuration. 


## System Tiers

Building the Mar's Challenge system is not an easy task. You and your team will find a lot of real word problems that plague our industry. You will have to determine what do you need to get the job done and what do you want to work for maximizing your chances of survival. 

The following is list of all the tiers in the solution with explanations for each of them:

### Tier 1 | Sensors Array ###

The Sensors Array is build using Raspberry Pis. Each member in the team will receive one. Your goal is to configure them to provide you with the atmospheric sensor information. One problem though, the sensor may go down due to radiation, you and your team may need to detect the shutdown and send a wake up call.

**Functionality:**
- Provides Temperature Information.
- Provides Radiation Information.
- Provides Solar Flare Information.
- End-Point for waking up the Sensor Array.

**Implementation:**
- Setup the Raspberry Pis.
- Configure the Raspberry Pi for Wi-Fi. 
- Install Golang.
- Execute the provided Mars Atmospheric simulator (Web Socket service application written in GO).
- Example of output: **[http://hackathon-engine.cloudapp.net/](http://hackathon-engine.cloudapp.net/ "http://hackathon-engine.cloudapp.net/")**.

Setting up the RaspberryPi with the weather simulator: **[Setting up the RaspBerry Pi](https://github.com/emccode/hackathon-mars/blob/master/documentation/Raspberry-Go-Weather-Simulator-Setup.md)**.

A container Implementation is located here: [https://hub.docker.com/r/emccode/mars-challenge-client/](https://hub.docker.com/r/emccode/mars-challenge-client/ "Sensor Client Service")

### Tier 2 | Sensor Gateway and Aggregation ###

The Sensors gateway collects all the data feeds from each of one of the Sensors. It will detect if any of the sensors has gone offline and send wake up calls to restart the system. It relays the sensor information to the  Aggregation tier.

**Functionality:**
- Relay information to the Aggregation Tier.
- Detect and Wake up Sensor Arrays that have gone offline (extra).

**Implementation:**
- Consume the Mars Atmospheric Simulator Web Sockets for each Raspberry Pi. This can be up to 5 devices. 
- Detect if one of the sensor feeds has gone down. Then it will send a request to start the service again.
- Relay the Data to the Aggregation Tier.

Implementations of the Gateway tier are located in this folder: [https://github.com/emccode/mars-challenge/tree/master/gateway](https://github.com/emccode/mars-challenge/tree/master/gateway)

A container with a working gateway can be found here: [https://hub.docker.com/r/emccode/mars-challenge-gateway-py/](https://hub.docker.com/r/emccode/mars-challenge-gateway-py/ "Gateway and Aggregator Service")




### Tier 3 | Data Repository Tier
The data Repository tier stores the data in the system. This tier can be implemented using any data repository.

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- Stores the data for further usage.
- Stores the Backup data.
- Stores the log data from all services.

**Implementation:**
- Needs to be a clustered service.
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.


### Tier 4 | Data Analysis Tier
The  data analysis tier takes the data from the Data Aggregation tier and determines what actions to take (Shields up/Down). Participant can use any technology they choose. The process can be automated or manual.  

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- May Shield Up/Down information to the MARS Challenge engine or delegate that functionality to the Monitoring and Control Dashboard Tier.

**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.


### Tier 5 | Team Client and Control Dashboard Tier
The  Team Client and Control dashboard is the core of the teams HQ operations. This service allows the team in Mars to view their sensor data, the analysis information, and shield status. In addition the dashboard needs to send the join/exit challenge  to the Mars Challenge engine Websocket.

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- Displays information for sensor data, data analysis and shield status.
- Display information about the team status on the MARS Challenge engine.


**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.

**Note:** Refer to the [Testing the Command and Control Center](#testing-the-command-and-control-center) section on this document for more information


### Tier 6 | Data Backup
The Data Backup tier involves taking the sensor data and backs it up on one(1) minute batches for future analysis. This is a bonus tier. 

**Functionality:**
- Receives per second Data from the Data Aggregation tier, a queue or the data repository and performs a 1 minute backup. the Gateway array and calculate the Average temperature and radiation for all the sensors.
- Stores the data in the Data Repository Tier.

**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.


## Points Table 

The Points table shows the points received for each one of the tasks that your team completes. Points are distributed per:

- Application Tier
- Deployment and Management 
- Completing the challenge
- Bonus Points

**Detailed Points Table is located [here](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")**