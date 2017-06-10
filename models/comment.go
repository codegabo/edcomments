package models

import "github.com/jinzhu/gorm"

// Comment comentario  del sistema
type Comment struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	ParentID uint   `json:"parentid"`
	Votes    int32  `json:"votes"`
	Content  string `json:"content"`
	// el campo hasvote permite saber si el usuario ya realizo un voto al comentario
	HasVote int8 `json:"hasvote" gorm:"-"`
	// el sampo user contiene la informacion del usuario que creo el comentario
	// esto es un hack
	// con este hack puedo consultar la información de usuario cuando se consulte un comentario
	User []User `json:"user,omitempty"`
	// el campo childrenson las respuestas de los comentarios
	Children []Comment `json:"children,omitempty"`
}
