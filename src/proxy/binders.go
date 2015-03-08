package proxy

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

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
