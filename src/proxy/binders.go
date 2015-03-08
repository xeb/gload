package proxy

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

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
}

func BindAgent(address string) {
	asock, _ = zmq.NewSocket(zmq.PUB)
	asock.Bind(address)
}
