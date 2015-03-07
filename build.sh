#!/bin/sh
go build -o=./bin/agent src/agent/main.go
go build -o=./bin/boss src/boss/main.go
go build -o=./bin/proxy src/proxy/main.go
