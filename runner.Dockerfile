FROM alpine:3
RUN apk update && \
    apk add --no-cache ca-certificates mailcap

CMD ["/server"]