nexus
=====

Simple Worker Lib

### Example #1

``` go

package main 

import(
	"github.com/squiidz/nexus"
)

func main() {
	// NewNexus(Number of worker, Job to Do)
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
	// EmptyNexus(), set parameter by hand
	man := nexus.EmptyNexus()
	
	// Set some probe to work
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