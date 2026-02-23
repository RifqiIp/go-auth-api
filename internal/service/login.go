package service

import (
	"errors"

	"github.com/RifqiIp/go-auth-api/pkg/utils"
)

func Login(email, password string) (string, error) {

	// validasi
	if email == "" || password == "" {
		return "", errors.New("email and password required")
	}

	// cek user ada
	hash, exists := users[email]
	if !exists {
		return "", errors.New("invalid email or password")
	}

	// compare password
	err := utils.ComparePassword(hash, password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// ⬇️ SET ROLE (sementara hardcode)
	role := "user"
	if email == "admin@example.com" {
		role = "admin"
	}

	// generate JWT
	token, err := utils.GenerateToken(email, role)
	if err != nil {
		return "", err
	}

	return token, nil
}
