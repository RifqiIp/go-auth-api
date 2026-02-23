package middleware

import (
	"net/http"

	"github.com/RifqiIp/go-auth-api/internal/response"
)

// AdminOnly
// → memastikan user punya role = admin
func AdminOnly(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		role, ok := GetUserRole(r.Context())
		if !ok {
			response.JSON(w, http.StatusUnauthorized, "role not found in context", nil)
			return
		}

		if role != "admin" {
			response.JSON(w, http.StatusForbidden, "admin access only", nil)
			return
		}

		// lanjut ke handler
		next.ServeHTTP(w, r)
	})
}