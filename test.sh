#!/bin/sh
./build.sh
bin/proxy &
bin/boss
for i in {1..5}; do bin/agent; done
killall proxy