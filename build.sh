#!/usr/bin/env sh

if [ ! -f "public/index.html" ]
then
    echo "Could not find index.html in public directory"
    exit 1
fi

CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server