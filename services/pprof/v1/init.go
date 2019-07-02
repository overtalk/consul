package pprof

import (
	"net/http/pprof"

	"github.com/qinhan-shu/consul/module"
)

// AddPprof : 添加pprof
func AddPprof(gate module.Gate) {
	gate.RegisterRoute(module.Route{Method: "GET", Path: "/debug/pprof/", Handler: pprof.Index})
	gate.RegisterRoute(module.Route{Method: "GET", Path: "/debug/pprof/cmdline", Handler: pprof.Cmdline})
	gate.RegisterRoute(module.Route{Method: "GET", Path: "/debug/pprof/profile", Handler: pprof.Profile})
	gate.RegisterRoute(module.Route{Method: "GET", Path: "/debug/pprof/symbol", Handler: pprof.Symbol})
	gate.RegisterRoute(module.Route{Method: "GET", Path: "/debug/pprof/trace", Handler: pprof.Trace})
}
