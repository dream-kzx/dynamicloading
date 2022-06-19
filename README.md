# dynamicloading

A library for dynamically loading configurations

# Use
需要实现Observer接口，可以参考一下代码中的ObserverImpl
```go
package main

import (
	"fmt"
	"time"

	"github.com/dream-kzx/dynamicloading"
	"github.com/dream-kzx/dynamicloading/source/file"
)

type ObserverImpl struct {
}

func (t *ObserverImpl) UpdateConfig(configData []byte) {
	fmt.Println(string(configData))
}

func main() {
	fileSource := file.NewSource("conf/conf.json", true)
	impl := ObserverImpl{}

	manager := dynamicloading.New(fileSource, dynamicloading.WithPeriod(2000))
	manager.Register(&impl)
	_ = manager.Start()

	time.Sleep(30 * time.Second)
}
```

# 自定义日志
可以通过实现github.com/dream-kzx/dynamicloading/logger的Logger接口，来定制日志
```json
type DefaultLogger struct {
}

func (d DefaultLogger) Debug(msg string) {
	log.Printf("Debug | %s", msg)
}

func (d DefaultLogger) Info(msg string) {
	log.Printf("Info  | %s", msg)
}
func (d DefaultLogger) Warn(msg string) {
	log.Printf("Warn  | %s", msg)
}
func (d DefaultLogger) Error(msg string) {
	log.Printf("Error | %s", msg)
}

logger := &DefaultLogger{}

manager := dynamicloading.New(fileSource, dynamicloading.WithPeriod(2000), dynamicloading.WithLogger(logger))
```

# 自定义配置源
需要实现ConfigSource接口，可以参考source/file/file.go
```go
type ConfigSource interface {
	Read() ([]byte, error)
}
```


