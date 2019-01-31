[![Build Status](https://travis-ci.org/raphi011/scores.svg?branch=development)](https://travis-ci.org/raphi011/scores)

# scores

## Motivation

This is mainly a hobby project of mine to improve my skills creating a modern javascript SPA + GO backend with all the newest bells and whistles.

### Incomplete list of features / tools:

- Serverside Rending (courtesy of the awesome [nextjs](https://nextjs.org/) framework).
- Automated builds / tests / deployments with [travis-ci](https://travis-ci.org).
- Debugging of Front + Backend (via VS-Code).
- Automated database migrations.
- Logging to the ELK stack.
- Webscraping of http://www.volleynet.at/beach.
- Plugable persistance architecture, currently it's possible to store the data with Postgresql, MySQL and Sqlite3. No-SQL db's could also be supported.
- Well tested against all supported data stores
  - Integration tests
  - Unit tests
  - Browser E2E tests (TODO)

## Components

### Backend

Srape the data from the official Austrian beach volleynet homepage and present it as a REST api.

### Web Frontend

Signup/out of Tournaments, browse and filter through tournaments. Get notifications, ...

### Teleram BOT

TODO

## Build locally

Development is done on Linux with VS-Code.

To get up and running follow these steps:

1. Install Node / Go 1.11+ / Docker + Docker Compose
1. Run npm install in ./web-client
1. Run docker-compose up
1. Start Frontend / Backend in VS-Code
1. Create test admin account by navigating to `localhost/api/debug/new-admin`
1. Open `localhost` in your browser of choice and login

## FAQ

- _Do you plan to earn money with this project?_  
  Nope I'm doing this purely for educational and practical reasons :).

- _Do you store volleynet passwords when logging in to tournaments?_  
  No, the only time you need to enter the volleynet password is when signing up/out of tournaments - since there is no offical API I'm unable to authenticate without having the user provide his/her cleartext password. But rest assured (or look at the sourcecode ;) ) that I will not do anything evil with it.

## Contributing

To contribute just open an issue and tell me how you would like to help!
