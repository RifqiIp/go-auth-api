package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/RifqiIp/go-auth-api/internal/handler"
	"github.com/RifqiIp/go-auth-api/internal/middleware"
)

func main() {

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on :" + port)

	// public routes
	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/login", handler.Login)

	// protected
	http.Handle("/profile", middleware.JWTAuth(http.HandlerFunc(handler.Profile)))

	http.Handle(
	"/admin",
	middleware.JWTAuth(
		middleware.AdminOnly(
			http.HandlerFunc(handler.AdminDashboard),
		),
	),
)

	http.ListenAndServe(":"+port, nil)
}