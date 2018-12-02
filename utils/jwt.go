package utils

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserClaims with()
type UserClaims struct {
	ID int
	jwt.StandardClaims
}

// JwtSecretKey secret to sign JWT tokens
var JwtSecretKey = []byte(os.Getenv("SECRET"))

// SignJwt with()
func SignJwt(identifier int) (string, error) {
	claims := UserClaims{
		ID: identifier,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}
