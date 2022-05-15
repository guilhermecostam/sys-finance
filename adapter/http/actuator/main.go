package actuator

import (
	"encoding/json"
	"net/http"
)

// Health is a method to indicate server is alive or not
func Health(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("content-type", "application/json")

	profile := HealthBody{"alive"}

	returnBody, err := json.Marshal(profile)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = responseWriter.Write(returnBody)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// HealthBody is a struct for the health type return
type HealthBody struct {
	Status string `json:"status"`
}
