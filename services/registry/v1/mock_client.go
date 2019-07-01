package registry

import (
	"github.com/qinhan-shu/consul/module"
)

type mockClient struct{}

func (c *mockClient) Register() error {
	return nil
}

func (c *mockClient) DeRegister() error {
	return nil
}

func (c *mockClient) GetAllServers(serverType module.ServerType) []string {
	return nil
}
