package detector

import (
	"log"

	"github.com/qinhan-shu/consul/module"
)

type Detector struct {
	watchChan chan module.ConsulWatchServers
	confs     []module.ConsulWatchConf
	monitors  []*monitor
}

func NewDetector(c ...module.ConsulWatchConf) *Detector {
	m := &Detector{
		watchChan: make(chan module.ConsulWatchServers),
		confs:     c,
	}
	for _, v := range c {
		m.monitors = append(m.monitors, newMonitor(v.ServerType, v.Tag, m.watchChan))
	}
	return m
}

func (m *Detector) Watch() <-chan module.ConsulWatchServers {
	for _, v := range m.monitors {
		go func(v *monitor) {
			if err := v.plan.Run("127.0.0.1:8500"); err != nil {
				log.Printf("Consul Watch Err: %+v\n", err)
			}
		}(v)
	}
	return m.watchChan
}
