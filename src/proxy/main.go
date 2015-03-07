package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"sync"
)

const defbossadd string = "tcp://*:10000"
const defagentadd string = "tcp://*:20000"

var servers chan int = make(chan int)

func BossBind(address string) {
	bsock, _ := zmq.NewSocket(zmq.REP)
	defer bsock.Close()
	bsock.Bind(address)

	for {
		msg, _ := bsock.Recv(0)
		fmt.Printf("[PROXY-boss] Received: %s\n", msg)
		v, _ := bsock.Send("PONG", 0)
		fmt.Printf("[PROXY-boss] Sent Pong %s\n", v)
	}

	servers <- 1
}

func AgentBind(address string) {
	asock, _ := zmq.NewSocket(zmq.REP)
	defer asock.Close()
	asock.Bind(address)

	for {
		msg, _ := asock.Recv(0)
		fmt.Printf("[PROXY-agent] Received: %s\n", msg)
		v, _ := asock.Send("PONG", 0)
		fmt.Printf("[PROXY-agent] Sent Pong %s\n", v)
	}

	servers <- 2
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Printf("[PROXY] Listening for Boss at '%s'\n", defbossadd)
	go BossBind(defbossadd)

	fmt.Printf("[PROXY] Listening for Agents at '%s'\n", defagentadd)
	go AgentBind(defagentadd)

	wg.Wait()
}
