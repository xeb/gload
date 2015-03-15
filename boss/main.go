package main

import (
	"fmt"
        zmq "github.com/pebbe/zmq4"
)

const defbossadd string = "tcp://localhost:11000"

func main() {
	fmt.Printf("[BOSS] Connecting to Proxy at '%s'\n", defbossadd)
	Connect(defbossadd)
	
	SendCommand("date")
	SendCommand("ls -la")
	//SendCommand("gload PING")
	SendCommand("curl http://www.google.com")
}

var proxysoc *zmq.Socket

func Connect(address string) {
        proxysoc, _ = zmq.NewSocket(zmq.REQ)
        proxysoc.Connect(address)
}

func Close() {
        proxysoc.Close()
}

func SendCommand(command string) {
        fmt.Printf("[BOSS] Sending %s.\n", command)
        _, _ = proxysoc.Send(command, 0)

        fmt.Printf("[BOSS] Sent %s.  Waiting for response\n", command)
        r, _ := proxysoc.Recv(0)
        fmt.Printf("[BOSS] Received %s\n", r)
}
