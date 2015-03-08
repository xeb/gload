package agent

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func Subscribe(address string) {
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
