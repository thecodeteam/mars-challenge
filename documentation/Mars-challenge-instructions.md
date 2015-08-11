# Mars Challenge Instructions 


### Contents

- [Introduction](#Introduction)
- [Goals](#Goals)
- [Requirements](#requirements)
- [Winning](#winning)
- [Challenge Phases](#challenge-phases)
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
- Deploy a distributed application using modern deployment technologies like  Docker Swarm, Compose, Messos, and Kubernetes).
- Develop an application on the Internet of Things (IoT) space.

## Requirements

This Challenge requires the following for their participants:

- Each Participant needs to bring their own computer/laptop
- Have a good attitude and plan to have serious technical fun
- Be a Team player 
- Adhere to the LinuxCon/ContainterCon/MessosCon code of Conduct

## Winning

[Each task you accomplish has a point value](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table"). The Team that has the most amount of points by the end of the 7 hours wins the challenge. 


# Challenge Phases

The Challenge is designed to be completed in 7 hours. is organized 





![Mars Challenge Participants Stage 1](https://github.com/emccode/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage1.jpg)


![Mars Challenge Participants Stage 2](https://github.com/emccode/mars-challenge/blob/master/documentation/images/marshackathon-Participant-stage2.jpg)




## Mars System's Tiers

Building the Mar's Challenge system is not an easy task. You and your team will find a lot of real word problems that plague our industry. You will have to determine if wh 

### Tier 1 | Sensors Array ###

The Sensors Array is build using Raspberry Pis. Each member in the team will receive one. Your goal is to configure them to provide you with the atmospheric sensor information. One problem though, the sensor may go down due to radiation, you and your team mayneed to detect the shutdown and send a wake up call.

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

**Detailed Points Table is located [here](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-points-table.md "Mars Challenge Points Table") **