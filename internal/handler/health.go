package handler

import (
	"net/http"
	"github.com/RifqiIp/go-auth-api/internal/response"
)

// Health
// endpoint sederhana untuk cek server hidup
func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "OK", nil)
}