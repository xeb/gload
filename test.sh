#!/bin/sh
set -e
go test -v ./...
./build.sh
bin/proxy &
bin/agent &
bin/agent &
bin/agent &
sleep 1
bin/boss
sleep 2
killall proxy
killall agent
echo Success