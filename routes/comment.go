package routes

import (
  "github.com/gorilla/mux"
  "github.com/urfave/negroni"
  "github.com/golang-es/edcomments/controllers"
)
// si el token es valido se ejecutara el SubRouter
 func SetCommentRouter(router *mux.Router) {
   prefix "/api/comments"
   subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter(.StrictSlash)(true)
   SubRouter.HandleFunc("/", controllers.CommentsCreate).Methods("POST")

   router.PathPrefix(prefix).Handler(
     negroni.New(
       negroni.HandlerFunc(controllers.VakidateToken),
       negroni.Wrap(SupRouter),
     )
   )
 }
