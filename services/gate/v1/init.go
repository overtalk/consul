package gate

import (
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/qinhan-shu/consul/module"
	"github.com/qinhan-shu/consul/utils/parse"
)

// Gate : 网关
type Gate struct {
	port      int
	engine    *gin.Engine
	routesMap sync.Map
}

// NewGate : 构造函数
func NewGate() *Gate {
	return &Gate{
		port:   8000,
		engine: gin.Default(),
	}
}

// RegisterRoute : 注册服务
func (g *Gate) RegisterRoute(r module.Route) {
	if _, ok := g.routesMap.Load(r.Path); ok {
		log.Fatalf("repeated path : %s", r.Path)
	}
	g.routesMap.Store(r.Path, r)
}

// Start : 启动服务
func (g *Gate) Start() {
	g.routesMap.Range(func(k, v interface{}) bool {
		path, err := parse.StringWithError(k)
		if err != nil {
			log.Fatalf("illegal http path[%v], not string, parse error [%v]", k, err)
		}
		route := v.(module.Route)

		switch route.Method {
		case "GET":
			{
				g.engine.GET(path, route.Handler)
			}
		case "POST":
			{
				g.engine.POST(path, route.Handler)
			}
		default:
			{
				log.Fatalf("illegal http method[%s]", route.Method)
			}
		}
		return true
	})

	g.engine.Run(fmt.Sprintf(":%d", g.port))
}
