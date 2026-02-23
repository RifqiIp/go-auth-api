package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	// middleware/jwt.go

	"github.com/RifqiIp/go-auth-api/internal/response"
	"github.com/golang-jwt/jwt/v5"
)

// key type khusus (BEST PRACTICE)
type contextKey string

const UserEmailKey contextKey = "userEmail"

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// JWTAuth
// middleware validasi JWT + inject email ke context
func JWTAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1️⃣ ambil header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.JSON(w, http.StatusUnauthorized, "missing authorization header", nil)
			return
		}

		// 2️⃣ format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.JSON(w, http.StatusUnauthorized, "invalid authorization format", nil)
			return
		}

		tokenString := parts[1]

		// 3️⃣ parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// pastikan method sesuai
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return getJWTSecret(), nil
		})

		if err != nil || !token.Valid {
			response.JSON(w, http.StatusUnauthorized, "invalid or expired token", nil)
			return
		}

		// 4️⃣ ambil email dari claims
		claims := token.Claims.(jwt.MapClaims)
		email, ok := claims["email"].(string)
		if !ok {
			response.JSON(w, http.StatusUnauthorized, "invalid token payload", nil)
			return
		}

		// 5️⃣ inject email ke context
		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		r = r.WithContext(ctx)

		// 6️⃣ lanjut ke handler
		next.ServeHTTP(w, r)
	})
}
