package registry

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"sync/atomic"

	"github.com/qinhan-shu/consul/module"
)

const (
	offline int32 = 0
	online  int32 = 1
)

// Service defines registry service
type Service struct {
	*http.ServeMux
	HTTPServer     *http.Server
	status         int32
	handlerFuncMap map[string]http.HandlerFunc
	client         module.RegistryClient
}

// NewService creates a new registry service
func NewService(registryClient module.RegistryClient) *Service {
	s := &Service{
		status: online,
	}
	s.ServeMux = http.NewServeMux()
	s.HTTPServer = &http.Server{
		Handler: s.ServeMux,
	}

	s.handlerFuncMap = map[string]http.HandlerFunc{
		"/debug/pprof/":        pprof.Index,
		"/debug/pprof/cmdline": pprof.Cmdline,
		"/debug/pprof/profile": pprof.Profile,
		"/debug/pprof/symbol":  pprof.Symbol,
		"/debug/pprof/trace":   pprof.Trace,
		"/check":               s.checkHandleFunc,
		"/mem":                 s.memStatsHandleFunc,
		"/updateStatus":        s.updateStatusHandleFunc,
	}

	s.client = new(mockClient)
	if registryClient != nil {
		s.client = registryClient
	}

	return s
}

// IsOffline checks if service is offline
func (s *Service) IsOffline() bool {
	return atomic.LoadInt32(&s.status) == offline
}

// GetStatus outputs `online` or `offline`
func (s *Service) GetStatus() string {
	if s.IsOffline() {
		return "offline"
	}
	return "online"
}

func (s *Service) register() error {
	if s.IsOffline() {
		atomic.StoreInt32(&s.status, online)
		return s.client.Register()
	}

	return nil
}

func (s *Service) deRegister() error {
	if s.IsOffline() {
		return nil
	}

	atomic.StoreInt32(&s.status, offline)
	return s.client.DeRegister()
}

func (s *Service) updateStatus(status string) error {
	switch status {
	case "offline":
		if err := s.deRegister(); err != nil {
			return err
		}
	case "online":
		if err := s.register(); err != nil {
			return err
		}
	default:
		return fmt.Errorf(`status must be set to "online" or "offline"`)
	}

	log.Printf("service is ready to %s\n", status)
	return nil
}
