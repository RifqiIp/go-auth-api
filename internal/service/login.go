package service

import (
	"errors"

	"github.com/RifqiIp/go-auth-api/pkg/utils"
)

func Login(email, password string) (string, error) {

	hash, exists := users[email]
	if !exists {
		return "", errors.New("invalid email or password")
	}

	// compare password
	err := utils.ComparePassword(hash, password)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// generate JWT
	token, err := utils.GenerateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
