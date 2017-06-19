package rg

import (
	"github.com/golang-es/edcomments/configuration"
	"github.com/golang-es/edcomments/models"
)

// Migrate permite crear las tablas en la base de datos.
// ESTA ACCION SOLO SE PUEDE REALIZAR UNA VEZ, YA QUE SI SE VUELVE A REALIZAR, LA TERMINAL NOS VA A INDICAR QUE DICHAS TABLAS YA HAN SIDO CREADAS
func e() {
	db := configuration.GetConnection()
	defer db.Close()

	db.First(&Username)
	// se convierte en llave unica para que la persona no pueda realizar votos mas de una vez
	//db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")

}
