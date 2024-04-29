package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	MAX_ATTEMPTS := 20
	for attempts := 0; err != nil && attempts < MAX_ATTEMPTS; attempts++ {
		hashed_password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	}
	return string(hashed_password), err
}

func ComparePasswords(hashed_password string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password)) == nil
}
