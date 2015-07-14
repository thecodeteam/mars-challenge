# EMCCode Hackaton #



## The Mars Challenge ##

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
- Build the System using the following Tiers/Services (this will be explain in more details):

	|Tier|Description|
	|----|-----------|
	|Tier 1| Sensors (Raspberry)|
	|Tier 2| Sensor Gateway (service)|
	Tier 3| Aggregator Tier (service)|
	Tier 4| Repository (DB, ECS, etc....)|
	Tier 5| Data Analysis / Data Aggregation (Service / Spark / Hadoop)|
	Tier 6| Display Sensor data,  Aggregate information, and shield status (UI)|
	Tier 7| Backup System|

- UI (Tier 6) to display information: Show Readings & Shield Status


**High Level Architecture**

![High level architecture diagram](https://github.com/emccode/hackathon-mars/blob/master/documentation/images/Mars-challenge-high-level-architecture.png )


**You get Extra Points for using the following items in your team's implementation:**

- EMC ECS as a data repository
- Using Spark/Hadoop for data processing
- Survive Poor network communications between the Gateway Tier (Tier 2) and Aggregation Tier (Tier 3) using [Helios Burn](https://github.com/emccode/HeliosBurn "Helios Burn Fault Injection Platform")
- Running a Docker container in Raspberry Pi with the Sensor Software (provided Go application) 
- Work with limited resources (only drink water)
- Show historical Data in the Display Tier (Tier 6)


**What would you accomplish from participating on this Hackaton?**

This Hackaton will challenge you to accomplish the following: 

- Build a modern distributed application (Platform 3) using best practices
- Deploy a distributed application using Containers 
- Deploy a distributed application using modern deployment technologies (Messos, Kubernetes)
- Develop an application on the Internet of Things (IoT) space



## Mars Challenge Application Tiers Explanation

### Tier 1 | Sensors Array ###

The Sensors Array is build using Raspberry Pis. Each member in the team will receive one. Your goal is to configure them to provide you with the atmospheric sensor information. 

**Functionality:**
- Provides Temperature Information
- Provides Radiation Information
- Provides Solar Flare Information
- End-Point for waking up the Sensor Array

**Implementation:**
- Setup the Raspberry Pis
- Configure the Raspberry Pi for WiFi 
- Install Golang
- Execute the provided Mars Atmospheric simulator (Web Socket service application written in GO)


#### Tier 2 | Sensor Gateway ###

The Sensors gateway collects all the data feeds from each of one of the Sensors. It will detect if any of the sensors has gone offline and send wake up calls to restart the system. It relays the sensor information to the  Aggregation tier

**Functionality:**
- Relay information to the Aggregation Tier
- Detect and Wake up Sensor Arrays that have gone offline (extra)

**Implementation:**


#### Tier 3 | Aggregator Tier ###

 
#### Tier 4 | Aggregator Tier ###


#### Tier 5 | Aggregator Tier ###


#### Tier 6 | Aggregator Tier ###


#### Tier 7 | Aggregator Tier ###



