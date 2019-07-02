package config

import (
	"fmt"
	"log"
)

// InitConfig is to get config
// Fatal if failed to get config
func (c *Config) InitConfig() {
	configMap, err := c.configSource.GetConfig()
	if err != nil {
		log.Fatalf("failed to init config : %v", err)
	}

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	c.configMap = configMap
}

// ReloadConfig is to reload config
// error if failed to get config
func (c *Config) ReloadConfig() error {
	configMap, err := c.configSource.GetConfig()
	if err != nil {
		return err
	}

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	c.configMap = configMap

	return nil
}

// GetConfigByName is to get config value by config key
func (c *Config) GetConfigByName(configName string) (string, error) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	configValue := c.configMap[configName]
	if configValue == "" {
		return "", fmt.Errorf("missing config : %s", configName)
	}

	return configValue, nil
}
