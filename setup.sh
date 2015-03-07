#!/bin/sh


# install protoc if necessary

echo Getting protobuf
go get -a github.com/google/protobuf
pwd
cd $GOPATH/src/github.com/google/protobuf
./autogen.sh
./configure --prefix=/usr
make 
make check
sudo make install
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
