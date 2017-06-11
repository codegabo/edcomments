package routes

import (
	"github.com/golang-es/edcomments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// si el token es valido se ejecutara el SubRouter
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
