package handler

import (
	"net/http"

	"github.com/RifqiIp/go-auth-api/internal/middleware"
	"github.com/RifqiIp/go-auth-api/internal/response"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	email, ok := middleware.GetUserEmail(r.Context())
	if !ok {
		response.JSON(w, http.StatusUnauthorized, "user not found in context", nil)
		return
	}

	response.JSON(w, http.StatusOK, "profile access granted", map[string]string{
		"email": email,
	})
}