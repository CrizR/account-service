FROM golang:1.9.1 AS build-env

# TODO:
#   * Set ENV (enviornment variable) for Firebase config


ADD . /go/src/github.com/ecclesia-dev/account-service/
RUN cd /go/src/github.com/ecclesia-dev/account-service/app && go get && go build -o account-service

WORKDIR /go/src/github.com/ecclesia-dev/account-service/app

ENTRYPOINT ./account-service | tee ../logs/account-service.log