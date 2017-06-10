package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/golang-es/edcomments/migration"
	"github.com/golang-es/edcomments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	// paquete de golang
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la base de datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Ha comenzado la migración")
		migration.Migrate()
		log.Println("Ha finalizado la migración")
	}

	// Inicia la ruta
	router := routes.InitRoutes()

	// Inicia los middleWares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	log.Println("Iniciado el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("ha finalizado la ejecucion del programa")
}
