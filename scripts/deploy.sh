#!/bin/bash

docker-compose -f docker-compose.prod.yml build
docker-compose -f docker-compose.prod.yml down
docker-compose -f docker-compose.prod.yml up -d --force-recreate --remove-orphans 