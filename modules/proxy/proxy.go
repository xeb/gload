package proxy

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

var asock *zmq.Socket
var bsock *zmq.Socket

var listening bool = true

func Bind(agentadd string, bossadd string) {
	asock, _ = zmq.NewSocket(zmq.PUB)
	asock.Bind(agentadd)

	bsock, _ := zmq.NewSocket(zmq.REP)
	bsock.Bind(bossadd)

	for listening {
		// Wait for a command from BOSS
		fmt.Printf("[PROXY-boss] Receiving...\n")
		msg, _ := bsock.Recv(0)
		fmt.Printf("[PROXY-boss] Received: %s\n", msg)

		// Send to Agents
		_, _ = asock.Send(msg, 0)
		fmt.Printf("[PROXY-agent] Dispatched: %s\n", msg)

		// Inform the Boss, TODO: get response times from another channel
		_, _ = bsock.Send("--pretend these are results--", 0)
		fmt.Println("[PROXY-boss] Sent Message")

	}
}

func Unbind() {
	listening = false
}
