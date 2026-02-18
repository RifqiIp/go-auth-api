package main

import (
	"fmt"
	"net/http"

	"github.com/RifqiIp/go-auth-api/internal/handler"
)

func main() {
	fmt.Println("Server running on :8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	http.ListenAndServe(":8080", nil)
}
