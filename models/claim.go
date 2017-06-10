// Claim el paquete jwt-go ayuda a la creacion de los tokens
package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim token de usuario
type Claim struct {
	User `json:"user"`
	// se llaman los campos standard del paquete jwt
	// como la fecha de expiración, uso, destino y demás
	jwt.StandardClaims
}
