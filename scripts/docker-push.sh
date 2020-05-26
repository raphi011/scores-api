#!/bin/bash

cwd=$(dirname "$0")
version=$($cwd/version.sh)

echo $version

docker login \
    --username="${DOCKER_USERNAME}" \
    --password="${DOCKER_PASSWORD}"

docker build \
    -t raphi011/scores-api:$version \
    -t raphi011/scores-api:latest \
    .

docker push raphi011/scores-api:$version
docker push raphi011/scores-api:latest
