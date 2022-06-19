package dynamicloading

import (
	"fmt"
	"time"

	"github.com/dream-kzx/dynamicloading/api"
	"github.com/dream-kzx/dynamicloading/code"
	"github.com/dream-kzx/dynamicloading/logger"
)

type ConfigLoadManager struct {
	configSource api.ConfigSource

	observers []api.Observer

	logger logger.Logger

	ticker *time.Ticker
	// period 动态读取配置的周期，单位为ms, 默认10s
	period int
}

func New(source api.ConfigSource, options ...Option) *ConfigLoadManager {
	manager := &ConfigLoadManager{
		configSource: source,
		observers:    make([]api.Observer, 0),
		period:       10000,
	}

	for _, option := range options {
		option(manager)
	}

	if manager.logger == nil {
		manager.logger = &logger.DefaultLogger{}
	}

	return manager
}

func (m *ConfigLoadManager) Start() error {
	m.ticker = time.NewTicker(time.Millisecond * time.Duration(m.period))
	go func() {
		m.loop()
	}()
	return nil
}

func (m *ConfigLoadManager) Stop() {
	m.ticker.Stop()
}

func (m *ConfigLoadManager) Register(observer api.Observer) {
	m.observers = append(m.observers, observer)
}

func (m *ConfigLoadManager) loop() {
	for range m.ticker.C {
		data, err := m.configSource.Read()
		if err != nil {
			if err == code.NotChangeError {
				m.logger.Debug(err.Error())
				continue
			}

			m.logger.Error(fmt.Sprintf("read source failed: %s", err.Error()))
			continue
		}
		m.notify(data)
	}
}

func (m *ConfigLoadManager) notify(data []byte) {
	m.logger.Debug(fmt.Sprintf("notify config: %s", data))

	for i := range m.observers {
		m.observers[i].UpdateConfig(data)
	}
}
