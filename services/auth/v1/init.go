package auth

import (
	"github.com/qinhan-shu/consul/module"
)

// Auth : 登陆模块
type Auth struct {
}

// NewAuth : 构造函数
func NewAuth() *Auth {
	return &Auth{}
}

// Register : 注册服务
func Register(g module.Gate) {
	m := NewAuth()
	g.RegisterRoute(module.Route{Method: "GET", Path: "/login", Handler: m.Login})
}
