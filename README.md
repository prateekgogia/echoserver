# echoserver

Simple gRPC based echoserver, which returns a string containing whatever the \
client sends in the request.

## Updating server protobuf API

Update the server API in api/api.proto file, this will generate the server and \
client gRPC libraries

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

Note: Make sure $PATH on the host contains $GOPATH/bin, else do \
`export PATH=$PATH:$GOPATH/bin/`
