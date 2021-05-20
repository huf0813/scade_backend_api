package domain

import "github.com/dgrijalva/jwt-go"

type CustomToken struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
