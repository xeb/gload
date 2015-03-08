package main

import (
	"fmt"
        zmq "github.com/pebbe/zmq4"
)

const defbossadd string = "tcp://*:11000"
const defagentadd string = "tcp://*:20000"

func main() {

	fmt.Printf("[PROXY] Listening for Agents at '%s'\n", defagentadd)
	fmt.Printf("[PROXY] Listening for Boss at '%s'\n", defbossadd)
	Bind(defagentadd, defbossadd)
}

var asock *zmq.Socket
var bsock *zmq.Socket

func Bind(agentadd string, bossadd string) {
        asock, _ = zmq.NewSocket(zmq.PUB)
        asock.Bind(agentadd)

        bsock, _ := zmq.NewSocket(zmq.REP)
        bsock.Bind(bossadd)

        for {
                fmt.Printf("[PROXY-boss] Receiving...\n")
                msg, _ := bsock.Recv(0)
                fmt.Printf("[PROXY-boss] Received: %s\n", msg)

                // Send to Agent Socket
                _, _ = asock.Send(msg, 0)
                fmt.Printf("[PROXY-agent] Dispatched: %s\n", msg)

                _, _ = bsock.Send("PONG", 0)
                fmt.Println("[PROXY-boss] Sent PONG")

        }
}
