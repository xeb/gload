package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

const defbossadd string = "tcp://*:11000"
const defagentadd string = "tcp://*:20000"

var asock *zmq.Socket
var done chan bool = make(chan bool)

func BindBoss(address string) {
	bsock, _ := zmq.NewSocket(zmq.REP)
	defer bsock.Close()
	bsock.Bind(address)

	for {
		msg, _ := bsock.Recv(0)
		fmt.Printf("[PROXY-boss] Received: %s\n", msg)

		// Send to Agent Socket
		_, _ = asock.Send(msg, 0)
		fmt.Printf("[PROXY-agent] Dispatched: %s\n", msg)

		_, _ = bsock.Send("PONG", 0)
		fmt.Println("[PROXY-boss] Sent PONG")
	}

	done <- true
}

func BindAgent(address string) {
	asock, _ = zmq.NewSocket(zmq.PUB)
	asock.Bind(address)
}

func main() {

	wg.Add(1)

	fmt.Printf("[PROXY] Listening for Agents at '%s'\n", defagentadd)
	BindAgent(defagentadd)

	fmt.Printf("[PROXY] Listening for Boss at '%s'\n", defbossadd)
	BindBoss(defbossadd)

	<-done
}
