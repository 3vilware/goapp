package main

import (
	"log"
	"flag"
	"github.com/goapp/migration"
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
}