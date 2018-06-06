#!/bin/bash

PROG=all
CHECKOUT=master
PLATFORM=release
URL=https://scores.raphi011.com
PORT_FRONTEND=3000
PORT_BACKEND=8080
DATA_DIRECTORY=/home/raphi/scores-data

for i in "$@"
do
case $i in
    -c=*|--checkout=*)
    CHECKOUT="${i#*=}"
    shift
    ;;
    -b|--beta)
    PLATFORM=beta
    PORT_FRONTEND=3001
    PORT_BACKEND=8081
    URL=https://beta.raphi011.com
	DATA_DIRECTORY=/home/raphi/scores-data-beta
    shift
    ;;
    *)
    # unknown option
    ;;
esac
done

if [[ -n $1 ]]; then
    PROG=$1
fi


echo "Deploying ${PROG} on ${PLATFORM} at version ${CHECKOUT}"

cd /home/raphi/go/src/github.com/raphi011/scores
git checkout ${CHECKOUT}
git pull

case "$PROG" in 
	"frontend")
		echo "Building frontend"
		./scripts/build-frontend.sh ${PLATFORM} ${URL}
		;;
	"backend")
		echo "Building backend"
		./scripts/build-backend.sh ${PLATFORM} ${URL}
		;;
	"scrape")
		echo "Building scraper"
		./scripts/build-scrape.sh ${PLATFORM}
		;;
	"all")
		echo "Building frontend"
		./scripts/build-frontend.sh ${PLATFORM} ${URL} &
		echo "Building backend"
		./scripts/build-backend.sh ${PLATFORM} ${URL} &
		echo "Building scraper"
		./scripts/build-scrape.sh ${PLATFORM}
		;;
	*)
		echo "Invalid prog: ${PROG}"
		exit 1
		;;
esac

wait

case "$PROG" in 
	"frontend")
		sudo docker run -it -d --name scores-frontend-${PLATFORM} -p 127.0.0.1:${PORT_FRONTEND}:3000 raphi011/scores-frontend-${PLATFORM};;
	"backend")
		sudo docker run -it -d --name scores-backend-${PLATFORM} -v ${DATA_DIRECTORY}:/srv/scores -p 127.0.0.1:${PORT_BACKEND}:8080 raphi011/scores-backend-${PLATFORM};;
	"scrape")
		sudo docker run -it -d --name scores-scrape-${PLATFORM} --net=host raphi011/scores-scrape-${PLATFORM} -url http://localhost:${PORT_BACKEND};;
	*)
		sudo docker run -it -d --name scores-frontend-${PLATFORM} -p 127.0.0.1:${PORT_FRONTEND}:3000 raphi011/scores-frontend-${PLATFORM}
		sudo docker run -it -d --name scores-backend-${PLATFORM} -v ${DATA_DIRECTORY}:/srv/scores -p 127.0.0.1:${PORT_BACKEND}:8080 raphi011/scores-backend-${PLATFORM}
		sudo docker run -it -d --name scores-scrape-${PLATFORM} --net=host raphi011/scores-scrape-${PLATFORM} -url http://localhost:${PORT_BACKEND};;

		;;
esac

echo "Done"
