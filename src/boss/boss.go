package main

import (
	"fmt"
	boss "github.com/xeb/gload/src/boss"
)

const defbossadd string = "tcp://localhost:11000"

func main() {
	fmt.Printf("[BOSS] Connecting to Proxy at '%s'\n", defbossadd)
	boss.Connect(defbossadd)
	boss.SendCommand("PING")
}
