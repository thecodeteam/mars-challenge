#!/bin/sh

if [ "$WS_ENDPOINT" ]; then
    sed -i 's,localhost:8080/ws,'$WS_ENDPOINT',g' /opt/mars-challenge/dashboard/app/scripts/controllers/main.js
fi

/usr/sbin/nginx -g "daemon off;"
