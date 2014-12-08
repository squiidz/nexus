package main

import (
	"io/ioutil"
	//"net/http"

	nex "github.com/squiidz/nexus"
	plug "github.com/squiidz/nexus/plugin"
)

func main() {
	nexus := nex.New()
	nexus.NewProbe().NewJob(Bla)
	pr := nexus.NewProbe()
	pr.NewJob(Blob)

	nexus.Start(plug.DataSize, plug.CheckErr)
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
