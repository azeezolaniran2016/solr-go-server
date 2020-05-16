FROM golang:1.13.1-alpine3.10

WORKDIR /github.com/azeezolaniran2016/solr-server

ENV CGO_ENABLED=0\
    GOOS=linux

RUN apk add --no-cache make git bash curl gcc musl-dev 

RUN go get github.com/oxequa/realize