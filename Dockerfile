FROM golang:1.17-alpine as dev-env

WORKDIR /app

FROM dev-env as build-env
COPY . /app/

RUN CGO_ENABLED=0 go build -o /echoserver

FROM alpine:3.10 as runtime

COPY --from=build-env /echoserver /usr/local/bin/echoserver
RUN chmod +x /usr/local/bin/echoserver

ENTRYPOINT ["echoserver"]