FROM donatowolfisberg/spa-server as builder

COPY public public
RUN ./build.sh

FROM scratch
COPY --from=builder /app/server /server
CMD ["/server"]