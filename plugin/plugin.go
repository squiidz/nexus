package plugin

import (
	"bytes"
	"fmt"

	"github.com/squiidz/nexus"
)

func CheckErr(n *nexus.Nexus) {
	for _, p := range n.Probes {
		if p.Err != nil {
			fmt.Printf("\x1b[41m[Probe #%d]\x1b[0m Error => %s\n", p.Id, p.Err.Error())
		}
	}
}

func DataSize(n *nexus.Nexus) {
	var size int
	for _, p := range n.Probes {
		b := bytes.NewBuffer([]byte(""))
		p.Extract(b)
		fmt.Println(b.String())
		size += b.Len()
	}
	fmt.Println("TOTAL SIZE :", size)
}
