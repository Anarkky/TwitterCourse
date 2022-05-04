package main

import (
	"log"

	"github.com/Anarkky/TwitterCourse/db"
	"github.com/Anarkky/TwitterCourse/handlers"
)

func main() {
	if db.VerifyConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.TwitterHandler()
}
