package main

import (
	"fmt"

	consulApi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

// 使用consul源码中的watch包监听服务变化
func main() {
	var (
		err    error
		params map[string]interface{}
		plan   *watch.Plan
	)

	params = make(map[string]interface{})
	params["type"] = "service"
	params["service"] = "serverNode"
	params["passingonly"] = false
	params["tag"] = "serverNode"
	plan, err = watch.Parse(params)
	if err != nil {
		panic(err)
	}
	plan.Handler = func(index uint64, result interface{}) {
		if entries, ok := result.([]*consulApi.ServiceEntry); ok {
			for k, v := range entries {
				fmt.Println(k)
				fmt.Println("health checks = ", v.Checks.AggregatedStatus())
				fmt.Println("id = ", v.Service.ID)
			}
		}
	}

	if err = plan.Run("127.0.0.1:8500"); err != nil {
		panic(err)
	}

}

//func register() {
//	var (
//		err    error
//		client *consulApi.Client
//	)
//	client, err = consulApi.NewClient(&consulApi.Config{Address: "127.0.0.1:8500"})
//	if err != nil {
//		panic(err)
//	}
//	err = client.Agent().ServiceRegister(&consulApi.AgentServiceRegistration{
//		ID:   "",
//		Name: "test",
//		Tags: []string{"SERVER"},
//		Port: 9527,
//		Check: &consulApi.AgentServiceCheck{
//			HTTP: "",
//		},
//	})
//	if err != nil {
//		panic(err)
//	}
//}
