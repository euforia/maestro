package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	symphFile = flag.String("f", ".symphony.yml", "Symphony file")
)

func main() {
	flag.Parse()

	orch, err := LoadOrchestra(*symphFile)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	if err = orch.Play(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}
}
