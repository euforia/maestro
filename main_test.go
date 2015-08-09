package main

import (
	"testing"
)

var (
	testConfigDir = "./test/"
)

func Test_findConfigs(t *testing.T) {
	list, err := findConfigs(testConfigDir)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%v", list)
}
