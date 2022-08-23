# Simple spa-server

[![Docker Build](https://img.shields.io/docker/cloud/build/donatowolfisberg/spa-server)](https://hub.docker.com/r/donatowolfisberg/spa-server)
[![Image Version](https://img.shields.io/docker/v/donatowolfisberg/spa-server?sort=semver)](https://hub.docker.com/r/donatowolfisberg/spa-server)
[![Image Size](https://img.shields.io/docker/image-size/donatowolfisberg/spa-server?sort=date)](https://hub.docker.com/r/donatowolfisberg/spa-server)
[![Docker Pulls](https://img.shields.io/docker/pulls/donatowolfisberg/spa-server)](https://hub.docker.com/r/donatowolfisberg/spa-server)

We have all been there. You create a small backend with a single-page-application frontend and want to deploy it. Then you
create a nginx dockerfile and copy the nginx config you looked up 2 years ago from one of your other projects. This
without really knowing much about it and besides having to run full-fledged nginx server running for really no good
reason.

To solve this I created a small simple web server that does html5 routing. It does not come with the 99 features that
you don't use. It is written in golang and utilizes the embed feature so that there is just a single binary in your
final stage docker image.

It will send you the file if it finds it in the else it will send you the index.html file. This behavior is useful if you want to do client-side routing.

For example the empty image size of nginx is ~130MB in contrast the final build of the example in this repo is just
6.18MB.

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

## Options
The following options can be configured through environment variables.

| Env Name              | Default  | 
| --------------------- | -------- | 
| PORT                  | 8080     | 
| ADDRESS               | 0.0.0.0  | 
| READ_TIMEOUT_SECONDS  | 5        | 
| WRITE_TIMEOUT_SECONDS | 10       |
| IDLE_TIMEOUT_SECONDS  | 120      |
| BASE_HREF             | /        |
| CONFIG_JSON           | {}       |

* `BASE_HREF` is used to replace the `href` content in the `index.html`'s string `<base href="/"`, where the original string must match exactly the one mentioned here
* `CONFIG_JSON` must be json object that will be provided as the response for the request path `/config.json`

The `Cache-Control` is a one minute validity for `/index.html` and `/config.json`, and immutable for the rest of the responses.
## Build local

```shell
docker build . -t spa-server 
```
