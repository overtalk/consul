package registrar

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	consulAPI "github.com/hashicorp/consul/api"

	"github.com/qinhan-shu/consul/module"
)

// Registrar : 注册器
type Registrar struct {
	mutex *sync.RWMutex

	serverPort int // consul 服务器的注册端口
	listenPort int // 服务检测 check 的端口

	serverType module.ServerType
	tag        string
	pod        podInfo

	*http.ServeMux
	checkServer  *http.Server
	servers      map[string][]string
	consulClient *consulAPI.Client
	//consulWatchPlans []*consulWatchPlan
	//watchChan        chan consulWatchServers
	//waitChan         chan serverCount
	//checkMap         map[mode.ServerType]bool
	//pvpServerIndex   int32
}

type podInfo struct {
	Name      string
	Namespace string
	IP        string
}

// NewRegistrar ： 构造函数
func NewRegistrar(serverListen, registryListen int, serverType module.ServerType) *Registrar {
	c := &Registrar{
		mutex:      new(sync.RWMutex),
		serverPort: registryListen,
		listenPort: serverListen,
		pod:        getPodInfo(),
		serverType: serverType,
		servers:    make(map[string][]string),
	}

	c.ServeMux = http.NewServeMux()
	c.checkServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", serverListen),
		Handler: c.ServeMux,
	}

	// 构造consulApi client
	config := consulAPI.DefaultConfig()
	// TODO: 从配置中获取consul地址
	config.Address = "127.0.0.1:8500"

	client, err := consulAPI.NewClient(config)
	if err != nil {
		log.Fatalf("Consul Register Err: %+v\n", err)
	}
	c.consulClient = client

	return c
}
