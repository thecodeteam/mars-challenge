# Mars Challenge Instructions

## Contents

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

Welcome to the Mars Challenge! You and your team are about to embark on a
space adventure that will challenge your **developer, operator, data analysis
and infrastructure** skills. The Challenge puts you in the arid lands of Mars:

*You and your team just landed on Mars. As you prepare your base of operations,
you receive word that massive Sun storms are coming your way. Now radio contact
with Earth has been lost. Your base has protective electromagnetic shields that
can protect you from the radiation, but can only be running for a few minutes
at a time without recharging.*

*Your only chance of survival is to monitor the current temperature and
radiation levels in the planet's atmosphere to detect sun flares and activate
your base shields for protection.*

*You only have a few hours to implement a sensor array, build and deploy the
monitoring application to engage/disengage your shields, then fine tune an
algorithm based on your data analysis that decides when to charge your shields
and when to engage them for protection. Will you and your team survive?*

You and your team will have at your disposal the necessary tools to survive and
win the challenge, however you will need all wits and skills to work together
and implement a solution that allows you to survive and outlast other teams.

## Goals

The Mars Challenge was built by people passionate about technology, for people
just as passionate or more. The goals for the challenge are the following:

- Network with people that share the same interests as you do
- Gather on teams to accomplish something fun and learn something while doing it
- Build and deploy a modern distributed application using best practices.
- Deploy a distributed application using containers.
- Practice container configuration, execution and debugging.
- Deploy a distributed application using modern deployment technologies like
  Docker Swarm, Compose, Mesos, or Kubernetes.
- Develop an application on the Internet of Things (IoT) space.

## Requirements

This Challenge requires the following or their participants:

- Each Participant needs to bring their own computer/laptop
- Have a good attitude and plan to have serious technical fun
- Be a Team player
- Adhere to the LinuxCon/ContainerCon/MesosCon Code of Conduct

## Winning

