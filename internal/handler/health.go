package handler

import "net/http"

// Health
// endpoint sederhana untuk cek server hidup
func Health(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, "OK", nil)
}