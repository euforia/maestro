package main

import (
	"fmt"
	"strings"

	"github.com/euforia/loom"
)

type Symphony struct {
	Performer   string   `yaml:"performer"`
	Parallel    bool     `yaml:"parallel"`
	User        string   `yaml:"user"`
	Composition []string `yaml:"composition"` // set of commands to run on a given host
}

func (s *Symphony) Play() (err error) {
	if s.IsLocal() {
		return s.runLocal()
	} else {
		return s.runRemotes()
	}
}

func (s *Symphony) IsLocal() bool {
	return s.Performer == "local"
}

func (s *Symphony) runLocal() (err error) {
	var (
		local  loom.Config
		output string
	)

	for _, comp := range s.Composition {
		if output, err = local.Local(comp); err != nil {
			return
		}
		fmt.Printf("[local] %s\n", strings.TrimRight(output, "\n"))
	}

	return
}

func (s *Symphony) runRemote(host string) (err error) {
	var remote loom.Config
	remote.Host = host
	if len(s.User) > 0 {
		remote.User = s.User
	}

	var output string
	for _, comp := range s.Composition {
		fmt.Printf("[%s] %s\n", host, comp)
		if output, err = remote.Run(comp); err != nil {
			return
		}
		fmt.Printf("[%s] %s\n", host, strings.TrimRight(output, "\n"))
	}
	return
}

func (s *Symphony) runRemotes() (err error) {
	hostList := strings.Split(s.Performer, ",")
	if s.Parallel {
		for _, v := range hostList {
			go func() {
				host := strings.TrimSpace(v)
				if err = s.runRemote(host); err != nil {
					return
				}
			}()
		}
	} else {
		for _, v := range hostList {
			host := strings.TrimSpace(v)
			if err = s.runRemote(host); err != nil {
				return
			}
		}
	}
	return
}
