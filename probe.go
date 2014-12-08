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

func (p *Probe) NewJob(j Job) {
	p.Jobs = append(p.Jobs, j)
}

func (p *Probe) Work() {
	p.Data = make(map[int]interface{}, len(p.Jobs))
	fmt.Printf("Probe #%d Start Working \n", p.Id)
	for i, j := range p.Jobs {

		task := reflect.ValueOf(j)
		//DEBUG: fmt.Println(reflect.TypeOf(j))
		in := []reflect.Value{}

		result := task.Call(in)

		for _, arg := range result {
			tmp := arg.Interface()

			switch tmp.(type) {
			case error:
				p.Err = tmp.(error)
			default:
				p.Data[i] = tmp
			}
		}
	}
	p.Wait.Done()
}

func (p *Probe) Extract(wri io.Writer) string {
	var n int
	for _, d := range p.Data {
		switch d.(type) {
		case string:
			n, _ = wri.Write([]byte(d.(string)))
		case int:
			n, _ = wri.Write(d.([]byte))
		case []byte:
			n, _ = wri.Write(d.([]byte))
		default:
			fmt.Println(reflect.TypeOf(d), "IS NOT A VALID TYPE")
		}
		//DEBUG: fmt.Println("TYPE IS :", reflect.TypeOf(d))
	}

	return fmt.Sprintf("Probe #%d: %d data writed", p.Id, n)
}

func (p *Probe) Stats() {
}
