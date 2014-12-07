package probe

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNexusCreation(t *testing.T) {

	nex := NewNexus(3, check)
	nex.NewProbes(2)
	nex.SetStarter(newStart)

	nex.Probes[4].NewJob(fart)

	nex.Start()

	for _, p := range nex.Probes {
		b := bytes.NewBuffer([]byte(""))
		p.Extract(b)
		fmt.Println(b)
	}
}

func check() []byte {
	data, _ := ioutil.ReadFile("test/test.txt")
	return data
}

func fart() string {
	return "Sorry if it smell like death"
}

func newStart(n *Nexus) {

	for _, p := range n.Probes {
		go p.Work(disp)
		fmt.Printf("Probe #%d start working [Jobs: %d] ! \n", p.Id, len(p.Jobs))
	}

}
