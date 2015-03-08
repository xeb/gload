#!/bin/sh
./build.sh
bin/proxy &
bin/agent &
bin/boss
killall proxy
killall agent