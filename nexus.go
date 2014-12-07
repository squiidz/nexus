package probe

import (
	//	"errors"
	//	"fmt"
	"sync"
)

type Starter func(n *Nexus)

type Nexus struct {
	Probes    []*Probe
	Job       []interface{}
	Errors    map[*Probe][]string
	Starter   Starter
	WaitStack sync.WaitGroup
}

func NewNexus(inst int, job ...interface{}) *Nexus {
	n := &Nexus{Job: job}
	n.NewProbes(inst)
	return n
}

func (n *Nexus) Start() {
	if n.Starter != nil {
		n.Starter(n)
	} else {

		for _, p := range n.Probes {
			go p.Work()
		}
	}
}

func (n *Nexus) SetStarter(s Starter) {
	n.Starter = s
}
