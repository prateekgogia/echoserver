FROM golang:1.12.5 as builder

LABEL maintainer="prateekgogia@hotmail.com"

RUN apt-get update -y && apt-get install -y ca-certificates && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    apt-get install unzip && \
    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip && \
    unzip protoc-3.8.0-linux-x86_64.zip -d /tmp/ && mv /tmp/bin/protoc /usr/bin/

ENV GOPATH=/workspace/golang/

COPY . /workspace/golang/src/github.com/prateekgogia/echoserver/

WORKDIR /workspace/golang/src/github.com/prateekgogia/echoserver

RUN protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/api.proto && \
    go build -v -o bin/server github.com/prateekgogia/echoserver/cmd/server && \
    go build -v -o bin/client github.com/prateekgogia/echoserver/cmd/client && \
    mv bin/server /usr/bin/ && \
    mv bin/client /usr/bin/ && mkdir cert && scripts/generate_certs.sh && \
    mv cert /usr/local/include/

FROM ubuntu:latest

COPY --from=builder /usr/bin/server /bin/server
COPY --from=builder /usr/bin/client /bin/client
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /usr/local/include/cert /usr/local/include/cert