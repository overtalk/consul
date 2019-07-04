package main

import (
	"fmt"

	"github.com/qinhan-shu/consul/module"
	"github.com/qinhan-shu/consul/services/registry/detector/v1"

	consulApi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

// 使用consul源码中的watch包监听服务变化
func main1() {
	var (
		err    error
		params map[string]interface{}
		plan   *watch.Plan
	)

	params = make(map[string]interface{})
	params["type"] = "service"
	params["service"] = "web"
	params["passingonly"] = false
	//params["tag"] = "serverNode"
	plan, err = watch.Parse(params)
	if err != nil {
		panic(err)
	}
	plan.Handler = func(index uint64, result interface{}) {
		fmt.Println("应该是服务发生了变化")
		if result == nil {
			return
		}

		if entries, ok := result.([]*consulApi.ServiceEntry); ok {
			for _, entry := range entries {
				if entry.Checks.AggregatedStatus() != consulApi.HealthPassing {
					fmt.Printf("ID = %s, address = %s 下线了\n", entry.Service.ID, entry.Service.Address)
					continue
				}
				fmt.Printf("ID = %s, address = %s 已经上线\n", entry.Service.ID, entry.Service.Address)
			}
		}
	}

	if err = plan.Run("127.0.0.1:8500"); err != nil {
		panic(err)
	}

}

// 使用consul源码中的watch包监听服务变化
func main() {
	m := detector.NewDetector(module.ConsulWatchConf{module.WebServerType, ""})
	for v := range m.Watch() {
		fmt.Println(v)
	}
}
