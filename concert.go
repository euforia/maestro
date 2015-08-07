package main

import (
	"fmt"
)

type Concert struct {
	Symphonies []Symphony
}

func NewConcert(symphs []Symphony) *Concert {
	return &Concert{symphs}
}

func (c *Concert) PlaySymphony(name string) (err error) {
	for _, v := range c.Symphonies {
		if v.Name == name {
			if err = v.Play(); err != nil {
				return
			}
			return
		}
	}
	err = fmt.Errorf("Symphony '%s' not found!", name)
	return
}

func (c *Concert) Start(symphonies ...string) (err error) {
	if len(symphonies) == 0 {
		// Play all symphonies
		for _, s := range c.Symphonies {
			if err = s.Play(); err != nil {
				return
			}
		}
	} else {
		// Play specifics by name
		fmt.Printf("Playing: %v\n", symphonies)
		for _, v := range symphonies {
			if err = c.PlaySymphony(v); err != nil {
				return
			}
		}
	}
	return
}
