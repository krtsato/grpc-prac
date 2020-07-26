FROM golang:1.14.6-buster

WORKDIR /reverse-proxy-sample

ENV GO111MODULE on
ENV PROTOBUF_VERSION 3.12.3
ENV PROTOBUF_URL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip

RUN set -ex \
  && apt update \
  && apt install -y --no-install-recommends unzip \
  && apt clean \
  && curl -LO ${PROTOBUF_URL} \
  && unzip protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -d ./protoc \
  && mv ./protoc/bin/* /usr/local/bin/ \
  && mv ./protoc/include/* /usr/local/include/ \
  && rm -rf ./{protoc-${PROTOBUF_VERSION}-linux-x86_64.zip,protoc} \
  && go get -u google.golang.org/grpc \
  && go get -u github.com/golang/protobuf/protoc-gen-go \
  && go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
  && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

COPY . .
