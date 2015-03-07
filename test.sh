#!/bin/sh
killall proxy
killall boss

go build -o=./bin/proxy src/proxy/main.go
go build -o=./bin/boss src/boss/main.go

bin/proxy &
bin/boss