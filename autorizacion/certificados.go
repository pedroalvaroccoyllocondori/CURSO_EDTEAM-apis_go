package autorizacion

import (
	"crypto/rsa"
	"io/ioutil"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	llaveFirma        *rsa.PrivateKey
	llaveVerificacion *rsa.PublicKey
	once              sync.Once
)

//funcion exportada
func CargarArchivos(archivoPrivado, archivoPublico string) error {
	var err error
	once.Do(func() {
		err = cargarArchivos(archivoPrivado, archivoPublico)
	})
	return err
}

//funcion no exportada
func cargarArchivos(archivoPrivado, archivoPublico string) error {
	bytesPrivados, err := ioutil.ReadFile(archivoPrivado)
	if err != nil {
		return err
	}
	bytesPublicos, err := ioutil.ReadFile(archivoPublico)
	if err != nil {
		return err
	}

	return parseRSA(bytesPrivados, bytesPublicos)
}

func parseRSA(bytesPrivados, bytesPublicos []byte) error {
	var err error
	llaveFirma, err = jwt.ParseRSAPrivateKeyFromPEM(bytesPrivados)
	if err != nil {
		return err
	}
	llaveVerificacion, err = jwt.ParseRSAPublicKeyFromPEM(bytesPublicos)
	if err != nil {
		return err
	}
	return nil
}
