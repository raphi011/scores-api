#!/bin/bash

docker-compose -f docker-compose.yml -f docker-compose.prod.yml build
docker-compose -f docker-compose.yml -f docker-compose.prod.yml down
docker-compose -f docker-compose.yml -f docker-compose.prod.yml -d --remove-orphans up