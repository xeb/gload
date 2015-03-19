package agent

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os/exec"
	"strings"
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

		fmt.Println("[AGENT] Receiving")
		r, _ := proxysoc.Recv(0)
		fmt.Printf("[AGENT] Received %s\n", r)

		fmt.Printf("[AGENT] Executing shell command\n")

		stdout, e := Exec(r)

		if e != nil {
			fmt.Printf("[AGENT] Stderr was %s\n", e)
		} else {
			fmt.Printf("[AGENT] Stdout was %s\n", stdout)
		}
	}

	fmt.Println("[AGENT] Unsubscribing")
}

func Exec(s string) (stdout string, e error) {

	parts := strings.Split(s, " ")
	cmd := exec.Command(parts[0], strings.Join(parts[1:], " "))
	b, e := cmd.Output()
	stdout = string(b)
	return
}

func Close() {
	proxysoc.Close()
	subscribed = false
}
