#!/bin/sh

cd src/
mkdir -p time
mkdir -p httpio
protoc -I ./protos ./protos/time.proto --go_out=plugins=grpc:time
protoc -I ./protos ./protos/httpio.proto --go_out=plugins=grpc:httpio
go build -o=../bin/agent agent/main.go