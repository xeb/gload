package main

import (
	"fmt"
	"time"
	"github.com/xeb/gload/modules/proxy"
	"github.com/xeb/gload/modules/agent"
	"github.com/xeb/gload/modules/boss"
)

var (
	agentaddr string = "tcp://*:1111"
	bossaddr string = "tcp://*:2222"
)

func main() {
	fmt.Println("Starting Proxy...")

	go func() { 
		proxy.Bind(agentaddr, bossaddr)
	}()

	fmt.Println("Starting Agent...")
	go func() {
		agent.Subscribe(agentaddr)
	}()

	time.Sleep(time.Second)

	boss.Connect(bossaddr)

	boss.SendCommand("curl http://www.google.com")
	
}
