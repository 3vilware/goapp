package main

import (
	"net/http"
	"log"
	"flag"
	"github.com/goapp/migration"
	"github.com/goapp/routes"
	"github.com/urfave/negroni"

)

func main(){
	var migrate string 
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la base de datos") //variable que guarda, nombre del parametro
	flag.Parse()

	if migrate == "yes"{
		log.Println(" Comenzo la migracion...")
		migration.Migrate()
		log.Printf("Migracion Terminada  [OK]\n")
	}

	// inicia las rutas
	router := routes.InitRoutes()

	// inicia middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// iniciamos servidor
	server := &http.Server{
		Addr: ":8080",
		Handler: n,
	}

	log.Println("Iniciando servidor en 8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizo la ejecucion del servidor")

}