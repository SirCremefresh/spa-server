#!/usr/bin/env sh

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server