package main

import (
	"testing"
)

func TestExec(t *testing.T) {
	output, _ := Exec("ls -la")
	t.Logf("Output is %s", output)

	if len(output) == 0 {
		t.Error("Output is empty for Exec(\"ls -la\")")
	}
}

