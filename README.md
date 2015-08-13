# EMC{code} Hackathon

## The Mars Challenge

![Planet Mars](https://github.com/emccode/hackathon-mars/blob/master/documentation/images/mars-11608_640.jpg)


Welcome to the Mars Challenge, you and your team are about to embark into a 7 hours space adventure that will challenge your **developer, operational, data analysis and infrastructure** skills. The Challenge puts you in the arid lands of Mars: 


*You and your team just landed in Mars. As you prepare your base of operations, massive Sun storms are coming your way. That was the last communication you had from Earth. Now connectivity to earth has been cut, and although the base has enough energy, your protective shields can only be running for a few minutes at the time, without recharging.* 

*Your only chance of survival is to monitor the current temperature and radiation levels in the planet atmosphere to detect sun flares and activate your base shields for protection.*

*You only have a few hours, 7 of them, to implement a sensor array, build and deploy the monitoring application to engage/disengage your shields, then fine tune an algorithm based on your data analysis that decides when to charge your shields and when to engage them for protection. Will you and your team survive?*

You and your team will have at your disposal the necessary tools to survive and win the challenge, however you will need all wits and skills to work together and implement a solution that allows you to survive and compete against other teams to find out who would be the last survivor.


![Mars Challenge High Level Architecture](https://github.com/emccode/mars-challenge/blob/master/documentation/images/Mars-challenge-high-level-architecture.JPG)


**Available hardware and software resources for each team:**
- 1 x [Raspberry PI 2](http://www.amazon.com/CanaKit-Raspberry-Complete-Original-Preloaded/dp/B008XVAVAW/ref=sr_1_1?s=electronics&ie=UTF8&qid=1439267179&sr=1-1&keywords=raspberry+pi+2) per Team Member (take away)
- 1 x WIFI Dongle per Raspberry (take away)
- 5 x VMs in a Cloud Provider
- 1 x [EMC's ECS Test Drive](https://portal.ecstestdrive.com/) Account per Team
- 1 x [Redis Cloud](https://redislabs.com/pricing?service=redis) Account per Team


**Challenge Requirements:**
- Each Participant will need to bring their own Laptop/computer
- Deployment of system layers need to be implemented using Containers (not in the Raspberry Pi).
- Deployment of the system needs to be automated using Deployment Tools.
- All data from sensors and services needs to be logged.
- All data needs be backed on regular intervals.
- Team solutions need to connect to the Mars Challenge [Game controller](https://github.com/emccode/mars-challenge/tree/master/game-controller "Game Controller") 


**What would you accomplish from participating on this Hackathon?**

This Hackathon will challenge you to accomplish the following: 

- Build and deploy a modern distributed application using best practices.
- Deploy a distributed application using Containers. 
- Practice container configuration, execution and debugging. 
- Deploy a distributed application using modern deployment technologies like  Docker Swarm, Compose, Messos, and Kubernetes).
- Develop an application on the Internet of Things (IoT) space.

**For more details refer to the:[ Mars Challenge Instructions document](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-challenge-instructions.md "Mars challenge Instructions document")**


## License & Contributions

[Please follow this link for License and Contributions](https://github.com/emccode/mars-challenge/blob/master/documentation/Mars-Challenge-licence.md "License and Contributions")

 
## DevOps Tools 
 
### EMCCODE Build Server
- **Game Controller Build:** <a href="http://buildserver.emccode.com/viewType.html?-buildTypeId=MarsChallenge_GameControllerBuild&guest=1"><img src="http://buildserver.emccode.com/app/rest/builds/buildType:(id:MarsChallenge_GameControllerBuild)/statusIcon"/></a>
- **Sensor Service Build:** <a href="http://buildserver.emccode.com/viewType.html?buildTypeId=MarsChallenge_SensorGo_Main&guest=1"><img src="http://buildserver.emccode.com/app/rest/builds/buildType:(id:MarsChallenge_SensorGo_Main)/statusIcon"/></a>

### Code Coverage

- **Code Coverage(CodecovIO):** [![codecov.io](http://codecov.io/github/emccode/mars-challenge/coverage.svg?branch=master)](http://codecov.io/github/emccode/mars-challenge?branch=master)
- **Code Coverage(Coveralls):** [![Coverage Status](https://coveralls.io/repos/emccode/mars-challenge/badge.svg?branch=master&service=github)](https://coveralls.io/github/emccode/mars-challenge?branch=master)





