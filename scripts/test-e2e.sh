#!/bin/bash

# docker rm db
# docker rm web
docker rm -f test-e2e

docker pull raphi011/scores-web
docker pull raphi011/scores-test-e2e
docker pull postgres


# docker run --network=host --name db -e POSTGRES_PASSWORD=test -d postgres
# docker run --network=host --name web -e BACKEND_URL=http://localhost:8080 -d raphi011/scores-web

./api -provider postgres -connection "postgres://postgres:test@localhost?sslmode=disable" &

api_pid=$(echo $!)

docker run --network=host --name test-e2e -e CYPRESS_BASE_URL=http://localhost:3000 raphi011/scores-test-e2e

return_value=$?

kill -9 "$api_pid"

docker rm -f db
docker rm -f web

exit "$return_value"
