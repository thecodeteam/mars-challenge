#!/bin/sh

if [ "$WS_ENDPOINT" ]; then
    FILE=$(ls -d /app/scripts/s*)
    sed -i 's,localhost:8080/ws,'$WS_ENDPOINT',g' $FILE
fi

/usr/sbin/nginx -g "daemon off;"
