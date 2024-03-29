package gate

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/rs/cors"

	"github.com/qinhan-shu/go-utils/parse"

	"github.com/qinhan-shu/consul/module"
	"github.com/qinhan-shu/consul/services/pprof/v1"
	"github.com/qinhan-shu/consul/services/registry/registrar/v1"
)

// Gate : 网关
type Gate struct {
	routesMap sync.Map
	mux       *http.ServeMux
	server    http.Server
}

// NewGate : 构造函数
func NewGate(port int) *Gate {
	g := new(Gate)
	g.mux = http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	g.server = http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: c.Handler(g.mux),
	}

	// 增加pprof
	pprof.AddPprof(g)

	return g
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
		g.mux.HandleFunc(path, route.Handler)
		return true
	})

	consulRegistrar := registrar.NewRegistrar(8080, 9527, "web")
	consulRegistrar.Register()

	log.Printf("Starting web server on %s\n", g.server.Addr)
	if err := g.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("gate service ListenAndServe error: %v", err)
	}
}
