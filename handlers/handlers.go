package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* TwitterHandler set default PORT and set handler listener*/
func TwitterHandler() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	hanlder := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, hanlder))

}
