# Base image
FROM ubuntu

# Maintainer info
MAINTAINER Adrian Moreno <adrian.moreno@emc.com>

#  Update the sources list
RUN apt-get update

# Get the project
RUN apt-get update &&\
  apt-get install -y git-core nodejs-legacy npm nginx

RUN npm install bower -g

RUN cd /opt &&\
  git clone https://github.com/emccode/mars-challenge &&\
  cd /opt/mars-challenge/dashboard/ &&\
  bower install --allow-root --force-latest --config.interactive=false

# setup all the config files
RUN rm /etc/nginx/sites-enabled/default &&\
  ln -s /opt/mars-challenge/dashboard/dashboard_nginx.conf /etc/nginx/sites-enabled/dashboard

RUN service nginx restart

EXPOSE 80

CMD /bin/bash
