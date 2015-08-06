package main

import (
	"testing"
)

var (
	testOrch, _ = LoadOrchestra(testFile)
)

func Test_Orchestra_Play(t *testing.T) {
	for k, v := range testOrch {
		if err := v.Play(); err != nil {
			t.Fatalf("%s", k)
		}
	}
}

func Test_Orchestra_PlayParallel(t *testing.T) {
	testOrch["localhost,localhost"].Parallel = true
	for k, v := range testOrch {
		if err := v.Play(); err != nil {
			t.Fatalf("%s", k)
		}
	}
}

func Test_Orchestra_BadUser(t *testing.T) {
	testOrch["localhost"].User = "foobar"
	if err := testOrch["localhost"].Play(); err == nil {
		t.Fatalf("Should have failed!")
	}
}

func Test_Orchestra_NoUser(t *testing.T) {
	testOrch["localhost"].User = ""
	if err := testOrch["localhost"].Play(); err != nil {
		t.Fatalf("%s", err)
	}
}
