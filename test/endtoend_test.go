package gload

import (
	agent "github.com/xeb/gload/src/agent"
	boss "github.com/xeb/gload/src/boss"
	proxy "github.com/xeb/gload/src/proxy"
	"testing"
	"time"
)

var agentadd string = "tcp://*:50111"
var bossadd string = "tcp://*:50112"

func TestEndToEnd(t *testing.T) {

	t.Log("Proxy is Binding Agent")
	go proxy.Bind(agentadd, bossadd)

	t.Log("Agent is Subscribing")
	go agent.Subscribe(agentadd)
	defer agent.Close()

	t.Log("Boss is Connecting")
	boss.Connect(bossadd)
	defer boss.Close()

	time.Sleep(time.Second * 5)

	t.Log("Boss is Sending Command")
	boss.SendCommand("TEST")
}
