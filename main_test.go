package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestUsageInstructions(t *testing.T) {
	out, _ := execGo("run", "main.go", "--help")

	if !strings.Contains(out, "Path to the BoltDB file") {
		t.Fatalf(`unexpected output`)
	}
}

func execGo(args ...string) (string, error) {
	out, err := exec.Command("go", args...).CombinedOutput()

	return string(out), err
}
