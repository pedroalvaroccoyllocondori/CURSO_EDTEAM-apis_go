package autorizacion

import (
	"errors"
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

func ValidarToken(t string) (modelo.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &modelo.Claim{}, verificarFuncion)
	if err != nil {
		return modelo.Claim{}, err
	}
	if !token.Valid {
		return modelo.Claim{}, errors.New("token no valido")
	}
	claim, ok := token.Claims.(*modelo.Claim)
	if !ok {
		return modelo.Claim{}, errors.New("no se puede obtener los claims")
	}
	return *claim, nil
}

func verificarFuncion(t *jwt.Token) (interface{}, error) {
	return llaveVerificacion, nil
}
