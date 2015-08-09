package main

import (
	"testing"
)

var (
	testFile = "./test/symphony.yml"
)

func Test_LoadSymphony(t *testing.T) {
	symphs, err := LoadSymphony(testFile)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%#v", symphs)
}
