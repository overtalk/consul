package module

import (
	"net/http"
)

// Route : 路由配置
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Gate : 网关接口
type Gate interface {
	RegisterRoute(r Route)
}
