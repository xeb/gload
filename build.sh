#!/bin/sh
rm -rf $GOPATH/pkg # pre-existing pkgs causing me some grief
mkdir -p bin
go build -o=./bin/agent agent/main.go
go build -o=./bin/boss boss/main.go
go build -o=./bin/proxy proxy/main.go
