#!/bin/bash

cwd=$(dirname $0)

gpg --batch --passphrase $CERTS_PW -d $cwd/../deploy/docker-auth.tar.gz.gpg | tar xzf -

export DOCKER_CERT_PATH=$(pwd)
export DOCKER_TLS_VERIFY=1
export DOCKER_HOST=${DOCKER_DEPLOY_URL:-scores.network:2376}

docker pull raphi011/scores-api:latest

docker stop scores-api

docker rm scores-api

docker run \
    -d \
    -p 8080 \
    --network="scores" \
    --network-alias=["api"] \
    --name scores-api \
    --mount 'type=volume,src=scores-api,dst=/srv/scores' \
    raphi011/scores-api:latest
