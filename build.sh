#!/bin/sh

cd src
mkdir -p time
mkdir -p httpio

# We probably dont need to go through the process of generating .pb.go's each time, but it seems worth it -- assuming it is possible
if hash protoc 2> /dev/null; then
	echo Generating proto models
	protoc -I ./protos ./protos/time.proto --go_out=plugins=grpc:time
	protoc -I ./protos ./protos/httpio.proto --go_out=plugins=grpc:httpio
fi
go build -o=../bin/gloada agent/main.go