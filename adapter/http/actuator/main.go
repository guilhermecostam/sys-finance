package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody is a struct for the health type return
type HealthBody struct {
	Status string `json:"status"`
}

// Health is a method to indicate server is alive or not
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("content-type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)
}
