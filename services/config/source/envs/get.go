package source

import (
	"fmt"
	"os"

	"github.com/qinhan-shu/consul/module"
)

// GetConfig return config
func (e *Env) GetConfig() (module.ConfigMap, error) {
	e.Lock()
	defer e.Unlock()

	conf := make(map[string]string)
	for _, key := range e.require {
		value, isExist := os.LookupEnv(key)
		if !isExist {
			err := fmt.Errorf(`Config "%s" is absent`, key)
			return nil, err
		}
		conf[key] = value
	}

	return conf, nil
}
