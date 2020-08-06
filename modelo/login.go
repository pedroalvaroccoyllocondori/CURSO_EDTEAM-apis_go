package modelo

import (
	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Email      string `json:"email"`
	contraseña string `json:"contraseña"`
}
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
