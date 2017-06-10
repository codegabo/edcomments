package main

import (
	"flag"
	"log"

	"github.com/golang-es/edcomments/migration"
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
}
