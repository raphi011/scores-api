#!/bin/bash

PLATFORM=$1
URL=$2

sudo docker rm -f scores-frontend-${PLATFORM}
sudo docker build -t raphi011/scores-frontend-${PLATFORM} ../web --build-arg backend_url=${URL}
