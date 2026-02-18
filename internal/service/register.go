package service

import "errors"

/*
Register
→ fungsi BUSINESS LOGIC
→ tidak tahu HTTP, JSON, status code
*/
func Register(email, password string) error {

	// 1️⃣ Aturan bisnis
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	// 2️⃣ Nanti: cek DB, hash password, dll
	// (sementara fake)

	return nil
}
