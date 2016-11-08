# EMC {code} Hackathon

[![Join the chat at https://gitter.im/emccode](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/emccode?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Stories in Ready](https://badge.waffle.io/emccode/mars-challenge.svg?label=ready&title=Ready)](http://waffle.io/emccode/mars-challenge)

**April 8th, 2016:** Mars Hackathon @ University of Utah

**November 10th 2015:** Mars Hackathon @ [LISA](https://www.usenix.org/conference/lisa15) 2015


**August 26th 2015:** Read about the results of our first **[Hackathon event at LinuxCon/ContainerCon](http://blog.emccode.com/2015/08/26/mars-challenge-hackathon-at-containerconlinuxcon-2015/ "Mars Challenge Hackathon @ LinuxCon/ContainerCon 2015")**

## The Mars Challenge

![Planet Mars](https://github.com/emccode/hackathon-mars/blob/master/documentation/images/mars-11608_640.jpg)


Welcome to the Mars Challenge, you and your team are about to embark into a 7 hours space adventure that will challenge your **developer, operational, data analysis and infrastructure** skills. The Challenge puts you in the arid lands of Mars:


*You and your team just landed in Mars. As you prepare your base of operations, massive Sun storms are coming your way. That was the last communication you had from Earth. Now connectivity to earth has been cut, and although the base has enough energy, your protective shields can only be running for a few minutes at the time, without recharging.*

*Your only chance of survival is to monitor the current temperature and radiation levels in the planet atmosphere to detect sun flares and activate your base shields for protection.*

*You only have a few hours, 7 of them, to implement a sensor array, build and deploy the monitoring application to engage/disengage your shields, then fine tune an algorithm based on your data analysis that decides when to charge your shields and when to engage them for protection. Will you and your team survive?*

You and your team will have at your disposal the necessary tools to survive and win the challenge, however you will need all wits and skills to work together and implement a solution that allows you to survive and compete against other teams to find out who would be the last survivor.


![Mars Challenge High Level Architecture](https://github.com/emccode/mars-challenge/blob/master/documentation/images/Mars-challenge-high-level-architecture.gif)


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

Mars Challenge is freely distributed under the [MIT License](http://codedellemc.github.io/sampledocs/LICENSE "LICENSE"). See LICENSE for details.
