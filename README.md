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
- Deployment of tiers need to be implemented using Containers (not in the Raspberry Pi).
- Deployment of the application needs to be automated using Deployment Tools.
- All data from sensors and services needs to be logged.
- All data needs be backed on regular intervals.
- Need to detect and calculate when to turn on the shields (refer to data table).


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

- EMC ECS as a data repository.
- Using Spark/Hadoop for data processing.
- Survive Poor network communications between the Gateway Tier (Tier 2) and Aggregation Tier (Tier 3) using [Helios Burn](https://github.com/emccode/HeliosBurn "Helios Burn Fault Injection Platform").
- Running a Docker container in Raspberry Pi with the Sensor Software (provided Go application).
- Work with limited resources (only drink water).
- Show historical Data in the Display Tier (Tier 6).


**What would you accomplish from participating on this Hackathon?**

This Hackathon will challenge you to accomplish the following: 

- Build a modern distributed application (Platform 3) using best practices.
- Deploy a distributed application using Containers. 
- Deploy a distributed application using modern deployment technologies (Messos, Kubernetes).
- Develop an application on the Internet of Things (IoT) space.




 




