package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/RifqiIp/go-auth-api/internal/handler"
	"github.com/golang-jwt/jwt/v5"
)

// secret key
// HARUS sama dengan yang dipakai di utils/jwt.go

// ...
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// JWTAuth
// middleware untuk validasi JWT
func JWTAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1️⃣ ambil Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handler.JSON(w, http.StatusUnauthorized, "missing authorization header", nil)
			return
		}

		// 2️⃣ format harus: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			handler.JSON(w, http.StatusUnauthorized, "invalid authorization format", nil)
			return
		}

		tokenString := parts[1]

		// 3️⃣ parse & validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// pastikan signing method benar
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			handler.JSON(w, http.StatusUnauthorized, "invalid or expired token", nil)
			return
		}

		// 4️⃣ token valid → lanjut ke handler
		next.ServeHTTP(w, r)
	})
}
