package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Anarkky/TwitterCourse/middleware"
	"github.com/Anarkky/TwitterCourse/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* TwitterHandler set default PORT and set handler listener*/
func TwitterHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.VerifyDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	hanlder := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, hanlder))

}
