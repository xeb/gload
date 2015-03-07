package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

const defagentadd string = "tcp://localhost:20000"

func ProxySend(address string) {
	proxysoc, _ := zmq.NewSocket(zmq.REQ)
	defer proxysoc.Close()

	proxysoc.Connect(address)

	_, _ = proxysoc.Send("PING", 0)

	fmt.Println("[AGENT] Sent PING.  Waiting for response")
	r, _ := proxysoc.Recv(0)
	fmt.Printf("[AGENT] Received %s\n", r)
}

func main() {
	fmt.Printf("[AGENT] Connecting to Proxy at '%s'\n", defagentadd)
	ProxySend(defagentadd)
}
