package service

import (
	"errors"

	"github.com/RifqiIp/go-auth-api/pkg/utils"
)

func Register(email, password string) error {

	if email == "" {
		return errors.New("email is required")
	}

	if password == "" {
		return errors.New("password is required")
	}
	// cek user sudah ada
	if _, exists := users[email]; exists {
		return errors.New("user already exists")
	}

	// hash password
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	// simpan ke fake DB
	users[email] = hash

	return nil
}
