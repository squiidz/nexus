package nexus

import (
	"sync"
	"time"
)

type Starter func(n *Nexus)

type Nexus struct {
	Probes    []*Probe
	Job       []interface{}
	Errors    map[*Probe][]string
	Starter   Starter
	WaitStack sync.WaitGroup
}

func NewNexus(nbrProbes int, job ...interface{}) *Nexus {
	n := &Nexus{Job: job}
	n.NewProbes(nbrProbes)
	return n
}

func New() *Nexus {
	return &Nexus{}
}

// Start the Probes with the Default or custom starter
func (n *Nexus) Start(options ...func(*Nexus)) {
	if n.Starter != nil {
		n.Starter(n)
	} else {
		n.DefaultStarter()
	}
	for _, op := range options {
		op(n)
	}
}

// Default Nexus Jobs Starter
func (n *Nexus) DefaultStarter() {
	for _, p := range n.Probes {
		n.WaitStack.Add(1)
		go p.Work()
		time.Sleep(time.Second * 1)
	}
	n.WaitStack.Wait()
}

// The User can provid he's own starter
func (n *Nexus) SetStarter(s Starter) {
	n.Starter = s
}
