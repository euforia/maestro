package main

import (
	"testing"
)

var (
	testFile = "./test/symphony.yml"
)

func Test_LoadSymphonies(t *testing.T) {
	symphs, err := LoadSymphonies(testFile)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%#v", symphs)
}

/*
func Test_Orchestra_Play(t *testing.T) {
	for k, v := range testOrch.symphonies {
		if err := v.Play(); err != nil {
			t.Fatalf("%s", k)
		}
	}
}

func Test_Orchestra_PlayParallel(t *testing.T) {
	testOrch.symphonies["localhost,localhost"].Parallel = true
	for k, v := range testOrch.symphonies {
		if err := v.Play(); err != nil {
			t.Fatalf("%s", k)
		}
	}
}

func Test_Orchestra_BadUser(t *testing.T) {
	testOrch.symphonies["localhost"].User = "foobar"
	if err := testOrch.symphonies["localhost"].Play(); err == nil {
		t.Fatalf("Should have failed!")
	}
}

func Test_Orchestra_NoUser(t *testing.T) {
	testOrch.symphonies["localhost"].User = ""
	if err := testOrch.symphonies["localhost"].Play(); err != nil {
		t.Fatalf("%s", err)
	}
}
*/
