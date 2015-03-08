#!/bin/sh
go build -o=./bin/agent src/agent.go
go build -o=./bin/boss src/boss.go
go build -o=./bin/proxy src/proxy/main.go
