
## Points Distribution Table ##

For each one of the tasks completed your team will get points. These Points will be added at the end of the Challenge. The team that has the most amount of points will win the challenge: 

- [Challenge Competition](#Challenge-Competition)
- [Points per Tier](#TIER-1-|-Sensors-Array) 
- [Bonus Points](#BONUS-Challenge)
- [DevOps Points](#Dev/Ops:) 



**Challenge Competition:**

|####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|CC-1|Competitor|CC| Compete against other teams in  Official Game| 30| 
|CC-2|Challenge Champion|CC| Winning the an official competition| 20| 



**TIER 1 | Sensors Array:**

|####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|T1-1 |Sensor Software Running in Raspberry|Tier 1|Get the Sensor program running in the Raspberry Pi|10|
|T1-2 |My Sensor Containers|Tier 1| Build a container with the Sensor Software and Deploy it on DockerHub.com. Use this container for the rest of the challenge|5|


**TIER 2 | The Gateway/Aggregation:**

|####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|T2-1|Multi-Sensors Gateway|Tier 2| Build a gateway that can read from multiple Sensors and can calculate the AVG between the sensor Data|5|
|T2-2|Data smarts| Tier 2| Build functionality on the gateway where the gateway sends all the RAW sensor and AVG data to a storage repository|5|
|T2-3*|Repair Crew|Tier 2| Build functionality on the gateway to detect and send wake up calls to sensors that fail|5|
|T2-4|My Gateway Container|Tier 2| Build a container with the Gateway Software and Deploy it on DockerHub.com. Use this container for the rest of the challenge|5|


**TIER 3 | Data Repository:**

|####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|T3-1 |Data Hog|Tier 3|Store all data from system: Sensor Data, Logs, etc...|10|
|T3-2 |The Black Box|Tier 3| Create a separate storage for data backups.Needs to be durable at least 3 copies. |5|


**TIER 4| Data Analysis Tier**

|####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|T4-1|Analytic Brain|Tier 4| Setup a cluster of Haddop/Spark/F# that to analyze data. This tier will help you to better anticipate change.|15|
|T4-2|My Analytic Container|Tier 4| Build a container with the Data Analysis tier and deploy it on DockerHub.com. Use this container for the rest of the challenge|5|


**Tier 5 | Team Client and Control Dashboard Tier:**

|#####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|T5-1|Command and Control|Tier 5| Implement a service that can retrieve data from the Gateway/Game Controller, join the game, and use the sensor data to take actions to protect the base|5|
|T5-2|My C&C Container|Tier 5| Build a container with the C&C tier and deploy it on DockerHub.com. Use this container for the rest of the challenge|5|


**Tier 6 | Data Backup:**

|#####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|T6-1|For the Ages|Tier 6|Build a Backup service that takes data and performs backups and stores them in Tier 3 Data repository. Can be combined with the "The Black Box" for added points|5|
|T6-2|My Backup Service Container |Tier 6| Build a container with the Data Backup service tier and deploy it on DockerHub.com. Use this container for the rest of the challenge|5|

**BONUS Challenge:**

|####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|BC-1|DevOps Power|BC| Deploy a CD/CI tool-chain, DevOps style, that includes Tiers 1,2,3,4, and 5 so your team can develop the algorithm to survive the event|20|
|BC-2|Limited Resources|BC| The team commits to only drink water for the entire challenge. It is Mars after all!|10|
|BC-3|The FOSS|BC| The team commits to share all the information about what you are doing. Free and open-source software works even in Space!|10|
|BC-4|Evil Genius|BC| The team commits to not sharing any information with anybody. You and your team are Evil Geniuses!|10|


**Dev/Ops:**

|#####|Name|Tier|Description|Points|
|----|----|----|-----------|------|
|DO-1|Service Discovery|DO| Implement Service Discovery between all the containers deployed in the solution| 5|
|DO-2|Service Monitoring|DO| Implement Service Monitoring of all containers deployed in the solution| 5|
|DO-3|Service Configuration|DO| Implement Service configuration for all the Tiers/containers| 5|
|DO-4|Service Orchestration|DO| Implement Service Orchestration for all the Tiers/containers| 5|
|DO-5|Docker Me up!|DO| Deployment of the implemented system using Docker Tooling| 15|
|DO-6|Kubernetes Me up!|DO| Deployment of the implemented system using Kubernetes Tooling| 15|
|DO-7|Messos Me up!|DO| Deployment of the implemented system using Messos Tooling| 15|
|DO-8|Deploy Me up!|DO| Deployment of the implemented system using Puppet/Ansible/Chef/Saltstack or other Tooling| 10|





