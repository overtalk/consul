package registry

import (
	"fmt"
	"net/http"

	"github.com/qinhan-shu/consul/utils/perf"
)

func (s *Service) checkHandleFunc(w http.ResponseWriter, r *http.Request) {
	if s.IsOffline() {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	w.Write([]byte("")) // nolint: errcheck
}

func (s *Service) updateStatusHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.updateStatus(r.PostForm.Get("status")); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s *Service) memStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, isPretty := r.Form["pretty"]
	memStatsBytes, err := perf.GetMemStats(isPretty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(memStatsBytes) // nolint: errcheck
}
