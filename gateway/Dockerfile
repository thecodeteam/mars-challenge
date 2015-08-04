# Base image
FROM ubuntu

# Maintainer info
MAINTAINER Patrick Butler Monterde] <patrick.butlermonterde@emc.com>

# Get the project
RUN apt-get update &&\
  apt-get install -y git-core

# Get the project
RUN cd /opt &&\
  git clone https://github.com/emccode/mars-challenge

# Set working directory
WORKDIR /opt/mars-challenge/gateway

# Run Sensors Gateway
CMD python gateway-python.py

