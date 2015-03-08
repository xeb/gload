package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

const defagentadd string = "tcp://localhost:20000"

func ProxySend(address string) {
	proxysoc, _ := zmq.NewSocket(zmq.SUB)
	defer proxysoc.Close()

	proxysoc.Connect(address)
	proxysoc.SetSubscribe("")

	fmt.Println("[AGENT] Connected")

	for {
		r, _ := proxysoc.Recv(0)
		fmt.Printf("[AGENT] Received %s\n", r)
		// _, _ = proxysoc.Send("STDOUT", 0)

		// fmt.Println("[AGENT] Sent PING.  Waiting for response")
	}
}

func main() {
	fmt.Printf("[AGENT] Connecting to Proxy at '%s'\n", defagentadd)
	ProxySend(defagentadd)
}
