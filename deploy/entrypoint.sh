#!/bin/sh
# Replace env vars from nginx template file to it's prod conf
set -e
VARS_TO_SUBSTITUTE='$NGINX_PORT $AUTH_HOST $AUTH_PORT $MOVIE_HOST $MOVIE_PORT'
envsubst "$VARS_TO_SUBSTITUTE" < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf
exec "$@"