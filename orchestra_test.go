package main

import (
	"encoding/json"
	"testing"
)

var (
	testFile = "./test/symphony.yml"
)

func Test_LoadOrchestra(t *testing.T) {
	if testOrch, err := LoadOrchestra(testFile); err != nil {
		t.Fatalf("%s", testOrch)
	} else {
		b, _ := json.MarshalIndent(testOrch, "", " ")
		t.Logf("%s", b)
	}
}
