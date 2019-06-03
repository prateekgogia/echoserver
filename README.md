# echoserver

Simple gRPC based echoserver, which returns a string containing whatever the client sends in the request.

## Build

This repo contains a Makefile and a Dockerfile, all the dependencies are added
to the Dockerfile.

To build gRPC server run the following command
```
    make build
```

This command will generate a Docker image called echoserver, this image contains
both a server and test client library to test the server.

## Usage

To run the echoserver in a container use the following commands. Name the Docker container
`server`, this way when the client container is started it can run the network namespace
as server container.

### Run in Server mode

`docker run -d --rm --name=server echoserver /bin/server`

### Run in test client mode

`docker run --rm --net=container:server echoserver /bin/client -msg=hi`

## Testing

To run the complete test suite including unit tests and integrations tests

```
make test
```

## Example

### Server mode

```
    root@ubuntu:/# docker run -ti --rm --name=server echoserver /bin/server
    2019/06/03 04:45:40 Receive message hi
```

### test client mode

```
    root@ubuntu:/# docker run --rm --net=container:server echoserver /bin/client -msg=hi
    2019/06/03 04:45:40 Response from server: hi
```

## Config flags

### Server

```
root@ubuntu:/# docker run -ti --rm --name=server echoserver /bin/server -help
Usage of /bin/server:
  -port int
        gRPC server port (default 8080)
```

### test client

```
root@ubuntu:/# docker run --rm --net=container:server echoserver /bin/client -help
Usage of /bin/client:
  -msg string
        Message to be sent to the server (default "hello")
  -port int
        gRPC server port (default 8080)
```

## Updating server protobuf API

Update the server API in api/api.proto file, this will generate the server and client gRPC libraries

### Install protoc

```
apt-get install unzip && \
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip && \
unzip protoc-3.8.0-linux-x86_64.zip -d /tmp/ && mv /tmp/bin/protoc /usr/bin/
```

### Install protoc-gen-go plugin

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

Note: Make sure $PATH on the host contains $GOPATH/bin, else do `export PATH=$PATH:$GOPATH/bin/`
