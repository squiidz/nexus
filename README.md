nexus
=====

Simple Worker Lib

### Example

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

### TODO

- Add Workload spliter
- Modify the easy of use for big task

### License

MIT 