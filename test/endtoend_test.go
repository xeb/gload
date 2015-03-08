package gload

import (
	agent "github.com/xeb/gload/src/agent"
	boss "github.com/xeb/gload/src/boss"
	"testing"
)

var agentadd string = "ipc://agent.gload.test"
var bossadd string = "ipc://boss.gload.test"

func TestEndToEnd(t *testing.T) {
	agent.Subscribe(agentadd)
	agent.Close()

	boss.Connect(bossadd)
	boss.SendCommand("TEST")
}
