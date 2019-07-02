package module

import ()

// ServerType : 服务节点类型
type ServerType string

const (
	// WebServerType : web 服务节点
	WebServerType ServerType = "web"
)

type consulWatchServers struct {
	serverType ServerType
	servers    []string
}

// Registrar : consul 注册器
type Registrar interface {
	//Wait()             // 等待特定的服务上线
	Register() error // 向consul注册服务
	//DeRegister() error // 向consul取消注册
}

// Detector ： consul 检测器
type Detector interface {
	Watch(serverType ServerType) <-chan consulWatchServers // 服务watch
	GetAllServers(serverType ServerType) []string          // 服务发现
}
