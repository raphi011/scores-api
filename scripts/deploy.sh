#!/bin/bash

VERSION=$(git describe --always --long) docker-compose -f docker-compose.prod.yml up -d --build