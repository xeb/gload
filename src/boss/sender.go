package boss

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

var proxysoc *zmq.Socket

func Connect(address string) {
	proxysoc, _ = zmq.NewSocket(zmq.REQ)
	// defer proxysoc.Close()

	proxysoc.Connect(address)
}

func SendCommand(command string) {
	_, _ = proxysoc.Send(command, 0)

	fmt.Printf("[BOSS] Sent %s.  Waiting for response\n", command)
	r, _ := proxysoc.Recv(0)
	fmt.Printf("[BOSS] Received %s\n", r)
}
