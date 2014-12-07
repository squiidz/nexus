package probe

import (
	"io"
	"reflect"
	"sync"
)

type Job interface{}

type Probe struct {
	Id   int
	Info interface{}
	Data map[int]interface{}
	Jobs []interface{}
	Errs string
	Wait *sync.WaitGroup
}

func (n *Nexus) NewProbe() {
	n.Probes = append(n.Probes, &Probe{Id: len(n.Probes) + 1, Jobs: n.Job, Wait: &n.WaitStack})
}

func (n *Nexus) NewProbes(ps int) {
	for i := 0; i < ps; i++ {
		n.Probes = append(n.Probes, &Probe{Id: len(n.Probes) + 1, Jobs: n.Job, Wait: &n.WaitStack})
	}
}

func (p *Probe) Work() {
	p.Data = make(map[int]interface{}, len(p.Jobs))
	for i, j := range p.Jobs {
		task := reflect.ValueOf(j)
		in := make([]reflect.Value, 0)

		result := task.Call(in)

		var cont []interface{}

		for _, arg := range result {
			cont = append(cont, arg.Interface())
		}
		p.Data[i] = cont

	}
	p.Wait.Done()
}

func (p *Probe) Extract(wri io.Writer) {
	for _, d := range p.Data {
		switch d.(type) {
		case string:
			wri.Write([]byte(d.(string)))
		case int:
			wri.Write(d.([]byte))
		case []byte:
			wri.Write(d.([]byte))
		}
	}
}

func (p *Probe) NewJob(j Job) {
	p.Jobs = append(p.Jobs, j)
}
