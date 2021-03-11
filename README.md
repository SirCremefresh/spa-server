# Simple spa-server

```donatowolfisberg/spa-server```
[![Docker Build](https://img.shields.io/docker/cloud/build/donatowolfisberg/spa-server)](https://hub.docker.com/r/donatowolfisberg/spa-server)
[![Image Version](https://img.shields.io/docker/v/donatowolfisberg/spa-server?sort=semver)](https://hub.docker.com/r/donatowolfisberg/spa-server)
[![Image Size](https://img.shields.io/docker/image-size/donatowolfisberg/spa-server?sort=date)](https://hub.docker.com/r/donatowolfisberg/spa-server)
[![Docker Pulls](https://img.shields.io/docker/pulls/donatowolfisberg/spa-server)](https://hub.docker.com/r/donatowolfisberg/spa-server)

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

Create a "Dockerfile" where you start with the image "donatowolfisberg/spa-server". There you copy your application to
the public folder. After you run the "build.sh" script. This bundles your frontend with the server into a single binary.
In the next build-step start from "scratch" and copy your binary into the root path. At last, you need to add
CMD ["server"] to get it started.

```dockerfile
FROM donatowolfisberg/spa-server as builder

COPY public public
RUN ./build.sh

FROM scratch

COPY --from=builder /app/server /server
CMD ["/server"]
```

## Build local

```shell
docker build . -t spa-server 
```
