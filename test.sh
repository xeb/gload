#!/bin/sh
./build.sh
bin/proxy &
bin/agent &
sleep 1
bin/boss
killall proxy
killall agent