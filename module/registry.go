package module

import ()

// ServerType : 服务节点类型
type ServerType string

const (
	// WebServerType : web 服务节点
	WebServerType ServerType = "web"
)

// ConsulWatchConf : 服务配置
type ConsulWatchConf struct {
	ServerType ServerType
	Tag        string
}

// ConsulWatchServers ： 可用服务
type ConsulWatchServers struct {
	ServerType ServerType
	Servers    []string
}

// Registrar : consul 注册器
type Registrar interface {
	//Wait()             // 等待特定的服务上线
	Register() error   // 向consul注册服务
	DeRegister() error // 向consul取消注册
}

// Detector ： consul 检测器
type Detector interface {
	Watch() <-chan ConsulWatchServers // 服务watch
}
