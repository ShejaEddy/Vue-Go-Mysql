package routes

import (
	"encoding/json"
	"github.com/go-xorm/xorm"
	jwt "github.com/projects/vue-go-1st/api/jwt"
	Passwords "github.com/projects/vue-go-1st/api/models/passwords"
	Users "github.com/projects/vue-go-1st/api/models/users"
	"net/http"
)

func Login(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		creds := &Users.User{}
		err := json.NewDecoder(r.Body).Decode(creds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		username := creds.Username
		password := creds.Password
		if len(username) == 0 || len(password) == 0 {
			http.Error(w, "An username and password are required", http.StatusBadRequest)
			return
		}

		users, err := Users.Index(db, &Users.User{Username: username})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			http.Error(w, "user or Password invalid", http.StatusUnauthorized)
			return
		}

		user := users[0]

		if !Passwords.Compare(user.Password, password) {
			http.Error(w, "Invalid Password", http.StatusUnauthorized)
			return
		}

		user.Password = ""

		token, err := jwt.CreateToken(user.ID)
		if err != nil {
			http.Error(w, "Unable to create jwt", http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(packet)
	}

}
