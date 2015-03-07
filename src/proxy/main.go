package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

const defbossadd string = "tcp://*:10000"

func BossBind(address string) {
	bosssock, _ := zmq.NewSocket(zmq.REP)
	defer bosssock.Close()
	bosssock.Bind(address)

	msg, _ := bosssock.Recv(0)
	fmt.Printf("[PROXY] Received: %s\n", msg)
	v, e := bosssock.Send("TEST", 0)
	fmt.Printf("[PROXY] Sent Pong %s\n", v)
	if e != nil {
		panic(e)
	}

}

func main() {
	fmt.Printf("[PROXY] Listening for Boss at '%s'\n", defbossadd)
	BossBind(defbossadd)
}
