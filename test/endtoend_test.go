package gload

import (
	agent "github.com/xeb/gload/src/agent"
	boss "github.com/xeb/gload/src/boss"
	proxy "github.com/xeb/gload/src/proxy"
	"testing"
)

var agentadd string = "tcp://*:50111"
var bossadd string = "tcp://*:50112"

func TestEndToEnd(t *testing.T) {

	t.Log("Proxy is Binding Agent")
	proxy.BindAgent(agentadd)

	t.Log("Proxy is Binding Boss")
	proxy.BindBoss(bossadd)

	t.Log("Agent is Subscribing")
	agent.Subscribe(agentadd)

	t.Log("Boss is Connecting")
	boss.Connect(bossadd)

	t.Log("Boss is Sending Command")
	boss.SendCommand("TEST")

	t.Log("Closing Agent")
	agent.Close()
}
