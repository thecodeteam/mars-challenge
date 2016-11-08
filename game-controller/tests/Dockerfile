# Base image
FROM ubuntu

# Maintainer info
MAINTAINER Patrick Butler Monterde <Patrick.butlermonterde@emc.com>

# Get the project
RUN apt-get update &&\
  apt-get install -y git-core python python-pip python-dev

RUN cd /opt &&\
  git clone https://github.com/codedellemc/mars-challenge &&\
  cd /opt/mars-challenge/game-controller/tests/

# Set the default directory where CMD will execute
WORKDIR /opt/mars-challenge/game-controller/tests

CMD python game-controller-harness.py
