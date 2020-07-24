FROM golang:1.14.6-alpine3.12 

WORKDIR /first-sample

ENV GO111MODULE on

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache protoc \
  && rm -rf /var/cache/apk/* \
  && go get -u google.golang.org/grpc \
  && go get -u github.com/golang/protobuf/protoc-gen-go \
  && go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

COPY . .
