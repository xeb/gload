package main

import (
	agent "github.com/xeb/gload/src/agent"
)

const discovery string = "http://localhost:4001/v2/keys/"

func main() {
	agent.Bind(10000)
}
