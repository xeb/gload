#!/bin/sh
go build -o=./bin/agent src/execs/agent.go
go build -o=./bin/boss src/execs/boss.go
go build -o=./bin/proxy src/proxy/main.go
