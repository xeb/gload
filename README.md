[![Build Status](https://travis-ci.org/xeb/gload.svg?branch=master)](https://travis-ci.org/xeb/gload)

WORK IN PROGRESS

# gload
Golang-based Load Testing.  

# Three Binaries
* ***Boss*** The REPL client for sending commands to remote agents via the coordinator
* ***Broker*** The broker that should be publically accessible to all agents.  Responsible for cluster management.
* ***Agent***  The remote process that connects to the broker, receives commands, executes them & returns results.  All stdin, stdout, stderr.

