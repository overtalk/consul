package registrar

import (
	"fmt"
	"log"
	"time"

	consulAPI "github.com/hashicorp/consul/api"
)

// Register implements registry client interface
func (c *Registrar) Register() error {
	registration := new(consulAPI.AgentServiceRegistration)
	registration.ID = getConsulID(string(c.serverType))
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

	time.Sleep(time.Second)

	c.ServeMux.HandleFunc("/check", c.updateStatusHandleFunc)

	go func() {
		if c.checkServer.Addr != "" {
			log.Printf("Start registry service on %s\n", c.checkServer.Addr)
			if err := c.checkServer.ListenAndServe(); err != nil {
				log.Fatalf("start registry service error: %v", err)
			}
		}
	}()

	return c.consulClient.Agent().ServiceRegister(registration)
}
