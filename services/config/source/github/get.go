package source

import (
	"encoding/json"
	"fmt"

	"github.com/qinhan-shu/consul/module"
)

// GetConfig return config
func (g *Github) GetConfig() (module.ConfigMap, error) {
	const fileName = "server.json"

	data, err := g.fetch(fileName)
	if err != nil {
		return nil, err
	}

	config := make(module.ConfigMap)

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if _, ok := config["ISSUCCEED"]; !ok {
		return nil, fmt.Errorf("failed to get config from config scorce (github version), error message [ %v ]", config)
	}

	return config, nil
}
