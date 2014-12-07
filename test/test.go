package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	//"net/http"

	nex "github.com/squiidz/nexus"
)

func main() {
	nexus := nex.EmptyNexus()
	nexus.NewProbe().NewJob(Bla)
	nexus.NewProbe().NewJob(Blob)
	nexus.NewProbe().NewJob(Blip)

	nexus.Start(LogAct, CheckErr)
}

func LogAct(n *nex.Nexus) {
	var size int
	for _, p := range n.Probes {
		b := bytes.NewBuffer([]byte(""))
		p.Extract(b)
		fmt.Println(b.String())
		size += b.Len()
	}
	fmt.Println("TOTAL SIZE :", size)
}

func CheckErr(n *nex.Nexus) {
	for _, p := range n.Probes {
		fmt.Printf("Probe #%d: %s\n", p.Id, p.Err.Error())
	}
}

func Bla() error {
	_, err := ioutil.ReadFile("oups.txt")
	if err != nil {
		return err
	}
	return nil
}

func Blob() error {
	_, err := ioutil.ReadFile("moule.txt")
	if err != nil {
		return err
	}
	return nil
}

func Blip() int {
	i := 9
	return i * i
}

/*
func GetInfo() []byte {
	ans := fetch()
	data, err := ioutil.ReadAll(ans.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer ans.Body.Close()

	return data
}

func fetch() *http.Response {
	c := http.Client{}
	req, _ := http.NewRequest("GET", "http://google.com", nil)
	resp, _ := c.Do(req)
	return resp
}
*/
