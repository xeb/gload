#!/bin/sh

go get -a github.com/google/protobuf
cd $GOPATH/src/github.com/google/protobuf
./autogen.sh
./configure
make 
make check
make install
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
cd $GOPATH/src/github.com/xeb/gload

cd src/
mkdir -p time
mkdir -p httpio
protoc -I ./protos ./protos/time.proto --go_out=plugins=grpc:time
protoc -I ./protos ./protos/httpio.proto --go_out=plugins=grpc:httpio
go build -o=../bin/agent agent/main.go