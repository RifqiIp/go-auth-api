package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RifqiIp/go-auth-api/internal/service"
)

/*
RegisterRequest
→ struktur DATA dari HTTP request
→ hanya dipakai di handler
*/
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
Response
→ struktur response JSON
*/
type Response struct {
	Message string `json:"message"`
}

/*
Register
→ fungsi ini adalah HTTP HANDLER
→ tugasnya:
  1. validasi HTTP
  2. panggil service
  3. kirim HTTP response
*/
func Register(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ Pastikan METHOD POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2️⃣ Decode JSON BODY
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	// 3️⃣ Validasi REQUEST (bukan bisnis)
	if req.Email == "" || req.Password == "" {
		http.Error(w, "email & password required", http.StatusBadRequest)
		return
	}

	// 4️⃣ Panggil SERVICE (bisnis logic)
	err = service.Register(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 5️⃣ Response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	json.NewEncoder(w).Encode(Response{
		Message: "register success",
	})
}
