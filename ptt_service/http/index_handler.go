package http

import (
	"encoding/json"
	"net/http"
	"ptt/domain"
)

// ServerStatus is a simple domain object for reporting server status
type ServerStatus struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// IndexHandler represents the root / uri of the server
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var status = &ServerStatus{"ok", domain.Version}
	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}
