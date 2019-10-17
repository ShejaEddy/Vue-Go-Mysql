package routes

import (
	"encoding/json"
	"github.com/go-xorm/xorm"
	jwt "github.com/projects/vue-go-1st/api/jwt"
	"net/http"
)

func GetUser(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenVal := r.Header.Get("X-Access-Token")

		user, err := jwt.GetUserFromToken(db, tokenVal)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}

}
