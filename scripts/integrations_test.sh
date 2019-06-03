#!/bin/bash
# set -x

MESSAGE="hola"
IMAGE_NAME=echoserver
CONTAINER_NAME=server

# Start server
docker run -d --rm --name=server ${IMAGE_NAME} /bin/server

## Check reply from server
OUTPUT=$(docker run -ti --rm --net=container:${CONTAINER_NAME} ${IMAGE_NAME} /bin/client -msg=${MESSAGE})
RESULT=$(echo $OUTPUT | awk '{ print $NF }' | tr -d '\r')
if [ $(echo $OUTPUT | awk '{ print $NF }' | tr -d '\r') != "$MESSAGE" ]; then
    exit 1
fi

## Close server
docker rm -f server

exit 0