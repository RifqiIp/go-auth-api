package handler

import (
	"net/http"

	"github.com/RifqiIp/go-auth-api/internal/response"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {

	response.JSON(w, http.StatusOK, "admin dashboard", map[string]string{
		"status": "ok",
	})
}