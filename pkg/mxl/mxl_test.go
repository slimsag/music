package mxl

import (
	"os"
	"os/exec"
	"testing"
)

func TestBasic(t *testing.T) {
	doc, err := UnmarshalFile("basic.want.xml")
	if err != nil {
		t.Fatal(err)
	}
	if err := MarshalFile("basic.got.xml", doc); err != nil {
		t.Fatal(err)
	}
	diffEqual(t, "basic.want.xml", "basic.got.xml")
}

func diffEqual(t *testing.T, a, b string) {
	cmd := exec.Command("git", "diff", "--no-index", a, b)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
}
