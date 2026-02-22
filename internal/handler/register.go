package handler // package ini bernama handler, dipakai untuk HTTP layer

import (
	"encoding/json" // untuk decode & encode JSON
	"net/http"      // package utama HTTP di Go

	"github.com/RifqiIp/go-auth-api/internal/service" // panggil logic bisnis
)

// RegisterRequest
// struct ini merepresentasikan DATA yang dikirim client lewat HTTP body
// hanya dipakai di layer handler
type RegisterRequest struct {
	Email    string `json:"email"`    // ambil field "email" dari JSON
	Password string `json:"password"` // ambil field "password" dari JSON
}

// Register
// fungsi ini adalah HTTP HANDLER
// akan dipanggil saat endpoint /register diakses
func Register(w http.ResponseWriter, r *http.Request) {

	// 1️⃣ cek apakah method HTTP adalah POST
	// jika bukan POST → tolak request
	if r.Method != http.MethodPost {
		JSON(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return // hentikan fungsi
	}

	// 2️⃣ siapkan variabel untuk menampung body request
	var req RegisterRequest

	// 3️⃣ decode JSON dari body ke struct RegisterRequest
	// contoh body:
	// { "email": "...", "password": "..." }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// jika JSON tidak valid
		JSON(w, http.StatusBadRequest, "invalid json", nil)
		return
	}

	// 4️⃣ panggil service.Register untuk logic bisnis
	// service akan melakukan validasi (email kosong, password kosong, dll)
	if err := service.Register(req.Email, req.Password); err != nil {
		// err adalah TYPE error → ubah ke string dengan err.Error()
		JSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// 5️⃣ jika semua sukses, kirim response berhasil
	JSON(w, http.StatusCreated, "register success", nil)
}