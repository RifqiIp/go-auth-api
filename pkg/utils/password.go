package utils

import "golang.org/x/crypto/bcrypt"

/*
HashPassword
→ ubah plain password jadi hash
*/
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(bytes), err
}

/*
ComparePassword
→ bandingkan hash vs password input
*/
func ComparePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
