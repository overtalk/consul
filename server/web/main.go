package main

import (
	"github.com/qinhan-shu/consul/services/auth/v1"
	"github.com/qinhan-shu/consul/services/gate/v1"
)

func main() {
	webGate := gate.NewGate(8000)
	auth.Register(webGate)
	webGate.Start()
}
