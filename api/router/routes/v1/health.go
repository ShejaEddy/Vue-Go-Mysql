package v1

import (
	"github.com/go-xorm/xorm"
	"net/http"
)

func Health(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("V1 is healthy"))
	}

}
