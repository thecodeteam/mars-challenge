# Mars Challenge Instructions 


### Contents

- [Introduction](#Introduction)
- [Goals](#Goals)
- [Requirements](#requirements)
- [Winning](#winning)
- [Challenge Phases](#challenge-phases)
- [Testing the Command and Control Center](#testing-the-command-and-control-center)
- [System Tiers](#system's-tiers)
- [Points Table](#points-table)


## Introduction 

Welcome to the Mars Challenge, you and your team are about to embark into a 7 hours space adventure that will challenge your **developer, operational, data analysis and infrastructure** skills. The Challenge puts you in the arid lands of Mars: 


*You and your team just landed in Mars. As you prepare your base of operations, massive Sun storms are coming your way. That was the last communication you had from Earth. Now connectivity to earth has been cut, and although the base has enough energy, your protective shields can only be running for a few minutes at the time, without recharging.* 

*Your only chance of survival is to monitor the current temperature and radiation levels in the planet atmosphere to detect sun flares and activate your base shields for protection.*

*You only have a few hours, 7 of them, to implement a sensor array, build and deploy the monitoring application to engage/disengage your shields, then fine tune an algorithm based on your data analysis that decides when to charge your shields and when to engage them for protection. Will you and your team survive?*

You and your team will have at your disposal the necessary tools to survive and win the challenge, however you will need all wits and skills to work together and implement a solution that allows you to survive and compete against other teams to find out who would be the last survivor.


## Goals 

This challenge was build by people passionate about technology for people as passionate or more about technology. The goals for the challenge are the following: 

- Network with people that share the same Passions as you do
- Gather on teams to accomplish something fun and learn something while doing it
- Build and deploy a modern distributed application application using best practices.
- Deploy a distributed application using Containers. 
- Practice container configuration, execution and debugging. 
- Deploy a distributed application using modern deployment technologies like  Docker Swarm, Compose, Messos, and Kubernetes).
- Develop an application on the Internet of Things (IoT) space.

## Requirements

This Challenge requires the following for their participants:

- Each Participant needs to bring their own computer/laptop
- Have a good attitude and plan to have serious technical fun
- Be a Team player 
- Adhere to the LinuxCon/ContainterCon/MessosCon code of Conduct

## Winning

The Challenge is designed to be completed in 7 hours. Each team will have to choose which features [(tasks)](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "List of Tasks") wants to implement and deploy as part of their solution. The only condition is to have a working solution that satisfies task [CC-1](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md) (been able to participate in an official Game). 

**[Each task your team accomplishes has a point value](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")**

For each task, functionality needs to be successfully demoed to a Judge, then points will be assigned. Point count for each team will be visible for the entire lenght of the contest.

**The team that has the most amount of points by the end of the seven hours wins the challenge.** 


## Challenge Approach

For this challenge we are providing most of the services, the codebases, and associated containers for the tiers that make the solution. 

You can use the implementation and containers provided or you can choose to build your own. There are [points](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table") associated for each decision that you make. 


These are the locations for each one o code that you can use to build your own containers

|Service Name|Folder Location|
|----|----|
|Sensor Client Service (Tier 1)|https://github.com/emccode/mars-challenge/tree/master/sensor-client|
|Gateway and Aggregator Service (Tier 2)|https://github.com/emccode/mars-challenge/tree/master/gateway|
|Data Repository (Tier 3)| |
|Data Analysis  (Tier 4)| |
|Team Command & Control (Tier 5)|https://github.com/emccode/mars-challenge/tree/master/clients/python|
|Game Controller|https://github.com/emccode/mars-challenge/tree/master/game-controller|
|Game Controller Dashboard|https://github.com/emccode/mars-challenge/tree/master/dashboard|


These are the containers avaliable for all the teams to use: 

|Service Name|Container Location|
|----|----|
|Sensor Client Service (Tier 1)| [https://hub.docker.com/r/emccode/mars-challenge-client/](https://hub.docker.com/r/emccode/mars-challenge-client/ "Sensor Client Service")|
|Gateway and Aggregator Service (Tier 2)| [https://hub.docker.com/r/emccode/mars-challenge-gateway-py/](https://hub.docker.com/r/emccode/mars-challenge-gateway-py/ "Gateway and Aggregator Service")|
|Game Challenge Controller|[https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard|[https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|


 We have divided the challenge in two sections to make things more comprehensive: 

**Section 1:** This section focuses on getting the sensor data from the Sensor services to the Command and Control Tier.  

 ![Mars Challenge Participants Stage 1](https://github.com/emccode/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage1.jpg)


|Service Name|Container Location|
|----|----|
|Sensor Client Service (Tier 1)| [https://hub.docker.com/r/emccode/mars-challenge-client/](https://hub.docker.com/r/emccode/mars-challenge-client/ "Sensor Client Service")|
|Gateway and Aggregator Service (Tier 2)| [https://hub.docker.com/r/emccode/mars-challenge-gateway-py/](https://hub.docker.com/r/emccode/mars-challenge-gateway-py/ "Gateway and Aggregator Service")|
|Game Challenge Controller|[https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard|[https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

A Demo implementations of a Command and Control (Tier 5) Service, implemented in Python, is located in the following folder: [https://github.com/emccode/mars-challenge/tree/master/clients/python](https://github.com/emccode/mars-challenge/tree/master/clients/python "Command and Control Demo")




![Mars Challenge Participants Stage 2](https://github.com/emccode/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage2.jpg)


## Testing the Command and Control Center

In order to Test your Team Command and Control center you can run the Mar's Game Controller in testing mode. You will need the to run the Game Controller and Game Controller's Dashboard. 

![Testing the Command and Control Center](https://github.com/emccode/mars-challenge/blob/master/documentation/images/Mars-challenge-testing-control-center.JPG)


The Game containers are located here: 

|Service Name|Container Location|
|----|----|
|Game Challenge Controller|[https://hub.docker.com/r/emccode/mars-challenge-controller/](https://hub.docker.com/r/emccode/mars-challenge-controller/ "Game Challenge Controller")|
|Controller Dashboard|[https://hub.docker.com/r/emccode/mars-challenge-dashboard/](https://hub.docker.com/r/emccode/mars-challenge-dashboard/ "Controller Dashboard")|

Use the following commands to setup the containers: 

 
First, start by executing the Game Controller Mode for Testing purposes:

    docker run -d --name=controller -p 80:8080 -e ADMIN_TOKEN=1234 emccode/mars-challenge-controller

Then execute the Game Controller Dashboard and map it to the endpoint used by the Game Controller: 

    docker run -d --name dashboard -e WS_ENDPOINT=localhost:80/ws -p 82:80 emccode/mars-challenge-dashboard

Where `-e WS_ENDPOINT=localhost:80/ws` is the location of the Game Controller and  `-p 82:80` is the endpoint of the Game Controller Dashboard.

A Demo of a Command and Control (Tier 5) Center, implemented in Python, is provided and  is located in the following folder: [https://github.com/emccode/mars-challenge/tree/master/clients/python](https://github.com/emccode/mars-challenge/tree/master/clients/python "Command and Control Demo")

The Demo implements all the interfaces provides by the Game Controller and implements Shield operations based on sensor data. 

**Note:** Refer to the [Game controller Information page](https://github.com/emccode/mars-challenge/tree/master/game-controller "Game Controller information page") for more details on the Game Controller API and configuration. 



## Mars System's Tiers

Building the Mar's Challenge system is not an easy task. You and your team will find a lot of real word problems that plague our industry. You will have to determine what do you need to get the job done and what do you want to work for maximizing your chances of survival. 

This is a list of all the tiers that a solution  do you want to build and 

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


### Tier 2 | Sensor Gateway and Aggregation ###

The Sensors gateway collects all the data feeds from each of one of the Sensors. It will detect if any of the sensors has gone offline and send wake up calls to restart the system. It relays the sensor information to the  Aggregation tier.

**Functionality:**
- Relay information to the Aggregation Tier.
- Detect and Wake up Sensor Arrays that have gone offline (extra).

**Implementation:**
- Consume the Mars Atmospheric Simulator Web Sockets for each Raspberry Pi. this can be up to 5 devices. 
- Detect if one of the sensor feeds has gone down. Then it will send a request to start the service again.
- Relay the Data to the Aggregation Tier.

Implementation of the Gateway tier are located in this folder: 


### Tier 3 | Data Repository Tier
The  data Repository tier stores the data in the system. This tier can be implemented using any data repository.

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
The  Team Client and Control dashboard is the core of the teams HQ operations. This service allows the team in Mars to view their sensor data, the analysis information, and shield status. In addition the dashboard needs to send the join/exit challenge  to the Mars Challenge engine websocket.

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- Displays information for sensor data, data analysis and shield status.
- Display information about the team status on the MARS Challenge engine.


**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.


### Tier 6 | Data Backup
The Data Backup tier takes care of taking the sensor data and back it up on 1 min batches for future analysis. This is a bonus tier. 

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
- Completing the scenarios
- Bonus Points

**Detailed Points Table is located [here](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table")**