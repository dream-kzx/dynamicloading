package api

type ConfigSource interface {
	Read() (interface{}, error)
}

type Observer interface {
	UpdateConfig(configData interface{})
}
