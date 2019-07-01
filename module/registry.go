package module

// ServerType : 服务节点类型
type ServerType string

const (
	// WebServerType : web 服务节点
	WebServerType ServerType = "web"
)

// RegistryClient : consul 客户端
type RegistryClient interface {
	Register() error
	DeRegister() error
	GetAllServers(serverType ServerType) []string
}
