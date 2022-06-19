package dynamicloading

import "github.com/dream-kzx/dynamicloading/logger"

type Option func(manager *ConfigLoadManager)

func WithPeriod(period int) Option {
	return func(manager *ConfigLoadManager) {
		manager.period = period
	}
}

func WithLogger(logger logger.Logger) Option {
	return func(manager *ConfigLoadManager) {
		manager.logger = logger
	}
}
