# Base image
FROM ubuntu

# Maintainer info
MAINTAINER Patrick Butler Monterde <Patrick.butlermonterde@emc.com>

# Get the project
RUN apt-get update &&\
  apt-get install -y git-core python python-pip python-dev

RUN pip install websocket-client

RUN cd /opt &&\
  git clone https://github.com/emccode/mars-challenge &&\
  cd /opt/mars-challenge/gateway/python/

# Set the default directory where CMD will execute
WORKDIR /opt/mars-challenge/gateway/python/

CMD python gateway-python.py
