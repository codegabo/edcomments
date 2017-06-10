package models

import "github.com/jinzhu/gorm"

// User aqui se esta creando un API para ser consumida via json
// en json no se utilizan las primeras letras en mayusculas
// se usa unique para que no se pueda repetir ese campo
// la notacion del json no es obligatoria
// si el campo password esta vacio se omite con la notacion omitempty
// el GORM puede OMITIR EL GUARDADO en la BD con un -
type User struct {
	gorm.Model
	Username        string    `json:"username" gorm:"not null;unique"`
	Email           string    `json:"email" gorm:"not null;unique"`
	Fullname        string    `json:"fullname" gorm:"not null"`
	Password        string    `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string    `json:"confirmPassword,omitempty" gorm:"-"`
	Picture         string    `json:"picture"`
	Comments        []Comment `json:"comments,omityempty"`
}

func main() {

}
