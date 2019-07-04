package detector

import (
	"fmt"
	"log"

	consulAPI "github.com/hashicorp/consul/api"
	consulWatch "github.com/hashicorp/consul/api/watch"

	"github.com/qinhan-shu/consul/module"
)

type monitor struct {
	plan            *consulWatch.Plan // consul 服务发现器
	changeCount     int               // 服务变更的次数
	watchServerType module.ServerType // 目标服务类型
	watchTag        string            // 目标服务的tag
}

func newMonitor(serverType module.ServerType, tag string, watchChan chan module.ConsulWatchServers) *monitor {
	p := &monitor{
		watchServerType: serverType,
		watchTag:        tag,
	}

	params := make(map[string]interface{})
	params["type"] = "service"
	params["service"] = string(p.watchServerType)
	if p.watchTag != "" {
		params["tag"] = p.watchTag
	}

	plan, err := consulWatch.Parse(params)
	if err != nil {
		log.Fatalf("Consul New Watch Plan  Err: %+v\n", err)
	}
	p.plan = plan
	p.plan.Handler = func(index uint64, raw interface{}) {
		if raw == nil {
			return
		}
		if entries, ok := raw.([]*consulAPI.ServiceEntry); ok {
			var servers []string
			for _, entry := range entries {
				// 如果服务没有通过健康检查，直接continue
				if entry.Checks.AggregatedStatus() != consulAPI.HealthPassing {
					continue
				}
				server := fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port)
				//if p.watchServerType == module.WebServerType {
				//	server = fmt.Sprintf("%s:%d", entry.Service.ID, entry.Service.Port)
				//}
				if p.changeCount == 0 {
					//if err := pingServer(p.watchServerType, server); err != nil {
					//	continue
					//}
				}
				servers = append(servers, server)
			}
			if len(servers) > 0 {
				// 可以加上ping方法
				//if p.watchServerType == mode.MatchServerType {
				//	sort.Strings(servers)
				//}
			} else {
				// first service change could be critical status
				if p.changeCount > 2 {
					//logger.GetLogger().Errorf("%s Server All Crashed.", p.watchServerType)
				}
			}
			p.changeCount++
			watchChan <- module.ConsulWatchServers{
				ServerType: serverType,
				Servers:    servers,
			}
		}
	}
	return p
}
