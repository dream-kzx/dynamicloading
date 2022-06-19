package api

type ConfigSource interface {
	Read() ([]byte, error)
}

type Observer interface {
	UpdateConfig(configData []byte)
}
