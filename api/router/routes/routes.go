package routes

import (
	"github.com/go-xorm/xorm"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})

}
func GetRoutes(db *xorm.Engine) Routes {
	return Routes{
		Route{
			Name:        "HealthCheck",
			Method:      "GET",
			Path:        "/health",
			HandlerFunc: Health(db),
		},
		Route{
			Name:        "Login",
			Method:      "POST",
			Path:        "/auth/login",
			HandlerFunc: Login(db),
		},
		Route{
			Name:        "GetUser",
			Method:      "GET",
			Path:        "/auth/user",
			HandlerFunc: GetUser(db),
		},
		Route{
			Name:        "Check",
			Method:      "POST",
			Path:        "/auth/check",
			HandlerFunc: Check(db),
		},
	}
}
