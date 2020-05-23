#!/bin/bash

cwd=$(dirname "$0")
version=$($cwd/version.sh)

echo $version

docker login \
    --username="${DOCKER_USERNAME}" \
    --password="${DOCKER_PASSWORD}"

docker build \
    -t raphi011/scores-backend:$version \
    -t raphi011/scores-backend:latest \
    .

docker push raphi011/scores-backend:$version
docker push raphi011/scores-backend:latest
