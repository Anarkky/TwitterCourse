package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Anarkky/TwitterCourse/db"
	"github.com/Anarkky/TwitterCourse/models"
)

/* Register is a function to create ner user in data base*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	//Body is stream type (read only) and can be only readed once then it's destroyed
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error while decoding User's body: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Error: - Email is mandatory", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Error: - Password must have at least 6 chars", 400)
		return
	}

	_, userFound, _ := db.VerifyExistingUser(t.Email)
	if userFound {
		http.Error(w, "Error: - Email user already exists", 400)
		return
	}

	_, status, err := db.RegisterUser(t)
	if err != nil {
		http.Error(w, "Error while creating the new user: "+err.Error(), 400)
		return
	}

	if status {
		http.Error(w, "Error: - Could not create the new user: ", 400)
		return
	}
}
