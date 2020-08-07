package autorizacion

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pedroalvaroccoyllocondori/apis_go/modelo"
)

func GenerarToken(data *modelo.Login) (string, error) {
	claim := modelo.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "ICAsoft",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenFirmado, err := token.SignedString(llaveFirma)
	if err != nil {
		return "", err
	}
	return tokenFirmado, nil
}
