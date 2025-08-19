package health

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/centrifugal/centrifuge"
)

// FullConfig of full health check handler.
type FullConfig struct{}

// FullHandler handles detailed health endpoint.
type FullHandler struct {
	node   *centrifuge.Node
	config FullConfig
}

// NewFullHandler creates new FullHandler.
func NewFullHandler(n *centrifuge.Node, c FullConfig) *FullHandler {
	h := &FullHandler{
		node:   n,
		config: c,
	}
	return h
}

type HealthResponse struct {
	Status    string                 `json:"status"`
	Timestamp time.Time              `json:"timestamp"`
	Uptime    string                 `json:"uptime"`
	Node      string                 `json:"node"`
	Version   string                 `json:"version"`
	Metrics   map[string]interface{} `json:"metrics"`
}

func (h *FullHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Build detailed health response
	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC(),
		Uptime:    getUptime(),
		Node:      h.node.ID(),
		Version:   "6.2.5", // You can make this dynamic
		Metrics: map[string]interface{}{
			"num_clients":  h.node.Hub().NumClients(),
			"num_users":    h.node.Hub().NumUsers(),
			"num_channels": h.node.Hub().NumChannels(),
		},
	}

	json.NewEncoder(w).Encode(response)
}

func getUptime() string {
	// Simple uptime calculation - you can enhance this
	return "running"
}
