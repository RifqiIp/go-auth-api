package service

import "errors"

/*
Login
→ business logic
*/
func Login(email, password string) error {

	// simulasi user benar
	if email != "rifqi@example.com" || password != "password123" {
		return errors.New("invalid email or password")
	}

	return nil
}