The Challenge is designed to be completed in the time allotted. Each team will
have to choose which features [(tasks)](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md)
to implement and deploy as part of their solution. **The only condition** is to
have a working solution that satisfies task **[CC-1](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md)** (demonstrate participation in an official Game).
Refer to the [Testing the Command and Control Center](#testing-the-command-and-control-center) section to get this part running.

**[Each task your team accomplishes has a point value](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")**

For each task, functionality needs to be successfully demoed to a Judge, then
points will be assigned. Point count for each team will be visible for the
entire length of the contest.

**The team that has the most amount of points at the end wins the challenge.**

### End of Challenge Shootout

At the conclusion of the Challenge, a best of 3 shootout is held between all
teams to see which algorithm lasts the longest on Mars. This is task
**[CC-2](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md)**
The winner of the shootout may not be the over all Challenge winner!


## Challenge Assets

For the Challenge we are providing most of the services, the codebases, and
associated containers for the tiers that make the solution.


### Services and Codebases locations

These are the locations for code samples and information for each one of the
services:

|Service Name|Folder Location|
|----|----|
|SensorSuite Service| [https://github.com/ghostplant/mars-challenge/tree/master/sensorsuite](https://github.com/ghostplant/mars-challenge/tree/master/sensorsuite)|
|Aggregator Service| [https://github.com/ghostplant/mars-challenge/tree/master/aggregator](https://github.com/ghostplant/mars-challenge/tree/master/aggregator)|
|Data Repository| Refer to the [System Tiers](#system-tiers) section for more information . |
|Data Analysis | Refer to the [System Tiers](#system-tiers) for more information.|
|Team Command & Control| [https://github.com/ghostplant/mars-challenge/tree/master/clients/python](https://github.com/ghostplant/mars-challenge/tree/master/clients/python)|
|Game Controller| [https://github.com/ghostplant/mars-challenge/tree/master/game-controller](https://github.com/ghostplant/mars-challenge/tree/master/game-controller)|
|Game Controller Dashboard| [https://github.com/ghostplant/mars-challenge/tree/master/dashboard](https://github.com/ghostplant/mars-challenge/tree/master/dashboard)|
|RaspBerry Pi Sensor Setup| [https://github.com/ghostplant/mars-challenge/blob/master/documentation/Raspberry-Go-Weather-Simulator-Setup.md](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Raspberry-Go-Weather-Simulator-Setup.md)|

### Containers

These are the containers available for all the teams to use. These containers
are the same images that will be used when running the end of game shootout.

|Service Name|Container Location|
|----|----|
|Aggregator Service| [https://hub.docker.com/r/emccode/mars-challenge-aggregator/](https://hub.docker.com/r/emccode/mars-challenge-aggregator/ "Aggregator Service")|
|Game Challenge Controller| [https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard| [https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

## Participant's Implementation

Participants have an opportunity to build multiple tiers of functionality. These
have been organized into the following groups:

 - Section 1: Getting Data From the Sensors to the Aggregator Service
 - Section 2: Perform Analytic and Data saving operations
 - Section 3: DevOps and Application Management

###  Section 1: Getting Data From the SensorSuite to the Aggregator Service

This section focuses on getting the sensor data from the SensorSuite to the
Aggregator Service, which then sends an aggregated stream to the Game
Controller.

 ![Mars Challenge Participants Stage 1](https://github.com/ghostplant/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage1.png).

In order to implement section 1 you will need the following components:

|Service Name|Container Location|
|----|----|
|SensorSuite Service (Tier 1)| [https://github.com/ghostplant/mars-challenge/tree/master/sensorsuite](https://github.com/ghostplant/mars-challenge/tree/master/sensorsuite)|
|Aggregator Service (Tier 2)| [https://hub.docker.com/r/emccode/mars-challenge-aggregator/](https://hub.docker.com/r/emccode/mars-challenge-aggregator/ "Aggregator Service")|
|Game Challenge Controller| [https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard| [https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

A Demo implementations of a Command and Control (Tier 5) Service, implemented in
Python, is located in the following folder: [https://github.com/ghostplant/mars-challenge/tree/master/clients/python](https://github.com/ghostplant/mars-challenge/tree/master/clients/python "Command and Control Demo")

Once you have the components in place, you and your team will have to create a
deployment and orchestrate how services are deployed in a specific order:

1. Game Controller
2. Dashboard
3. Sensor Suite's Solar Flare
4. Aggregator Service
5. Sensor Service
6. Team Command and Control Center service

**NOTE:** The components you deploy for early testing and working on your Team
Solution will be different than what you deploy for the end of game shootout.
To test and iterate, you will deploy all components yourself, outlined above.
At the end of the game, the judges will have deployed a central solar flare
component, an Aggregator, the Game Controller and Dashboard.

You will get [points](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")
for automating the deployment and orchestration of this section.

### Section 2: Perform Analytic and Data Saving Operations

**Section 2:** This section focuses on getting the rest of services that work
with the Team Command and Control center:

![Mars Challenge Participants Stage 2](https://github.com/ghostplant/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage2.png)

The first step is to make sure you have a Command and Control service
working. This is algorithm you create for deciding when to raise your shields
and when to charge them. You can test your implementation of the Command &
Control (C&C) service using the information in
[Testing the Command and Control Center](#testing-the-command-and-control-center)
section.

Once you have the C&C service ready and on a container, you can start adding the
other pieces that you may need. These are some of the options of the services
you can implement:

- Data Repository Tier
- Data Analytic Tier
- Backup Service

Please Refer to the **[System Tiers](#system-tiers)** section for more details
on these. Remember, each one of these tier will get you more [points](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")
and at the same time it will increase the complexity of the deployments.


### Section 3: DevOps and Application Management

Even on Mars, you will need to deploy your application, in your infrastructure
to get to the functionality you and your team have implemented. The goal is for
you and your team to get experience managing and deploying micro services
applications. Some important aspects of managing micro services are:

- **Service Discovery:** Implement Service Discovery between all the containers
  deployed in the solution.
- **Service Monitoring:** Implement Service Monitoring of all containers
  deployed in the solution.
- **Service Configuration:** Implement Service configuration for all the
  Tiers/containers.
- **Service Orchestration:** Implement Service Orchestration for all the
  Tiers/containers.
- **Automated Service Deployment:** Deployment of the implemented system using
  Docker Tooling, Swarm, Kubernetes, Mesos, Puppet/Ansible/Chef/Saltstack or similar.
- **Logging Router:** Deploy, connect and route the logs of the application
  using a logging router.

**Note:** We are providing VMs on a Cloud provider for you to do deploy your
system. You can use any tooling to automate this process.

## Testing the Command and Control Center

In order to Test your Team Command and Control center you can run the Mars Game
Controller in testing mode. This disables the Sensor API and causes the Game
Controller to locally generate sensor readings instead. You will need to run the
Game Controller and Game Controller's Dashboard.

![Testing the Command and Control Center](https://github.com/ghostplant/mars-challenge/blob/master/documentation/images/Mars-challenge-testing-control-center.png)

The Game containers are located here:

|Service Name|Container Location|
|----|----|
|Game Controller| [https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard| [https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

The provided containers are Docker containers. You need to have Docker installed
on the host(s) that you plan to run these on. You should start up the containers
in the following order:

1. Game Controller
2. Game Controller Dashboard
3. Team Command and Control Service

Please use the following commands to setup the containers:

First, start by executing the Game Controller Container. You can define both the
listening port: `-p 8080:8080` and the Admin Token: **`-e ADMIN_TOKEN=1234`** to
what you may need:

```sh
docker run -d --name=controller -p 8080:8080 -e ADMIN_TOKEN=1234 emccode/mars-challenge-controller
```

You can check the Game Controller service is running using a web browser against
the URL: `http://<host ip address>:8080`.

Second, execute the Game Controller Dashboard and map it to the endpoint used by
the Game Controller:

```
docker run -d --name dashboard -e WS_ENDPOINT=localhost:8080/ws -p 80:80 emccode/mars-challenge-dashboard
```

Where `-e WS_ENDPOINT=localhost:8080/ws` is the location of the Game Controller
and `-p 80:80` is the endpoint of the Game Controller Dashboard. You can check
the Game Controller Dashboard service running on the specified address. In our
example, the address would be: `http://<host ip address>:80`.

The last step is to execute the Command and Control Center code. You can run the
code directly or run it as a container.

A Demo of a Command and Control (Tier 5) Center, implemented in Python, is
provided and is located in the following folder: [https://github.com/ghostplant/mars-challenge/tree/master/clients/python](https://github.com/ghostplant/mars-challenge/tree/master/clients/python "Command and Control Demo")

The Demo implements all the interfaces provided by the Game Controller and
implements shield operations based on sensor data.

For our example we will run the code directly from command line/terminal:

1. Get the code from the code repository:
   `https://github.com/ghostplant/mars-challenge/`.
    You can perform a git clone:
    `git clone https://github.com/ghostplant/mars-challenge.git`
2. Navigate to the `/clients/python` folder
3. Edit the `Team-client.py` file. On lines 16 and 17, change the Values of the
   **IP address and ports** to the ones that you used for the Game Controller.
   These are the variables:
	- Line 16: `server_url = 'http://192.168.59.103:8080/api'`   # URL of the SERVER API
	- Line 17: `server_ws = 'ws://192.168.59.103:8080/ws'`       # URL of the Sensors Websocket
4. Save the File
5. Execute the following command:
   `sudo python Team-client.py`. The Program will wait until the game starts to
   register.
6. To Start a Game in the Game Controller. This will require for you to send a
   curl command to the Game Controller API. It will require administrator
   rights. Use the authentication token that you defined when you executed the
   Game Controller Docker container. In our example: the Admin Token is:
   **`-e ADMIN_TOKEN=1234`**

	```
	$ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
	HTTP/1.1 200 OK
	Date: Fri, 31 Jul 2015 09:30:48 GMT
	Content-Length: 12
	Content-Type: text/plain; charset=utf-8

	    Game started
	```

	It will return a 400 if the game is already running:

	```
	$ curl -i -H 'X-Auth-Token: 1234' -X POST http://localhost:8080/api/start
	HTTP/1.1 400 Bad Request
	Content-Type: text/plain; charset=utf-8
	Date: Fri, 31 Jul 2015 10:21:51 GMT
	Content-Length: 44

	Game is already started, not doing anything
	```

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


**Note:** Refer to the [Game controller Information page](https://github.com/ghostplant/mars-challenge/tree/master/game-controller "Game Controller information page") for more details on the Game Controller API and configuration.

## System Tiers

Building the Mars Challenge system is not an easy task. You and your team will
find a lot of real word problems that plague our industry. You will have to
determine what do you need to get the job done and what do you want to work for
maximizing your chances of survival.

The following is a list of all the tiers in the solution with explanations for
each of them.

### Tier 1 | Sensors Array ###

The Sensors Array is built using Raspberry Pis. Each team will be assigned one
for use. Your goal is to configure them to provide you with the atmospheric
sensor information.

**Functionality:**
- Provides Temperature Information.
- Provides Radiation Information.
- Receives Solar Flare Information.

**Implementation:**
- Connect to the Raspberry Pi.
- Execute the provided SensorSuite.
  (Web Socket service application written in GO).
- Configure SensorSuite to send data to Sensor Aggregator

**Note** You can execute the SensorSuite applications directly, or you can
deploy them in containers. There are points associated with each choice.

### Tier 2 | Sensor Aggregation ###

The Sensors Aggregator collects the data feeds from each of one of the Sensor
Arrays. It averages the temperature and radiation values, and relays that data
to the Game Controller.

**Functionality:**
- Receive, aggregate, and relay information to the Game Controller.

**Implementation:**
- Expose endpoint that each SensorSuite **Publisher** can connect to
- Receive and average data from all Sensor Arrays
- Relay the Data to the Game Controller.

Implementation of the Aggregator tier is located in this folder: [https://github.com/ghostplant/mars-challenge/tree/master/aggregator](https://github.com/ghostplant/mars-challenge/tree/master/aggregator)

A container with a working aggregator can be found here: [https://hub.docker.com/r/emccode/mars-challenge-aggregator/](https://hub.docker.com/r/emccode/mars-challenge-aggregator/ "Aggregator Service")

### Tier 3 | Data Repository Tier
The data Repository tier stores the data in the system. This tier can be
implemented using any data repository.

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
The  data analysis tier takes the data from the Data Aggregation tier and
determines what actions to take (Shields up/Down). Participant can use any
technology they choose. The process can be automated or manual.  

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- May send Shield Up/Down information to the MARS Challenge engine or delegate
that functionality to the Monitoring and Control Dashboard Tier.

**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.

### Tier 5 | Team Command and Control Tier
The Team Command and Control service is the core of your team's operations.
This service allows the team on Mars to view their sensor data, the analysis
information, and shield status. In addition the service needs to send the
join/exit challenge to the Mars Challenge engine Websocket.

**Functionality:**
- Receives sensor data from the Game Controller.
- Displays information for sensor data, data analysis and shield status.
- Display information about the team status on the MARS Challenge engine.
- Implements shield control algorithm

**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.

**Note:** Refer to the [Testing the Command and Control Center](#testing-the-command-and-control-center) section on this document for more information

### Tier 6 | Data Backup
The Data Backup tier involves taking the sensor data and backs it up on one (1)
minute batches for future analysis. This is a bonus tier.

**Functionality:**
- Receives per second Data from the Game Controller, a queue, or the data
  repository and performs a 1 minute backup.
- Stores the data in the Data Repository Tier.

**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.

## Points Table

The Points table shows the points received for each one of the tasks that your
team completes. Points are distributed per:

- Application Tier
- Deployment and Management
- Completing the challenge
- Bonus Points

**Detailed Points Table is located [here](https://github.com/ghostplant/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")**
