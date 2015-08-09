package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	symphFile     = flag.String("f", "", "Symphony file")
	configKeyword = "symph"
	configExt     = ".yml"
)

func findConfigs(dirname string) (configs []string, err error) {
	configs = make([]string, 0)

	var list []os.FileInfo
	if list, err = ioutil.ReadDir(dirname); err != nil {
		return
	}

	for _, f := range list {
		if strings.Contains(f.Name(), configKeyword) && strings.HasSuffix(f.Name(), configExt) {
			configs = append(configs, f.Name())
		}
	}

	return
}

func main() {
	flag.Parse()

	var (
		symphs = []Symphony{}
		err    error
	)

	if len(*symphFile) > 0 {

		if symphs, err = LoadSymphony(*symphFile); err != nil {
			fmt.Printf("Failed to load symphonies: %s\n", err)
			os.Exit(1)
		}

	} else {

		configs, err := findConfigs("./")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(2)
		} else if len(configs) < 1 {
			fmt.Printf("No configs found!\n")
			os.Exit(3)
		}

		fmt.Printf("Concerts found: %d %v\n", len(configs), configs)
		for _, c := range configs {
			ts, lErr := LoadSymphony(c)
			if lErr != nil {
				fmt.Printf("Failed to load symphonies: %s\n", lErr)
				os.Exit(1)
			}
			symphs = append(symphs, ts...)
		}
	}

	fmt.Printf("Symphonies loaded: %d\n", len(symphs))
	fmt.Printf("Symphonies requested: %d %v\n", len(flag.Args()), flag.Args())

	concert := NewConcert(symphs)
	if err = concert.Start(flag.Args()...); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}
}
