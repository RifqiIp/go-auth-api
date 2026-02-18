package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RifqiIp/go-auth-api/internal/service"
)

/*
LoginRequest
→ bentuk DATA dari request body
*/
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
Login
→ HTTP handler (controller)
*/
func Login(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ Validasi HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2️⃣ Decode JSON body
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	// 3️⃣ Validasi request (layer HTTP)
	if req.Email == "" || req.Password == "" {
		http.Error(w, "email & password required", http.StatusBadRequest)
		return
	}

	// 4️⃣ Panggil SERVICE
	err = service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// 5️⃣ Response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{
		Message: "login success",
	})
}
