#!/bin/bash

go get ./...
go test -v ./... -tags repository
