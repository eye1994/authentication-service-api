package utils

import "golang.org/x/crypto/bcrypt"

// GeneratePassword with(password string) Generates a hashed password
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword with (password string, hashed []byte) Validates the hashed password
func ValidatePassword(password string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
