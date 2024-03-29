FROM golang:1.21 as builder
LABEL maintainer="donato@wolfisberg.dev"
WORKDIR /app

COPY ["main.go", "go.mod", "build.sh", "./"]
RUN chmod +x build.sh
