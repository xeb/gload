package main

import (
	"fmt"
	"github.com/xeb/gload/modules/boss"
)

const defbossadd string = "tcp://localhost:11000"

func main() {
	fmt.Printf("[BOSS] Connecting to Proxy at '%s'\n", defbossadd)
	boss.Connect(defbossadd)

	boss.SendCommand("date")
	boss.SendCommand("ls -la")
	//SendCommand("gload PING")
	boss.SendCommand("curl http://www.google.com")
}
