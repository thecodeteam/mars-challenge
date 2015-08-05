# Base image
FROM ubuntu

# Maintainer info
MAINTAINER Patrick Butler Monterde <Patrick.butlermonterde@emc.com>

# Get the project
RUN apt-get update &&\
  apt-get install -y git-core

RUN cd /opt &&\
  git clone https://github.com/emccode/mars-challenge &&\
  cd /opt/mars-challenge/gateway/

# Set the default directory where CMD will execute
WORKDIR /opt/mars-challenge/gateway

CMD python gateway-python.py
