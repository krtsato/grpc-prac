FROM golang:1.14.6-alpine3.12 

WORKDIR /first-sample

RUN set -ox pipefail \
  && apk update \
  && apk add --no-cache protobuf \
  && rm -rf /var/cache/apk/*

RUN set -ox pipefail \
  && export GO111MODULE=on \
  && go get -u google.golang.org/grpc \
  && go get -u github.com/golang/protobuf/protoc-gen-go \
  && go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

COPY . .
