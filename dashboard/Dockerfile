# Base image
FROM ubuntu

# Maintainer info
MAINTAINER Adrian Moreno <adrian.moreno@emc.com>

# Get the project
RUN apt-get update &&\
  apt-get install -y git-core ruby ruby-dev nodejs-legacy npm nginx

RUN gem install compass

RUN npm install bower grunt-cli -g &&\
  npm install grunt

RUN cd /opt &&\
  git clone https://github.com/emccode/mars-challenge &&\
  cd /opt/mars-challenge/dashboard/ &&\
  npm install &&\
  bower install --allow-root --force-latest --config.interactive=false &&\
  grunt build &&\
  cp -r dist /app

# setup all the config files
RUN rm /etc/nginx/sites-enabled/default &&\
  ln -s /opt/mars-challenge/dashboard/dashboard_nginx.conf /etc/nginx/sites-enabled/dashboard

# Set the default directory where CMD will execute
WORKDIR /opt/mars-challenge/dashboard

EXPOSE 80

CMD sh start.sh
