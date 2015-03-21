package main

import (
	"fmt"
	"time"
	"github.com/xeb/gload/modules/proxy"
	"github.com/xeb/gload/modules/agent"
	"github.com/xeb/gload/modules/boss"
)

var (
	//agentaddr string = "tcp://*:1111"
	agentaddr string = "ipc:///tmp/agent"
	//bossaddr string = "tcp://*:2222"
	bossaddr string = "ipc:///tmp/boss"
)

func main() {
	fmt.Println("Starting Proxy...")

	go func() { 
		proxy.Bind(agentaddr, bossaddr)
	}()

	time.Sleep(time.Second * 2)

	fmt.Println("Starting Agent...")
	go func() {
		agent.Subscribe(agentaddr)
	}()

	time.Sleep(time.Second * 2)

	boss.Connect(bossaddr)

	boss.SendCommand("curl http://www.google.com")

	time.Sleep(time.Second * 2)
}
