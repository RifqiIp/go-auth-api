package handler

import (
	"encoding/json"
	"net/http"
)

// Response standar untuk semua API
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// helper kirim JSON response
func JSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{
		Message: message,
		Data:    data,
	})
}