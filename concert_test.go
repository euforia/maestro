package main

import (
	"testing"
)

var (
	testSymphonies, _ = LoadSymphonies(testFile)
)

func Test_Concert_Start(t *testing.T) {
	concert := NewConcert(testSymphonies)
	if err := concert.Start(); err != nil {
		t.Fatalf("%s", err)
	}
}

func Test_Concert_Start_ByNames(t *testing.T) {
	concert := NewConcert(testSymphonies)
	if err := concert.Start("date", "hostname"); err != nil {
		t.Fatalf("%s", err)
	}
}

func Test_Concert_Start_NameNotFound(t *testing.T) {
	concert := NewConcert(testSymphonies)
	if err := concert.Start("foo"); err == nil {
		t.Fatalf("Should have errored!")
	}
}
