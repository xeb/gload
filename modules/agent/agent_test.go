package agent

import (
	"testing"
)

func TestExec(t *testing.T) {
	output, _, sw := Exec("ls -la")
	t.Logf("Output is %s and took %dms", output, sw.Milliseconds())

	if len(output) == 0 {
		t.Error("Output is empty for Exec(\"ls -la\")")
	}
}
