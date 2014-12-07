package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/squiidz/probe"
)

func main() {
	nexus := probe.NewNexus(10, GetInfo)
	nexus.SetStarter(NewStarter)
	nexus.Start()
}

func NewStarter(n *probe.Nexus) {
	for _, p := range n.Probes {
		n.WaitStack.Add(1)
		go p.Work()
		fmt.Printf("Probe #%d start working [Jobs: %d] ! \n", p.Id, len(p.Jobs))
	}
	n.WaitStack.Wait()
}

func GetInfo() []byte {
	ans := fetch()
	data, _ := ioutil.ReadAll(ans.Body)
	defer ans.Body.Close()

	fmt.Println(string(data))

	return data
}

func fetch() *http.Response {
	c := http.Client{}
	req, _ := http.NewRequest("GET", "http://google.com", nil)
	resp, _ := c.Do(req)
	return resp
}
