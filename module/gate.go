package module

import (
	"github.com/gin-gonic/gin"
)

// Route : 路由配置
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// Gate : 网关接口
type Gate interface {
	RegisterRoute(r Route)
}
