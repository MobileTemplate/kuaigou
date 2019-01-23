#!/usr/bin/env bash

set -e

source ./build.sh

set -e

echo ''
echo 'auto action ...'
echo 'push image: ' $DOCKER_IMAGE
docker push $DOCKER_IMAGE
echo 'push rmi image: ' $DOCKER_IMAGE
docker rmi $DOCKER_IMAGE
