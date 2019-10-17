package v1

import (
	"github.com/go-xorm/xorm"
	jwt "github.com/projects/vue-go-1st/api/jwt"
	Routes "github.com/projects/vue-go-1st/api/router/routes"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")

		if _, err := jwt.ParseToken(token); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) (SubRoute map[string]Routes.SubRoutePackage) {
	SubRoute = map[string]Routes.SubRoutePackage{
		"/v1": {
			Routes: Routes.Routes{
				Routes.Route{
					Name:        "V1Health",
					Method:      "GET",
					Path:        "/health",
					HandlerFunc: Health(db),
				},
			},
			Middleware: Middleware,
		},
	}
	return SubRoute
}
