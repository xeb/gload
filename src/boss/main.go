package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

const defbossadd string = "tcp://localhost:10000"

func ProxySend(address string) {
	proxysoc, _ := zmq.NewSocket(zmq.REQ)
	defer proxysoc.Close()

	proxysoc.Connect(address)

	_, _ = proxysoc.Send("HELLO!", 0)

	fmt.Println("[BOSS] Sent Hello.  Waiting for response")
	r, _ := proxysoc.Recv(0)
	fmt.Printf("[BOSS] Received %s\n", r)
}

func main() {
	fmt.Printf("[BOSS] Connecting to Proxy at '%s'\n", defbossadd)
	ProxySend(defbossadd)
}
