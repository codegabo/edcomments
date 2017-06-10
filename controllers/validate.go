package controllers

import (
  "net/http"
//el paquete context es un paquete precompilado que permitira enviar la informacion del usuario entre los controladores
  "context"
  jwt "github.com/dgrijalva/jwt-go"
  "github.com/dgrijalva/jwt-go/request"
  "github.com/golang-es/edcomments/commons"
  "github.com/golang-es/edcomments/models"

)

//ValidateToken permite validar el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandleFunc) {
  var m models.Message
  // con esto se extrae la informacion del request
  token, err := request.ParseFromRequestWithClaims(
    r,
    request.OAuth2Extractor,
    &models.Claim{},
    func(t *jwt.Token) (interface{}, error) {
      return commons.PublicKey, nil
    },
  )

  if err != nil {
    m.Code = http.StatusUnauthorized
    switch err.(type){
    case *jwt.ValidationError:
     vError err.(*jwt.ValidationError)
        switch vError.Error {
      case jwt.ValidationErrorExpired:
        m.Message = "Su token ha expirado"
        commons.DisplayMessage(w, m)
        return
      case jwt.ValidationErrorSignatureInvalid:
        m.Message = "La firma del token no coincide"
        commons.DisplayMessage(w, m)
        return
      default:
        m.Message = "Su token no es válido"
        commons.DisplayMessage(w, m)
        return

      }
    }
  }
// si el token es valido se ejecuta el siguiente controlador
  if token.Valid {
    ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
    next(w, r.WithContext(ctx))
  } else {
    m.Code = http.StatusUnauthorized
    m.Message = "Su token no es válido"
    commons.DisplayMessage(w, m)
  }
}
