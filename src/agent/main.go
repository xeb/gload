package main

import (
	agent "github.com/xeb/gload/src/agent"
)

const defproxy string = "ipc://gload.proxy"

func main() {
	agent.Bind(defproxy)
}
