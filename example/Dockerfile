FROM donatowolfisberg/spa-server-builder as builder

COPY public public
RUN ./build.sh

FROM donatowolfisberg/spa-server-runner

COPY --from=builder /app/server /server
