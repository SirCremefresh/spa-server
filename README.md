# Simple spa-server

```donatowolfisberg/spa-server-builder```
[![Docker Build](https://img.shields.io/docker/cloud/build/donatowolfisberg/spa-server-builder)](https://hub.docker.com/r/donatowolfisberg/spa-server-builder)
[![Image Version](https://img.shields.io/docker/v/donatowolfisberg/spa-server-builder?sort=semver)](https://hub.docker.com/r/donatowolfisberg/spa-server-builder)
[![Image Size](https://img.shields.io/docker/image-size/donatowolfisberg/spa-server-builder?sort=date)](https://hub.docker.com/r/donatowolfisberg/spa-server-builder)
[![Docker Pulls](https://img.shields.io/docker/pulls/donatowolfisberg/spa-server-builder)](https://hub.docker.com/r/donatowolfisberg/spa-server-builder)  
```donatowolfisberg/spa-server-runneer```
[![Docker Build](https://img.shields.io/docker/cloud/build/donatowolfisberg/spa-server-runner)](https://hub.docker.com/r/donatowolfisberg/spa-server-runner)
[![Image Version](https://img.shields.io/docker/v/donatowolfisberg/spa-server-runner?sort=semver)](https://hub.docker.com/r/donatowolfisberg/spa-server-runner)
[![Image Size](https://img.shields.io/docker/image-size/donatowolfisberg/spa-server-runner?sort=date)](https://hub.docker.com/r/donatowolfisberg/spa-server-runner)
[![Docker Pulls](https://img.shields.io/docker/pulls/donatowolfisberg/spa-server-runner)](https://hub.docker.com/r/donatowolfisberg/spa-server-runner)

Who doesn't know it? You create a small backend with a single-page-application frontend and want to deploy it. Then you
create a nginx dockerfile and copy the nginx config you looked up 2 years ago from one of your other projects. This
without really knowing much about it and besides having to run full-fledged nginx server running for really no good
reason.

To solve this I created a small simple web server that does html5 routing. It does not come with the 99 features that
you don't use. It is written in golang and utilizes the embed feature so that there is just a single binary in your
final stage docker image.

## Example

There is an example setup in the [./example](https://github.com/SirCremefresh/spa-server/example) folder in
this [repository](https://github.com/SirCremefresh/spa-server).

## Usage

Create a "Dockerfile" where you start with the image "donatowolfisberg/spa-server-builder". There you copy your
application to the public folder. After you run the "build.sh" script. This bundles your frontend with the server into a
single binary. In the next build-step start from "donatowolfisberg/spa-server-runner" and copy your binary into the root
path.

```dockerfile
FROM donatowolfisberg/spa-server-builder as builder

COPY public public
RUN ./build.sh

FROM donatowolfisberg/spa-server-runner

COPY --from=builder /app/server /server

```

## Build local

```shell
docker build --file builder.Dockerfile . -t spa-server-builder 
docker build --file runner.Dockerfile . -t spa-server-runner   
```
