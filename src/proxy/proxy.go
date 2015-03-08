package main

import (
	"fmt"
	proxy "github.com/xeb/gload/src/proxy"
)

const defbossadd string = "tcp://*:11000"
const defagentadd string = "tcp://*:20000"

func main() {

	fmt.Printf("[PROXY] Listening for Agents at '%s'\n", defagentadd)
	fmt.Printf("[PROXY] Listening for Boss at '%s'\n", defbossadd)
	proxy.Bind(defagentadd, defbossadd)
}
