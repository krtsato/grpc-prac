FROM golang:1.14.6-alpine3.12 


ENV PROTOBUF_VERSION 3.12.3
ENV PROTOBUF_URL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/protobuf-cpp-$PROTOBUF_VERSION.tar.gz

RUN set -oux pipefail \
  && apk update \
  && apk add --no-cache bash vim curl build-base autoconf automake \
  && curl -L -o /first-sample/protobuf.tar.gz $PROTOBUF_URL \
  && tar -xvzf protobuf.tar.gz \
  && rm -rf /var/cache/apk/*

WORKDIR /protobuf-$PROTOBUF_VERSION

RUN set -oux pipefail \
  && ./autogen.sh \
  && ./configure --prefix=/usr/local/protobuf/3_9_0 \
  && make -j 4 \
  && make install \
  && go get -u google.golang.org/grpc \
  && go get -u github.com/golang/protobuf/protoc-gen-go \
  && go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

WORKDIR /first-sample

COPY . /first-sample
