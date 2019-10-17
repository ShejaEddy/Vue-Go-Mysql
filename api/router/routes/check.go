package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	jwt "github.com/projects/vue-go-1st/api/jwt"
	"net/http"
)

func Check(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenVal := r.Header.Get("X-Access-Token")
		userId, err := jwt.ParseToken(tokenVal)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		packet, err := json.Marshal(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}

}
