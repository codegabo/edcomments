package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang-es/edcomments/models"
)

// así se delaran variables a nivel de paquete
var (
	privateKey *rsa.PrivateKey
	// Como esta variable es publica  se incia con mayuscula, PublicKey se usar para validar el token
	PublicKey *rsa.PublicKey
)

// esta funcion se ejecutara cada vez que se importe desde otro paquete
func init() {
	//ReadFile, funciona para leer un archivo del sistema operativo
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}

	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("No se puedo leer el archivo público")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privateKey", err)
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a PublicKey", err)
	}
}

// GenerateJWT genera el token para el cliente
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			//si se desea que la sesion expire se desomenta la siguiente linesa
			//ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: "Escuela Digital",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return result
}
