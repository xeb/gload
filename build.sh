#!/bin/sh


# install protoc if necessary
hash protoc &> /dev/null
if [ $? -eq 1 ]; then
	echo Getting protobuf
	go get -a github.com/google/protobuf
	pwd
	cd /home/travis/gopath/src/github.com/google/protobuf
	./autogen.sh
	./configure --prefix=/usr
	make 
	make check
	make install
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	cd /home/travis/gopath/src/github.com/xeb/gload
fi

cd src/
mkdir -p time
mkdir -p httpio

# We probably dont need to go through the headache of generating .pb.go's each time, but it seems worth it -- assuming it is possible

protoc -I ./protos ./protos/time.proto --go_out=plugins=grpc:time
protoc -I ./protos ./protos/httpio.proto --go_out=plugins=grpc:httpio
go build -o=../bin/gloada agent/main.go