package registrar

import (
	"fmt"
	"log"

	consulAPI "github.com/hashicorp/consul/api"
)

// Register implements registry client interface
func (c *Registrar) Register() error {
	c.ServeMux.HandleFunc("/check", c.updateStatusHandleFunc)
	go func() {
		if c.checkServer.Addr != "" {
			log.Printf("Starting %s consul check server on %s\n", c.serverType, c.checkServer.Addr)
			if err := c.checkServer.ListenAndServe(); err != nil {
				log.Fatalf("Starting %s registry Starting error: %v", c.serverType, err)
			}
		}
	}()

	registration := new(consulAPI.AgentServiceRegistration)
	registration.ID = getConsulID(c.pod)
	registration.Name = string(c.serverType)
	registration.Port = c.serverPort
	if c.tag != "" {
		registration.Tags = []string{c.tag}
	}
	registration.Address = c.pod.IP
	registration.Check = &consulAPI.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, c.listenPort, "/check"),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s", //check失败后15秒删除本服务
	}

	return c.consulClient.Agent().ServiceRegister(registration)
}

func (c *Registrar) DeRegister() error {
	return c.consulClient.Agent().ServiceDeregister(getConsulID(c.pod))
}
