package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/RifqiIp/go-auth-api/internal/response"
	"github.com/golang-jwt/jwt/v5"
)

// key PRIVATE (tidak di-export)
type contextKey string

const userEmailKey contextKey = "userEmail"

// helper ambil secret
func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// GetUserEmail
// helper resmi untuk handler
func GetUserEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(userEmailKey).(string)
	return email, ok
}

// JWTAuth middleware
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.JSON(w, http.StatusUnauthorized, "missing authorization header", nil)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.JSON(w, http.StatusUnauthorized, "invalid authorization format", nil)
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return getJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			response.JSON(w, http.StatusUnauthorized, "invalid or expired token", nil)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		email, ok := claims["email"].(string)
		if !ok {
			response.JSON(w, http.StatusUnauthorized, "invalid token payload", nil)
			return
		}

		// inject ke context
		ctx := context.WithValue(r.Context(), userEmailKey, email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}