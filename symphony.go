package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/euforia/loom"
)

type Symphony struct {
	Name        string   `yaml:"name"`
	Performers  []string `yaml:"performer"`
	Parallel    bool     `yaml:"parallel"`
	User        string   `yaml:"user"`
	Composition []string `yaml:"composition"` // set of commands to run on a given host
}

func (s *Symphony) isLocal(performer string) bool {
	return performer == "local"
}

func (s *Symphony) runLocal() (err error) {
	var (
		local  loom.Config
		output string
	)

	for _, comp := range s.Composition {
		fmt.Printf("[local] %s\n", comp)
		if output, err = local.Local(comp); err != nil {
			return
		}
		for _, v := range strings.Split(output, "\n") {
			fmt.Printf("[local] %s", v)
		}
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
		fmt.Printf("[%s] %s", host, strings.TrimRight(output, "\n"))
	}
	return
}

func (s *Symphony) singleRun(host string) (err error) {
	if s.isLocal(strings.TrimSpace(host)) {
		return s.runLocal()
	} else {
		return s.runRemote(strings.TrimSpace(host))
	}
	return nil
}

func (s *Symphony) Play() (err error) {
	if s.Parallel {
		for _, v := range s.Performers {
			go func() {
				if err = s.singleRun(v); err != nil {
					return
				}
			}()
		}
	} else {
		for _, v := range s.Performers {
			if err = s.singleRun(v); err != nil {
				return
			}
		}
	}
	return
}

func LoadSymphony(cfgfile string) (symphs []Symphony, err error) {
	var (
		b []byte
	)
	if b, err = ioutil.ReadFile(cfgfile); err != nil {
		return
	}

	if err = yaml.Unmarshal(b, &symphs); err != nil {
		return
	}
	return
}
