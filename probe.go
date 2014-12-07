package nexus

import (
	"fmt"
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
	Err  error
	Wait *sync.WaitGroup
}

func (n *Nexus) NewProbe() *Probe {
	prob := &Probe{Id: len(n.Probes) + 1, Jobs: n.Job, Wait: &n.WaitStack}
	n.Probes = append(n.Probes, prob)

	return prob
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

		for j, arg := range result {
			cont = append(cont, arg.Interface())
			switch cont[j].(type) {
			case error:
				p.Err = cont[j].(error)
			default:
				p.Data[i] = cont[j]
			}
		}
	}
	p.Wait.Done()
}

func (p *Probe) Extract(wri io.Writer) {
	for _, d := range p.Data {
		switch d.(type) {
		case string:
			n, _ := wri.Write([]byte(d.(string)))
			fmt.Println(n)
		case int:
			n, _ := wri.Write(d.([]byte))
			fmt.Println(n)
		case []byte:
			n, _ := wri.Write(d.([]byte))
			fmt.Println(n)
		default:
			fmt.Println(reflect.TypeOf(d), "IS NOT A VALID TYPE")
		}
	}
}

func (p *Probe) NewJob(j Job) {
	p.Jobs = append(p.Jobs, j)
}
