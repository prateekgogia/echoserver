#!/bin/bash
openssl genrsa -out cert/server.key 2048
openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650 -subj "/C=US/ST=California/L=SanFrancisco/O=gRPC EchoServer/CN=localhost"
openssl req -new -sha256 -key cert/server.key -out cert/server.csr -subj "/C=US/ST=California/L=SanFrancisco/O=gRPC EchoServer/CN=localhost"
openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650