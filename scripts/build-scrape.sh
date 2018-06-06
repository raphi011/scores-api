#!/bin/bash

PLATFORM=$1

sudo docker rm -f scores-scrape-${PLATFORM}
sudo docker build -t raphi011/scores-scrape-${PLATFORM} ./cmd/terminal --build-arg app_env=production

