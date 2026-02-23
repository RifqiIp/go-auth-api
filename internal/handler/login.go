package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RifqiIp/go-auth-api/internal/response"
	"github.com/RifqiIp/go-auth-api/internal/service"
)

// LoginRequest
// data yang dikirim client saat login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login
// HTTP handler untuk endpoint /login
func Login(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ hanya izinkan POST
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}

	// 2️⃣ decode JSON body
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	// 3️⃣ panggil service login
	token, err := service.Login(req.Email, req.Password)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	// 4️⃣ response sukses
	response.JSON(w, http.StatusOK, "login success", map[string]string{
		"token": token,
	})
}