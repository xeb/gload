package main

import (
	"fmt"
	agent "github.com/xeb/gload/src/agent"
)

const defagentadd string = "tcp://localhost:20000"

func main() {
	fmt.Printf("[AGENT] Connecting to Proxy at '%s'\n", defagentadd)
	agent.Subscribe(defagentadd)
}
