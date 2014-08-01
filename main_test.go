package main

import (
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWikiCommand(t *testing.T) {
	Convey("wiki", t, func() {
		Convey("--help outputs the command line options", func() {
			out, _ := execGo("run", "main.go", "--help")

			So(out, ShouldContainSubstring, "Path to the BoltDB file")
		})
	})
}

func execGo(args ...string) (string, error) {
	out, err := exec.Command("go", args...).CombinedOutput()

	return string(out), err
}
