package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Orchestra map[string]*Symphony

func (o Orchestra) Play() error {
	for _, v := range o {
		if err := v.Play(); err != nil {
			return err
		}
	}
	return nil
}

func LoadOrchestra(cfgfile string) (Orchestra, error) {
	var orch Orchestra
	b, err := ioutil.ReadFile(cfgfile)
	if err != nil {
		return orch, err
	}
	if err = yaml.Unmarshal(b, &orch); err != nil {
		return orch, err
	}
	for k, v := range orch {
		v.Performer = k
	}
	return orch, nil
}
