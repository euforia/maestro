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

	for _, s := range symphs {
		if err = s.Play(); err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(2)
		}
	}
}
