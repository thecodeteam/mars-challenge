# Mars Challenge Instructions 














## Mars Challenge Application Tiers Explanation

The following diagram shows all the tiers: 

![High level architecture diagram](https://github.com/emccode/hackathon-mars/blob/master/documentation/images/Mars-challenge-high-level-architecture.png )

### Tier 1 | Sensors Array ###

The Sensors Array is build using Raspberry Pis. Each member in the team will receive one. Your goal is to configure them to provide you with the atmospheric sensor information. One problem though, the sensor may go down due to radiation, you and your team need to detect the shutdown and send a wake up call.

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


### Tier 2 | Sensor Gateway ###

The Sensors gateway collects all the data feeds from each of one of the Sensors. It will detect if any of the sensors has gone offline and send wake up calls to restart the system. It relays the sensor information to the  Aggregation tier.

**Functionality:**
- Relay information to the Aggregation Tier.
- Detect and Wake up Sensor Arrays that have gone offline (extra).

**Implementation:**
- Consume the Mars Atmospheric Simulator Web Sockets for each Raspberry Pi. this can be up to 5 devices. 
- Detect if one of the sensor feeds has gone down. Then it will send a request to start the service again.
- Relay the Data to the Aggregation Tier.


### Tier 3 | Aggregation Tier
The Aggregation Tier is a service that will take the information provided by the Gateway, aggregates the data from all sensors into an AVG. This is AVG is used by the Data Analysis tier to perform the shields and re-charing logic.

**Functionality**
- Receives per second Data from the Gateway array and calculate the Average temperature and radiation for all the sensors.
- Stores the data on the Data Store.
- Stores the data into a backup Service.

**Implementation**
- Needs to be implemented as a service
- Service needs to be executed in a container
- Service needs to be deployed as part of the application


### Tier 4 | Data Repository Tier
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


### Tier 5 | Data Analysis Tier
The  data analysis tier takes the data from the Data Aggregation tier and determines what actions to take (Shields up/Down). Participant can use any technology they choose. The process can be automated or manual.  

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- May Shield Up/Down information to the MARS Challenge engine or delegate that functionality to the Monitoring and Control Dashboard Tier.

**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.


### Tier 6 | Monitoring and Control Dashboard Tier
The  Monitoring dashboard is a service allows the team in Mars to view the sensor data, the analysis information, and shield status. In addition the dashboard needs to send the join/exit challenge  to the Mars Challenge engine websocket.

**Functionality:**
- Receives sensor data from the Data Aggregation tier.
- Displays information for sensor data, data analysis and shield status.
- Display information about the team status on the MARS Challenge engine.


**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.


### Tier 7 | Data Backup
The Data Backup tier takes care of taking the sensor data and back it up on 1 min batches for future analysis. This is a bonus tier. 

**Functionality:**
- Receives per second Data from the Data Aggregation tier, a queue or the data repository and performs a 1 minute backup. the Gateway array and calculate the Average temperature and radiation for all the sensors.
- Stores the data in the Data Repository Tier.

**Implementation:**
- Needs to be implemented as a service.
- Service needs to be executed in a container.
- Service needs to be deployed as part of the application.



## Points Table ##

This is the distribution of points for the Project: 

|Name|Tier|Description|Points|
|----|----|-----------|------|


