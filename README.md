# dynamicloading

A library for dynamically loading configurations

# Use

```go
package main

import (
	"fmt"
	"time"

	"github.com/dream-kzx/dynamicloading"
	"github.com/dream-kzx/dynamicloading/source/file"
)

type Test struct {
}

func (t *Test) UpdateConfig(configData []byte) {
	fmt.Println(string(configData))
}

func main() {
	fileSource := file.NewSource("conf/conf.json", true)
	test := Test{}

	manager := dynamicloading.New(fileSource, dynamicloading.WithPeriod(2000))
	manager.Register(&test)
	_ = manager.Start()

	time.Sleep(30 * time.Second)
}
```


