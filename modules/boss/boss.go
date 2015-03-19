package boss

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

var proxysoc *zmq.Socket

func Connect(address string) {
	proxysoc, _ = zmq.NewSocket(zmq.REQ)
	proxysoc.Connect(address)
}

func Close() {
	proxysoc.Close()
}

func SendCommand(command string) {
	fmt.Printf("[BOSS] Sending %s.\n", command)
	_, _ = proxysoc.Send(command, 0)

	fmt.Printf("[BOSS] Sent %s.  Waiting for response\n", command)
	r, _ := proxysoc.Recv(0)
	fmt.Printf("[BOSS] Received %s\n", r)
}
