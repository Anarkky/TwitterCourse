package middleware

import (
	"net/http"

	"github.com/Anarkky/TwitterCourse/db"
)

/* VerifyDB is a middleware that verify data base status*/
func VerifyDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.VerifyConnection() == 0 {
			http.Error(w, "Data Base connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
