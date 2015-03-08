package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"sync"
)

const defbossadd string = "tcp://*:10000"
const defagentadd string = "tcp://*:20000"

var servers chan int = make(chan int)
var asock *zmq.Socket

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

	servers <- 1
}

func BindAgent(address string) {
	fmt.Printf("[PROXY] Listening for Agents at '%s'\n", address)
	asock, _ = zmq.NewSocket(zmq.PUB)
	asock.Bind(address)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	BindAgent(defagentadd)

	fmt.Printf("[PROXY] Listening for Boss at '%s'\n", defbossadd)
	go BossBind(defbossadd)

	wg.Wait()
}
