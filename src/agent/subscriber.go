package agent

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

var proxysoc *zmq.Socket
var subscribed bool = true

func Subscribe(address string) {
	proxysoc, _ := zmq.NewSocket(zmq.SUB)
	defer proxysoc.Close()

	proxysoc.Connect(address)
	proxysoc.SetSubscribe("")

	fmt.Println("[AGENT] Connected")

	for subscribed {
		r, _ := proxysoc.Recv(0)
		fmt.Printf("[AGENT] Received %s\n", r)
		// _, _ = proxysoc.Send("STDOUT", 0)

		// fmt.Println("[AGENT] Sent PING.  Waiting for response")
	}
}

func Close() {
	proxysoc.Close()
	subscribed = false
}
