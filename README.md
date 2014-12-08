Nexus
=====

### Simple Golang Worker
Multiplex your Jobs to multiple worker.

### Example #1

``` go

package main 

import(
	"github.com/squiidz/nexus"
)

func main() {
	// NewNexus(Number of Probes, Jobs to Do)
	man := nexus.NewNexus(3, Job)
	man.Start()
}

func Job() string{
	return "Job Done"
}

```

### Example #2

``` go

package main 

import(
	"github.com/squiidz/nexus"
)

func main() {
	// New(), Create a empty Nexus
	man := nexus.New()

	// Add some probe to the Nexus, and set Jobs on them
	man.NewProbe().NewJob(Job)
	man.NewProbe().NewJob(Job2)

	// Start Working
	man.Start()
}

func Job() string{
	return "Job Done"
}

func Job2() []byte{
	return []byte("Some Bytes")
}

```

### TODO

- Add Workload spliter
- Modify the easy of use for big task

### License

MIT 