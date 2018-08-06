[![Build Status](https://travis-ci.org/raphi011/scores.svg?branch=master)](https://travis-ci.org/raphi011/scores)

# scores

## Motivation

Collect data about our hobby volleyball matches to get to know more about who's my best partner / what was my best day / etc.

Support for other sports is planned but as of yet only picking teams of 2 is possible.

## Components

### Backend

Collect the data and present it as a REST api.

### Web Frontend

Create new / show past matches and statistics to authenticated users

### Teleram BOT

Get your statistics directly in your Telegram (Group) Chat!

## Build locally

1. Install [Docker](https://docs.docker.com/install/) and [Docker Compose](https://docs.docker.com/compose/install/).

1. Start Docker service.

1. Run `docker-compose -f docker-compose.yml -f docker-compose.dev.yml up
`.

1. Go to `http://localhost`


### Missing:

* How to seed database
* Run without attaching delve debugger