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

	symphs, err := LoadSymphonies(*symphFile)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	concert := NewConcert(symphs)
	if err = concert.Start(flag.Args()...); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}
}
