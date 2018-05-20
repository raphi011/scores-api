#!/bin/bash

PLATFORM=$1
URL=$2

sudo docker rm -f scores-backend-${PLATFORM}
sudo docker build -t raphi011/scores-backend-${PLATFORM} .. --build-arg app_env=production --build-arg backend_url=${URL}
