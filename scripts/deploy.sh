#!/bin/bash

cwd=$(dirname $0)

gpg --batch --passphrase $CERTS_PW -d $cwd/../deploy/docker-auth.tar.gz.gpg | tar xzf -

export DOCKER_CERT_PATH=$(pwd)
export DOCKER_TLS_VERIFY=1
export DOCKER_HOST=${DOCKER_DEPLOY_URL:-scores.network:2376}

docker pull raphi011/scores-backend:latest

docker stop scores-backend

docker rm scores-backend

docker run \
    -d \
    -p 8080 \
    --network="scores" \
    --network-alias=["backend"] \
    --name scores-backend \
    --mount 'type=volume,src=scores-backend,dst=/srv/scores' \
    raphi011/scores-backend:latest
