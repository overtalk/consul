package config

import (
	"sync"

	"github.com/lunny/log"

	"github.com/qinhan-shu/consul/module"
	envs "github.com/qinhan-shu/consul/services/config/source/envs"
	local "github.com/qinhan-shu/consul/services/config/source/local"
)

// Config describes Config model
type Config struct {
	rwMutex      sync.RWMutex
	configSource module.ConfigSource
	configMap    module.ConfigMap
}

// NewConfig is the constructor of config model
func NewConfig() module.Config {
	var (
		source module.ConfigSource
		err    error
	)

	log.Infof("get config from source (local file)")
	source, err = local.NewConfigSource()
	if err != nil {
		log.Errorf("failed to get config from source (local file) : %v", err)
		log.Infof("get config from source (envs)")
		source, err = envs.NewConfigSource()
		if err != nil {
			log.Errorf("failed to get config from source (envs) : %v", err)
			log.Fatalf("failed to get config from source(envs & local file)")
		}
	}

	c := &Config{
		configSource: source,
	}

	c.InitConfig()

	return c
}
