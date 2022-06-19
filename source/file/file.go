package file

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dream-kzx/dynamicloading/api"
	"github.com/dream-kzx/dynamicloading/code"
)

var _ api.ConfigSource = (*Source)(nil)

type Source struct {
	filePath   string
	onlyChange bool
	lastTime   int64
}

func NewSource(filePath string, onlyChange bool) *Source {
	return &Source{
		filePath:   filePath,
		onlyChange: onlyChange,
	}
}

func (s *Source) Read() ([]byte, error) {
	file, err := os.OpenFile(s.filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("open file %s error: %s", s.filePath, err)
	}

	var currentTime int64
	if s.onlyChange {
		fileInfo, err := file.Stat()
		if err != nil {
			return nil, fmt.Errorf("read file[%s] info error: %s", s.filePath, err)
		}

		currentTime = fileInfo.ModTime().Unix()
		if currentTime <= s.lastTime {
			return nil, code.NotChangeError
		}
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file %s error: %s", s.filePath, err)
	}

	s.lastTime = currentTime
	return data, nil
}
