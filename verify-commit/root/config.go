package root

import (
	"errors"
	"github.com/mitchellh/mapstructure"
)

type Config struct {

	// branch:
	//   "*" :
	//     .png: true
	//   directory:
	//     "*": true
	IncludeOnly map[string]map[string]map[string]bool
	UnityMeta   bool
}

var errWrongConfig = errors.New("错误的配置格式")

func extractConfig(input interface{}) (config Config, err error) {
	if err = mapstructure.Decode(input, &config); err != nil {
		return
	}

	if config.IncludeOnly == nil {
		return config, errWrongConfig
	}

	return
}
