package main

import (
        zmq "github.com/pebbe/zmq4"
	"fmt"
)

const defagentadd string = "tcp://localhost:20000"

func main() {
	fmt.Printf("[AGENT] Connecting to Proxy at '%s'\n", defagentadd)
	Subscribe(defagentadd)
}

var proxysoc *zmq.Socket
var subscribed bool = true

func Subscribe(address string) {
        proxysoc, _ := zmq.NewSocket(zmq.SUB)
        defer proxysoc.Close()

        proxysoc.Connect(address)
        proxysoc.SetSubscribe("")

        fmt.Println("[AGENT] Connected")

        for subscribed {

                fmt.Println("[AGENT] Receiving")
                r, _ := proxysoc.Recv(0)
                fmt.Printf("[AGENT] Received %s\n", r)
                // _, _ = proxysoc.Send("STDOUT", 0)

                // fmt.Println("[AGENT] Sent PING.  Waiting for response")
        }

        fmt.Println("[AGENT] Unsubscribing")
}

func Close() {
        proxysoc.Close()
        subscribed = false
}
