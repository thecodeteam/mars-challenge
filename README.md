# EMC{code} Hackathon



## The Mars Challenge

![Planet Mars](https://github.com/emccode/hackathon-mars/blob/master/documentation/images/mars-11608_640.jpg)


You and your team just landed in Mars. As you prepare the base, massive Sun storms are coming your way. That was the last communication you had from Earth. Now connectivity to earth has been cut, and although the base has enough energy, your protective shields can only be running for a few minutes at the time, without recharging. 

Your only chance of survival is to monitor the current temperature and radiation levels in the planet atmosphere to detect sun flares and activate your base shields for protection.

You only have a few hours, 7 of them, to implement a sensor array, build and deploy the monitoring application to engage/disengage your shields, then fine tune an algorithm based on your data analysis that decides when to charge your shields and when to engage them for protection. Will you and your team survive? 


**Available hardware and software resources for each team:**
- 1 x Raspberry PI per Team Member (take away)
- 1 x WIFI Dongle per Raspberry (take away)
- 5 x VMs in a Cloud Provider
- 1 x EMC ECS Online Account per Team


**Challenge Requirements:**
- Sensors Measurements need to be taken as a time series.
- Deployment of tiers need to be implemented using Containers (not in the Raspberry Pi)
- Deployment of the application needs to be automated using Deployment Tools
- All data from sensors and services needs to be logged
- All data needs be backed on regular intervals
- Need to detect and calculate when to turn on the shields (refer to data table)


**Bonuses for:** 
- Build the System using the following Tiers/Services (this will be explain in more details):

	|Tier|Description|
	|----|-----------|
	|Tier 1| Sensors (Raspberry)|
	|Tier 2| Sensor Gateway (service)|
	Tier 3| Aggregation Tier (service)|
	Tier 4| Repository (DB, ECS, etc....)|
	Tier 5| Data Analysis (Service / Spark / Hadoop)|
	Tier 6| Monitoring and Control Dashboard Tier|
	Tier 7| Backup System|


**High Level Architecture:**

![High level architecture diagram](https://github.com/emccode/hackathon-mars/blob/master/documentation/images/marshackathon-Participant-diagram.jpg)



**You get Extra Points for using the following items in your team's implementation:**

- EMC ECS as a data repository
- Using Spark/Hadoop for data processing
- Survive Poor network communications between the Gateway Tier (Tier 2) and Aggregation Tier (Tier 3) using [Helios Burn](https://github.com/emccode/HeliosBurn "Helios Burn Fault Injection Platform")
- Running a Docker container in Raspberry Pi with the Sensor Software (provided Go application) 
- Work with limited resources (only drink water)
- Show historical Data in the Display Tier (Tier 6)


**What would you accomplish from participating on this Hackathon?**

This Hackathon will challenge you to accomplish the following: 

- Build a modern distributed application (Platform 3) using best practices
- Deploy a distributed application using Containers 
- Deploy a distributed application using modern deployment technologies (Messos, Kubernetes)
- Develop an application on the Internet of Things (IoT) space



## Mars Challenge Application Tiers Explanation

The following diagram shows all the tiers: 

![High level architecture diagram](https://github.com/emccode/hackathon-mars/blob/master/documentation/images/Mars-challenge-high-level-architecture.png )

### Tier 1 | Sensors Array ###

The Sensors Array is build using Raspberry Pis. Each member in the team will receive one. Your goal is to configure them to provide you with the atmospheric sensor information. One problem though, the sensor may go down due to radiation, you and your team need to detect the shutdown and send a wake up call.

**Functionality:**
- Provides Temperature Information
- Provides Radiation Information
- Provides Solar Flare Information
- End-Point for waking up the Sensor Array

**Implementation:**
- Setup the Raspberry Pis
- Configure the Raspberry Pi for Wi-Fi 
- Install Golang
- Execute the provided Mars Atmospheric simulator (Web Socket service application written in GO)
- Example of output: **[http://hackathon-engine.cloudapp.net/](http://hackathon-engine.cloudapp.net/ "http://hackathon-engine.cloudapp.net/")**

Setting up the RaspberryPi with the weather simulator: **[Setting up the RaspBerry Pi](https://github.com/emccode/hackathon-mars/blob/master/documentation/Raspberry-Go-Weather-Simulator-Setup.md)**


### Tier 2 | Sensor Gateway ###

The Sensors gateway collects all the data feeds from each of one of the Sensors. It will detect if any of the sensors has gone offline and send wake up calls to restart the system. It relays the sensor information to the  Aggregation tier

**Functionality:**
- Relay information to the Aggregation Tier
- Detect and Wake up Sensor Arrays that have gone offline (extra)

**Implementation:**
- Consume the Mars Atmospheric Simulator Web Sockets for each Raspberry Pi. this can be up to 5 devices. 
- Detect if one of the sensor feeds has gone down. Then it will send a request to start the service again
- Relay the Data to the Aggregation Tier.


### Tier 3 | Aggregation Tier
The Aggregation Tier is a service that will take the information provided by the Gateway, aggregates the data from all sensors into an AVG. This is AVG is used by the Data Analysis tier to perform the shields and re-charing logic

**Functionality**
- Receives per second Data from the Gateway array and calculate the Average temperature and radiation for all the sensors
- Stores the data on the Data Store
- Stores the data into a backup Service

**Implementation**
- Needs to be implemented as a service
- Service needs to be executed in a container
- Service needs to be deployed as part of the application


### Tier 4 | Data Repository Tier
The  data Repository tier stores the data in the system. This tier can be implemented using any data repository.

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- Stores the data for further usage
- Stores the Backup data
- Stores the log data from all services

**Implementation:**
- Needs to be a clustered service
- Needs to be implemented as a service
- Service needs to be executed in a container
- Service needs to be deployed as part of the application


### Tier 5 | Data Analysis Tier
The  data analysis tier takes the data from the Data Aggregation tier and determines what actions to take (Shields up/Down). Participant can use any technology they choose. The process can be automated or manual.  

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- May Shield Up/Down information to the MARS Challenge engine or delegate that functionality to the Monitoring and Control Dashboard Tier

**Implementation:**
- Needs to be implemented as a service
- Service needs to be executed in a container
- Service needs to be deployed as part of the application


### Tier 6 | Monitoring and Control Dashboard Tier
The  Monitoring dashboard is a service allows the team in Mars to view the sensor data, the analysis information, and shield status. In addition the dashboard needs to send the information to the 

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- Displays information for sensor data, data analysis and shield status
- Display information about the team status on the MARS Challenge engine.


**Implementation:**
- Needs to be implemented as a service
- Service needs to be executed in a container
- Service needs to be deployed as part of the application


### Tier 7 | Data Backup
The Data Backup tier takes care of taking the sensor data and back it up on 1 min batches for future analysis. This is a bonus tier 

**Functionality:**
- Receives per second Data from the Data Aggregation tier, a queue or the data repository and performs a 1 minute backup. the Gateway array and calculate the Average temperature and radiation for all the sensors
- Stores the data in the Data Repository Tier

**Implementation:**
- Needs to be implemented as a service
- Service needs to be executed in a container
- Service needs to be deployed as part of the application



## Points Table ##

This is the distribution of points for the Project: 

|Name|Tier|Description|Points|
|----|----|-----------|------|




 




